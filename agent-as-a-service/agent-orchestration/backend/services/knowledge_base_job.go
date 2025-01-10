package services

import (
	"context"
	"strings"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/aikb20"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

func (s *Service) JobCreateAgentKnowledgeBase(ctx context.Context) error {
	err := s.JobRunCheck(
		ctx, "JobCreateAgentKnowledgeBase",
		func() error {
			return s.CreateAgentKnowledgeBase(ctx)
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}

func (s *Service) CreateAgentKnowledgeBase(ctx context.Context) error {
	list, err := s.KnowledgeUsecase.GetKnowledgeBaseByStatus(ctx, models.KnowledgeBaseStatusDone, 0, 10)
	if err != nil {
		return err
	}
	for _, item := range list {
		err = s.DeployAgentKnowledgeBase(ctx, item)
		if err != nil {
			logger.Info("JobCreateAgentKnowledgeBase", "CreateAgentKnowledgeBase",
				zap.Any("err", err), zap.Any("kb", item))
		}
	}
	return nil
}

func (s *Service) DeployAgentKnowledgeBase(ctx context.Context, info *models.KnowledgeBase) error {
	var err error
	var tx, contractAddress string
	priKey := s.conf.WalletDeployAIKB20
	if !s.conf.GetConfigKeyBool(info.NetworkID, "zk_sync") {
		client := s.GetEthereumClient(ctx, info.NetworkID)
		input := common.FromHex(aikb20.EternalAIKB20MetaData.Bin)
		tx, err = client.Transact("", priKey, input, nil)
		if err != nil {
			return err
		}
		txReceipt, err := client.WaitMinedTxReceipt(ctx, common.HexToHash(tx))
		if err != nil {
			return err
		}
		contractAddress = txReceipt.ContractAddress.Hex()
	} else {
		client := s.GetZkClient(ctx, info.NetworkID)
		tx, contractAddress, err = client.DeployContract(priKey, aikb20.EternalAIKB20MetaData.Bin, "")
		if err != nil {
			return err
		}
	}

	info.KB20TxDeploy = strings.ToLower(tx)
	info.KB20Address = strings.ToLower(contractAddress)
	info.Status = models.KnowledgeBaseStatusMinted
	err = s.KnowledgeUsecase.UpdateKnowledgeBaseById(ctx, info.ID, map[string]interface{}{
		"status":          info.Status,
		"kb_20_tx_deploy": info.KB20TxDeploy,
		"kb_20_address":   info.KB20Address,
	})
	if err != nil {
		return err
	}
	return nil
}
