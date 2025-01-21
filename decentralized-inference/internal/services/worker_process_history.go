package services

import (
	"context"
	"decentralized-inference/internal/models"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Service) GetModelWorkerProcessHistoryByInferenceId(ctx context.Context, inferenceId string) ([]*models.ModelWorkerProcessHistories, error) {
	filter := bson.M{
		"inference_id": inferenceId,
	}

	cursor, err := s.db.GetDB().Collection(models.ModelWorkerProcessHistories{}.CollectionName()).Find(ctx, filter, nil)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	var items []*models.ModelWorkerProcessHistories
	err = cursor.All(ctx, &items)
	if err != nil {
		return nil, err
	}
	return items, nil
}

func (s *Service) UpdateModelWorkerProcessHistory(ctx context.Context, id primitive.ObjectID, updatedData map[string]interface{}) error {
	filter := bson.D{
		{"_id", id},
	}
	update := bson.M{
		"$set": updatedData,
	}
	_, err := s.db.GetDB().Collection(models.ModelWorkerProcessHistories{}.CollectionName()).
		UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateModelWorkerProcessHistoryByInferenceId(ctx context.Context, filter bson.D, updatedData map[string]interface{}) error {
	update := bson.M{
		"$set": updatedData,
	}
	_, err := s.db.GetDB().Collection(models.ModelWorkerProcessHistories{}.CollectionName()).
		UpdateMany(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) InsertModelWorkerProcessHistories(ctx context.Context, data *models.ModelWorkerProcessHistories) error {
	result, err := s.db.GetDB().Collection(data.CollectionName()).InsertOne(ctx, data)
	if err != nil {
		return err
	}
	data.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}

func (s *Service) GetModelWorkerProcessHistoryByFilter(ctx context.Context, filter bson.M) (*models.ModelWorkerProcessHistories, error) {
	var data models.ModelWorkerProcessHistories
	err := s.db.GetDB().Collection(models.ModelWorkerProcessHistories{}.CollectionName()).FindOne(ctx, filter).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}

func (s *Service) GetModelWorkerProcessHistoryByID(ctx context.Context, id string) (*models.ModelWorkerProcessHistories, error) {
	var data models.ModelWorkerProcessHistories
	primitiveId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = s.db.GetDB().Collection(models.ModelWorkerProcessHistories{}.CollectionName()).FindOne(ctx, bson.M{
		"_id": primitiveId,
	}).Decode(&data)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}
		return nil, err
	}
	return &data, nil
}

func (s *Service) GetManyModelWorkerProcessHistoryByFilter(ctx context.Context, filter bson.M, opts ...*options.FindOptions) ([]*models.ModelWorkerProcessHistories, error) {
	var data []*models.ModelWorkerProcessHistories
	cur, err := s.db.GetDB().Collection(models.ModelWorkerProcessHistories{}.CollectionName()).Find(ctx, filter, opts...)
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
