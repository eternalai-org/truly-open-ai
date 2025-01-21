package services

import (
	"context"
	"decentralized-inference/internal/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *Service) GetJobConfig(job string) (*models.JobConfig, error) {
	var jobConfig models.JobConfig
	if err := s.db.GetDB().Collection(models.JobConfig{}.CollectionName()).FindOne(context.Background(), bson.D{{"job_name", job}}).Decode(&jobConfig); err != nil {
		return nil, err
	}
	return &jobConfig, nil
}

func (s *Service) addJobConfig(ctx context.Context, jobConfig *models.JobConfig) error {
	jobConfig.CreatedAt = time.Now().UTC()
	jobConfig.UpdatedAt = jobConfig.CreatedAt
	if _, err := s.db.GetDB().Collection(models.JobConfig{}.CollectionName()).InsertOne(ctx, jobConfig); err != nil {
		return err
	}
	return nil
}

func (s *Service) updateJobLastRun(ctx context.Context, job string, lastRun time.Time) error {
	updateDoc := bson.D{
		{"last_run", lastRun},
		{"updated_at", time.Now().UTC()},
	}

	if _, err := s.db.GetDB().Collection(models.JobConfig{}.CollectionName()).
		UpdateOne(ctx, bson.D{{"job_name", job}}, bson.D{{"$set", updateDoc}}); err != nil {
		return err
	}
	return nil
}

func (s *Service) updateJobConfigByFilter(ctx context.Context, filter bson.M, data map[string]interface{}) error {
	if _, err := s.db.GetDB().Collection(models.JobConfig{}.CollectionName()).
		UpdateOne(ctx, filter, bson.D{{"$set", data}}); err != nil {
		return err
	}
	return nil
}
func (s *Service) updateJobLastOffset(ctx context.Context, job string, offset int64) error {
	if _, err := s.db.GetDB().Collection(models.JobConfig{}.CollectionName()).UpdateOne(ctx, bson.D{{"job_name", job}}, bson.D{{"$set", bson.D{{"offset", offset}}}}); err != nil {
		return err
	}
	return nil
}
