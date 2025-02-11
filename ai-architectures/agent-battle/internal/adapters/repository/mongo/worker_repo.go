package mongo

import (
	"context"
	"time"

	"agent-battle/internal/core/model"
	"agent-battle/pkg/drivers/mongodb"
	"agent-battle/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IWorkerRepo interface {
	mongodb.Repository[*model.Worker]
	FindOrCreateWorkerByName(ctx context.Context, name string) (*model.Worker, error)
}

type workerRepo struct {
	mongodb.Repository[*model.Worker]
}

func (repo *workerRepo) FindOrCreateWorkerByName(ctx context.Context, name string) (*model.Worker, error) {
	worker := &model.Worker{}
	err := repo.FindOne(ctx, bson.M{"name": name}, worker)
	if err != nil && !utils.IsErrNoDocuments(err) {
		return nil, err
	}

	if !worker.Id.IsZero() {
		return worker, nil
	}

	worker = &model.Worker{
		Status: model.WorkerStatusEnable,
		Name:   name,
		State:  &model.State{},
	}
	worker.DateCreated = time.Now()
	worker.DateModified = worker.DateCreated
	worker.Id = primitive.NewObjectID()
	if _, err := repo.Create(ctx, worker); err != nil {
		return nil, err
	}
	return worker, nil
}

func NewWorkerRepo(db *mongo.Database, secondaryDb ...*mongo.Database) IWorkerRepo {
	return &workerRepo{
		Repository: mongodb.NewBaseRepository(&model.Worker{}, db, secondaryDb...),
	}
}
