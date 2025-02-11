package mongodb

import (
	"context"
	"reflect"
	"time"

	"agent-battle/pkg/logger"

	pg "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const MongoIdGenCollectionName = "id-gen"

type Repository[T Model] interface {
	Database() *mongo.Database
	DatabaseSecondary() *mongo.Database
	Collection() *mongo.Collection
	Find(ctx context.Context, filters bson.M, result interface{}, opts ...*FindOptions) error
	FindById(ctx context.Context, id primitive.ObjectID, value interface{}, opts ...*FindOneOptions) error
	FindOne(ctx context.Context, filters bson.M, value interface{}, opts ...*FindOneOptions) error
	Aggregations(ctx context.Context, result interface{}, pipeline primitive.A, opts ...*AggregateOptions) error
	Filter(ctx context.Context, filters bson.M, sortFields []string, sortValues []int, page int64, limit int64, result interface{}) (int64, error)
	CountDocuments(ctx context.Context, filter bson.M, opts ...*CountOptions) (int64, error)
	FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	FindOneAndReplace(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult
	Update(ctx context.Context, model interface{}, id primitive.ObjectID, dateModified time.Time, opts ...*options.FindOneAndReplaceOptions) error
	UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (int64, error)
	UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (int64, error)
	UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (int64, error)
	Create(ctx context.Context, model interface{}, opts ...*options.InsertOneOptions) (primitive.ObjectID, error)
	CreateMany(ctx context.Context, models []interface{}, opts ...*options.InsertManyOptions) ([]primitive.ObjectID, error)
	CreateManyWithTransaction(ctx context.Context, models []interface{}, opts ...*options.TransactionOptions) ([]primitive.ObjectID, error)
	Delete(ctx context.Context, id []primitive.ObjectID, opts ...*options.DeleteOptions) error
	WithTransaction(ctx context.Context, fn func(sessCtx mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error)
	CreateCompoundIndex(ctx context.Context, compoundIndex []string, unique bool, opts ...*options.CreateIndexesOptions) error
	CreateIndices(ctx context.Context, indices []string, unique bool, opts ...*options.CreateIndexesOptions) error
	CreateIndexes(ctx context.Context, indices []string, value int32, unique bool, opts ...*options.CreateIndexesOptions) error
	NextId(ctx context.Context, sequenceName *string) (uint, error)

	Generic[T]
}
type Generic[T Model] interface {
	// Example pipeline:
	//	{
	//	    $facet: {
	//	      items: [
	//	        {
	//	          $sort: {
	//	            date_created: -1,
	//	            _id: -1
	//	          }
	//	        },
	//	        { $limit: 1 },
	//	        {
	//	          $lookup: {
	// 				...
	//	          }
	//	        },
	//			{
	//	          $project: {
	//	            ....
	//	          }
	//	        }
	//	    ],
	//	    count: [
	//	      {
	//	        $count: "count"
	//	      }
	//	    ]
	//	  }
	//	}
	AggregationFacets(ctx context.Context, pipeline primitive.A, opts ...*AggregateOptions) ([]T, int64, error)
	GnrFind(ctx context.Context, filters bson.M, opts ...*FindOptions) ([]T, error)
	GnrFindById(ctx context.Context, id primitive.ObjectID, opts ...*FindOneOptions) (T, error)
	GnrFindOne(ctx context.Context, filters bson.M, opts ...*FindOneOptions) (T, error)
}

type BaseRepository[T Model] struct {
	collectionName string
	db             *mongo.Database
	secondaryDB    *mongo.Database
}

func NewBaseRepository[T Model](model T, primaryDB *mongo.Database, secondaryDB ...*mongo.Database) Repository[T] {
	if reflect.ValueOf(model).Kind() != reflect.Ptr {
		logger.AtLog.Fatal("model must is pointer")
	}
	base := &BaseRepository[T]{
		collectionName: model.CollectionName(),
		db:             primaryDB,
	}
	if len(secondaryDB) > 0 {
		base.secondaryDB = secondaryDB[0]
	}
	return base
}

func (b *BaseRepository[T]) Database() *mongo.Database {
	return b.db
}

func (b *BaseRepository[T]) DatabaseSecondary() *mongo.Database {
	if b.secondaryDB != nil {
		return b.secondaryDB
	}
	return b.Database()
}

func (b *BaseRepository[T]) Collection() *mongo.Collection {
	return b.db.Collection(b.collectionName)
}

func (b *BaseRepository[T]) Find(ctx context.Context, filters bson.M, result interface{}, opts ...*FindOptions) error {
	db := b.Database()
	newOpts := make([]*options.FindOptions, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, opt.FindOptions)
		if opt.readPreference {
			db = b.DatabaseSecondary()
		}
	}
	cur, err := db.Collection(b.collectionName).Find(ctx, filters, newOpts...)
	if err != nil {
		return err
	}
	if err := cur.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func (b *BaseRepository[T]) FindById(ctx context.Context, id primitive.ObjectID, value interface{}, opts ...*FindOneOptions) error {
	return b.FindOne(ctx, bson.M{"_id": id}, value, opts...)
}

func (b *BaseRepository[T]) FindOne(ctx context.Context, filters bson.M, value interface{}, opts ...*FindOneOptions) error {
	db := b.Database()
	newOpts := make([]*options.FindOneOptions, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, opt.FindOneOptions)
		if opt.readPreference {
			db = b.DatabaseSecondary()
		}
	}
	res := db.Collection(b.collectionName).FindOne(ctx, filters, newOpts...)
	if res.Err() != nil {
		return res.Err()
	}
	if err := res.Decode(value); err != nil {
		return err
	}
	return nil
}

func (b *BaseRepository[T]) Aggregations(ctx context.Context, result interface{}, pipeline primitive.A, opts ...*AggregateOptions) error {
	db := b.Database()
	newOpts := make([]*options.AggregateOptions, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, opt.AggregateOptions)
		if opt.readPreference {
			db = b.DatabaseSecondary()
		}
	}
	cur, err := db.Collection(b.collectionName).Aggregate(ctx, pipeline, newOpts...)
	if err != nil {
		return err
	}
	if err := cur.All(ctx, result); err != nil {
		return err
	}
	return nil
}

func (b *BaseRepository[T]) CountDocuments(ctx context.Context, filter bson.M, opts ...*CountOptions) (int64, error) {
	db := b.Database()
	newOpts := make([]*options.CountOptions, 0, len(opts))
	for _, opt := range opts {
		newOpts = append(newOpts, opt.CountOptions)
		if opt.readPreference {
			db = b.DatabaseSecondary()
		}
	}
	total, err := db.Collection(b.collectionName).CountDocuments(ctx, filter, newOpts...)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (b *BaseRepository[T]) FindOneAndUpdate(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return b.db.Collection(b.collectionName).FindOneAndUpdate(ctx, filter, update, opts...)
}

func (b *BaseRepository[T]) FindOneAndReplace(ctx context.Context, filter interface{}, update interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	return b.db.Collection(b.collectionName).FindOneAndReplace(ctx, filter, update, opts...)
}

func (b *BaseRepository[T]) Filter(ctx context.Context, filters bson.M, sortFields []string, sortValue []int, page int64, limit int64, result interface{}) (int64, error) {
	query := pg.New(b.db.Collection(b.collectionName)).Decode(result).Context(ctx).Page(page).Limit(limit)
	if len(sortFields) < 1 {
		query = query.Sort("date_modified", -1)
	} else {
		for i, sort := range sortFields {
			if i < len(sortValue) {
				query = query.Sort(sort, sortValue[i])
			}
		}
	}
	aggPaginatedData, err := query.Filter(filters).Find()
	if err != nil {
		return 0, err
	}
	return aggPaginatedData.Pagination.Total, err
}

func (b *BaseRepository[T]) Create(ctx context.Context, model interface{}, opts ...*options.InsertOneOptions) (primitive.ObjectID, error) {
	result, err := b.db.Collection(b.collectionName).InsertOne(ctx, model, opts...)
	if err != nil {
		return primitive.NilObjectID, err
	}
	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id, nil
	}
	return primitive.NilObjectID, nil
}

func (b *BaseRepository[T]) CreateManyWithTransaction(ctx context.Context, models []interface{}, opts ...*options.TransactionOptions) ([]primitive.ObjectID, error) {
	session, err := b.db.Client().StartSession()
	if err != nil {
		return nil, err
	}
	defer session.EndSession(ctx)
	callback := func(sessionContext mongo.SessionContext) (interface{}, error) {
		return b.db.Collection(b.collectionName).InsertMany(sessionContext, models)
	}
	results, err := session.WithTransaction(ctx, callback, opts...)
	if err != nil {
		return nil, err
	}
	var ids []primitive.ObjectID
	if insertResult, ok := results.(*mongo.InsertManyResult); ok {
		for _, result := range insertResult.InsertedIDs {
			if id, ok := result.(primitive.ObjectID); ok {
				ids = append(ids, id)
			}
		}
	}
	return ids, err
}

func (b *BaseRepository[T]) Update(ctx context.Context, model interface{}, id primitive.ObjectID, dateModified time.Time, opts ...*options.FindOneAndReplaceOptions) error {
	res := b.db.Collection(b.collectionName).FindOneAndReplace(
		ctx,
		bson.D{
			{Key: "_id", Value: id},
			{Key: "date_modified", Value: dateModified},
		},
		model,
		opts...,
	)
	return res.Err()
}

func (b *BaseRepository[T]) Delete(ctx context.Context, ids []primitive.ObjectID, opts ...*options.DeleteOptions) error {
	_, err := b.db.Collection(b.collectionName).DeleteMany(
		ctx,
		bson.M{"_id": bson.M{"$in": ids}},
		opts...,
	)
	if err != nil {
		return err
	}
	return nil
}

func (b *BaseRepository[T]) CreateCompoundIndex(ctx context.Context, compoundIndex []string, unique bool, opts ...*options.CreateIndexesOptions) error {
	var indices bson.D
	for _, index := range compoundIndex {
		indices = append(indices, bson.E{
			Key:   index,
			Value: 1,
		})
	}
	index := mongo.IndexModel{
		Keys:    indices,
		Options: options.Index().SetUnique(unique),
	}
	_, err := b.db.Collection(b.collectionName).Indexes().CreateOne(ctx, index, opts...)
	return err
}

// func indirect(reflectValue reflect.Value) reflect.Value {
// 	for reflectValue.Kind() == reflect.Ptr {
// 		reflectValue = reflectValue.Elem()
// 	}
// 	return reflectValue
// }

// func indirectType(reflectType reflect.Type) (_ reflect.Type, isPtr bool) {
// 	for reflectType.Kind() == reflect.Ptr || reflectType.Kind() == reflect.Slice {
// 		reflectType = reflectType.Elem()
// 		isPtr = true
// 	}
// 	return reflectType, isPtr
// }

func (b *BaseRepository[T]) WithTransaction(ctx context.Context, callback func(sessCtx mongo.SessionContext) (interface{}, error), opts ...*options.TransactionOptions) (interface{}, error) {
	session, err := b.db.Client().StartSession()
	defer session.EndSession(ctx)
	if err != nil {
		return nil, err
	}
	return session.WithTransaction(ctx, callback, opts...)
}

func (b *BaseRepository[T]) CreateIndices(ctx context.Context, indices []string, unique bool, opts ...*options.CreateIndexesOptions) error {
	return b.CreateIndexes(ctx, indices, 1, unique, opts...)
}

func (b *BaseRepository[T]) CreateIndexes(ctx context.Context, indices []string, value int32, unique bool, opts ...*options.CreateIndexesOptions) error {
	var indexModels []mongo.IndexModel
	for _, index := range indices {
		indexModel := mongo.IndexModel{
			Keys: bson.D{
				{Key: index, Value: value},
			},
			Options: options.Index().SetUnique(unique),
		}
		indexModels = append(indexModels, indexModel)
	}
	_, err := b.db.Collection(b.collectionName).Indexes().CreateMany(ctx, indexModels, opts...)
	return err
}

func (b *BaseRepository[T]) NextId(ctx context.Context, sequenceName *string) (uint, error) {
	findOptions := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	name := b.collectionName
	if sequenceName != nil && len(*sequenceName) > 0 {
		name = *sequenceName
	}
	counter := struct {
		ID  string `json:"id" bson:"_id"`
		Seq uint   `json:"seq" bson:"seq"`
	}{}
	err := b.Database().Collection(MongoIdGenCollectionName).
		FindOneAndUpdate(ctx,
			bson.M{"_id": name},
			bson.M{"$inc": bson.M{"seq": 1}},
			findOptions,
		).Decode(&counter)
	if err != nil {
		return 0, err
	}
	return counter.Seq, nil
}

func (b *BaseRepository[T]) UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (int64, error) {
	result, err := b.db.Collection(b.collectionName).UpdateByID(ctx, id, update, opts...)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (b *BaseRepository[T]) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (int64, error) {
	result, err := b.db.Collection(b.collectionName).UpdateOne(ctx, filter, update, opts...)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (b *BaseRepository[T]) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (int64, error) {
	result, err := b.db.Collection(b.collectionName).UpdateMany(ctx, filter, update, opts...)
	if err != nil {
		return 0, err
	}
	return result.ModifiedCount, nil
}

func (b *BaseRepository[T]) CreateMany(ctx context.Context, models []interface{}, opts ...*options.InsertManyOptions) ([]primitive.ObjectID, error) {
	results, err := b.db.Collection(b.collectionName).InsertMany(ctx, models, opts...)
	if err != nil {
		return nil, err
	}
	var ids []primitive.ObjectID
	for _, idResult := range results.InsertedIDs {
		if id, ok := idResult.(primitive.ObjectID); ok {
			ids = append(ids, id)
		}
	}
	return ids, err
}

func (b *BaseRepository[T]) AggregationFacets(ctx context.Context, pipeline primitive.A, opts ...*AggregateOptions) ([]T, int64, error) {
	facetsPipeline := []bson.M{}
	noNeedCount := false
	if len(opts) > 0 {
		opt := opts[0]
		if len(opt.sort) > 0 {
			facetsPipeline = append(facetsPipeline, bson.M{"$sort": opt.sort})
		}
		if opt.page != nil && *opt.page > 0 && opt.limit != nil && *opt.limit > 0 {
			skip := (*opt.page - 1) * *opt.limit
			facetsPipeline = append(facetsPipeline, bson.M{"$skip": skip})
		}
		if opt.limit != nil && *opt.limit > 0 {
			facetsPipeline = append(facetsPipeline, bson.M{"$limit": *opt.limit})
		}
		if opt.noNeedCount != nil && *opt.noNeedCount {
			noNeedCount = true
		}
		facetsPipeline = append(facetsPipeline, opt.facetsPipeline...)
	}
	if noNeedCount {
		pipeline = append(
			pipeline,
			bson.M{
				"$facet": bson.M{
					"items": facetsPipeline,
				},
			})
	} else {
		pipeline = append(
			pipeline,
			bson.M{
				"$facet": bson.M{
					"items": facetsPipeline,
					"count": bson.A{
						bson.M{"$count": "count"},
					},
				},
			},
			bson.M{
				"$project": bson.M{
					"items": 1,
					"total": bson.M{
						"$arrayElemAt": []interface{}{"$count.count", 0},
					},
				},
			})
	}
	result := []*Facet[T]{}
	if err := b.Aggregations(ctx, &result, pipeline, opts...); err != nil {
		return nil, 0, err
	}
	if len(result) <= 0 {
		return nil, 0, nil
	}
	return result[0].GetItems(), result[0].GetTotal(), nil
}

func (b *BaseRepository[T]) GnrFind(ctx context.Context, filters bson.M, opts ...*FindOptions) ([]T, error) {
	var rs []T
	if err := b.Find(ctx, filters, &rs, opts...); err != nil {
		return nil, err
	}
	return rs, nil
}

func (b *BaseRepository[T]) GnrFindById(ctx context.Context, id primitive.ObjectID, opts ...*FindOneOptions) (rs T, err error) {
	value := *new(T)
	if err := b.FindById(ctx, id, &value, opts...); err != nil {
		return rs, err
	}
	return value, nil
}

func (b *BaseRepository[T]) GnrFindOne(ctx context.Context, filters bson.M, opts ...*FindOneOptions) (rs T, err error) {
	value := *new(T)
	if err := b.FindOne(ctx, filters, &value, opts...); err != nil {
		return rs, err
	}
	return value, nil
}
