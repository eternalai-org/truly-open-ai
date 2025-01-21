package services

import (
	"context"
	"decentralized-inference/internal/models"
	"errors"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *Service) ListChainConfig(ctx context.Context) ([]*models.ChainConfig, error) {
	cur, err := s.db.GetDB().Collection(models.ChainConfig{}.CollectionName()).Find(ctx, bson.M{},
		&options.FindOptions{
			Sort: bson.M{"_id": -1},
		})
	if err != nil {
		return nil, err
	}
	var chainConfigs []*models.ChainConfig
	if err := cur.All(ctx, &chainConfigs); err != nil {
		return nil, err
	}

	return chainConfigs, nil
}

func (s *Service) FindChainConfig(ctx context.Context, chainId string) (*models.ChainConfig, error) {
	filter := bson.M{"chain_id": chainId}
	chainConfig := &models.ChainConfig{}
	if err := s.db.GetDB().Collection(models.ChainConfig{}.CollectionName()).FindOne(ctx, filter).
		Decode(chainConfig); err != nil {
		return nil, err
	}

	return chainConfig, nil
}

func (s *Service) GetListModelSupportsByChainId(ctx context.Context, chainId string) (map[string]string, error) {
	chainConfig, err := s.FindChainConfig(ctx, chainId)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return nil, err
	}

	if len(chainConfig.SupportModelNames) > 0 {
		return chainConfig.SupportModelNames, nil
	}
	return map[string]string{}, nil
}

func (s *Service) GetChainDefaultByModelName(ctx context.Context, modelName string) string {
	appConfig, err := s.GetAppConfig(ctx)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return ""
	}
	if len(appConfig.ModelToChain) > 0 {
		return appConfig.ModelToChain[modelName]
	}
	return ""
}
func (s *Service) InsertChainConfig(ctx context.Context, chain *models.ChainConfig) error {
	chain.CreatedAt = time.Now()
	_, err := s.db.GetDB().Collection(models.ChainConfig{}.CollectionName()).InsertOne(ctx, chain)
	return err
}

func (s *Service) UpdateChainConfigByFilter(ctx context.Context, filter bson.M, update bson.M) error {
	update["updated_at"] = time.Now().UTC()
	_, err := s.db.GetDB().Collection(models.ChainConfig{}.CollectionName()).UpdateOne(ctx, filter,
		bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) GetAppConfig(ctx context.Context) (*models.AppConfig, error) {
	filter := bson.M{}
	opt := options.FindOne().SetSort(bson.M{"_id": 1})
	info := &models.AppConfig{}
	if err := s.db.GetDB().Collection(info.CollectionName()).FindOne(ctx, filter, opt).
		Decode(info); err != nil {
		return nil, err
	}

	return info, nil
}

func (s *Service) InsertAppConfig(ctx context.Context, info *models.AppConfig) error {
	info.CreatedAt = time.Now()
	_, err := s.db.GetDB().Collection(info.CollectionName()).InsertOne(ctx, info)
	return err
}

func (s *Service) UpdateAppConfigByFilter(ctx context.Context, filter bson.M, update bson.M) error {
	_, err := s.db.GetDB().Collection(models.AppConfig{}.CollectionName()).UpdateOne(ctx, filter,
		bson.M{"$set": update})
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) AddOrUpdateContractSyncState(data *models.ContractSyncState) error {
	filter := bson.D{{"contract_address", strings.ToLower(data.ContractAddress)}}
	err := s.db.GetDB().Collection(models.ContractSyncState{}.CollectionName()).FindOneAndReplace(context.Background(), filter, data, options.FindOneAndReplace().SetUpsert(true))
	if err.Err() != nil {
		return err.Err()
	}
	return nil
}

func (s *Service) AddContractSyncState(data *models.ContractSyncState) error {
	_, err := s.db.GetDB().Collection(models.ContractSyncState{}.CollectionName()).InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) UpdateContractSyncStateByAddressAndJob(data *models.ContractSyncState) error {
	filter := bson.D{{"_id", data.ID}}
	data.UpdatedAt = time.Now()
	update := bson.D{{"$set", data}}
	_, err := s.db.GetDB().Collection(models.ContractSyncState{}.CollectionName()).
		UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetContractSyncState(contractAddress string, jobName string) (*models.ContractSyncState, error) {
	var contractSyncState models.ContractSyncState
	contractAddress = strings.ToLower(contractAddress)
	if err := s.db.GetDB().Collection(models.ContractSyncState{}.CollectionName()).FindOne(context.Background(),
		bson.D{
			{"contract_address", contractAddress},
			{"job", jobName},
		}).Decode(&contractSyncState); err != nil {
		return nil, err
	}
	return &contractSyncState, nil
}
