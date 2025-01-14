package knowledge

import (
	"context"
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/core/ports"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/internal/repository"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/logger"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/pkg/utils"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/serializers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/ethapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/lighthouse"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/trxapi"
	resty "github.com/go-resty/resty/v2"

	"go.uber.org/zap"
)

type knowledgeUsecase struct {
	knowledgeBaseRepo          repository.KnowledgeBaseRepo
	knowledgeBaseFileRepo      repository.KnowledgeBaseFileRepo
	agentInfoKnowledgeBaseRepo repository.IAgentInfoKnowledgeBaseRepo
	secretKey                  string

	networks      map[string]map[string]string
	ethApiMap     map[uint64]*ethapi.Client
	trxApi        *trxapi.Client
	ragApi        string
	lighthouseKey string
	webhookUrl    string
}

func (uc *knowledgeUsecase) CreateAgentInfoKnowledgeBase(ctx context.Context, model *models.AgentInfoKnowledgeBase) (*models.AgentInfoKnowledgeBase, error) {
	return uc.agentInfoKnowledgeBaseRepo.Create(ctx, model)
}

func (uc *knowledgeUsecase) WebhookFile(ctx context.Context, filename string, bytes []byte, id uint) (*models.KnowledgeBase, error) {
	kn, err := uc.knowledgeBaseRepo.GetKnowledgeBaseById(ctx, id)
	if err != nil {
		return nil, err
	}
	logger.Info("webhook_file", "start_webhook_file", zap.Any("knowledge_base_id", id), zap.Any("filename", filename))
	updatedFields := make(map[string]interface{})
	hash, err := lighthouse.UploadDataWithRetry(uc.lighthouseKey, fmt.Sprintf("%d_%s", time.Now().Unix(), filename), bytes)
	if err != nil {
		logger.Error("webhook_file_error", "upload_data_with_retry", zap.Error(err))
		updatedFields["status"] = models.KnowledgeBaseStatusProcessingFailed
		_ = uc.knowledgeBaseRepo.UpdateKnowledgeBaseById(ctx, id, updatedFields)
		return nil, err
	}

	updatedFields["status"] = models.KnowledgeBaseStatusDone
	updatedFields["filecoin_hash"] = hash
	if err := uc.knowledgeBaseRepo.UpdateKnowledgeBaseById(ctx, id, updatedFields); err != nil {
		return nil, err
	}
	return kn, nil
}

func (uc *knowledgeUsecase) Webhook(ctx context.Context, req *models.RagResponse) (*models.KnowledgeBase, error) {
	logger.Info("knowledgeUsecase", "Webhook", zap.Any("data", req))
	if req.Result == nil {
		return nil, nil
	}

	id, err := strconv.ParseUint(req.Result.Ref, 10, 64)
	if err != nil {
		return nil, err
	}

	kn, err := uc.knowledgeBaseRepo.GetKnowledgeBaseById(ctx, uint(id))
	if err != nil {
		return nil, err
	}

	updatedFields := make(map[string]interface{})
	updatedFields["filecoin_hash"] = req.Result.FilecoinHash
	if req.Result.FilecoinHash == "" {
		updatedFields["status"] = models.KnowledgeBaseStatusProcessingFailed
	} else {
		updatedFields["status"] = models.KnowledgeBaseStatusDone
	}
	// updatedFields["message"] = req.Message

	if err := uc.knowledgeBaseRepo.UpdateKnowledgeBaseById(ctx, kn.ID, updatedFields); err != nil {
		return nil, err
	}

	return kn, nil
}

func NewKnowledgeUsecase(
	knowledgeBaseRepo repository.KnowledgeBaseRepo,
	knowledgeBaseFileRepo repository.KnowledgeBaseFileRepo,
	agentInfoKnowledgeBaseRepo repository.IAgentInfoKnowledgeBaseRepo,
	secretKey string,
	ethApiMap map[uint64]*ethapi.Client,
	networks map[string]map[string]string,
	trxApi *trxapi.Client,
	ragApi string,
	lighthousekey string,
	webhookUrl string,
) ports.IKnowledgeUsecase {
	return &knowledgeUsecase{
		knowledgeBaseRepo:          knowledgeBaseRepo,
		knowledgeBaseFileRepo:      knowledgeBaseFileRepo,
		agentInfoKnowledgeBaseRepo: agentInfoKnowledgeBaseRepo,
		secretKey:                  secretKey,
		ethApiMap:                  ethApiMap,
		networks:                   networks,
		trxApi:                     trxApi,
		ragApi:                     ragApi,
		lighthouseKey:              lighthousekey,
		webhookUrl:                 webhookUrl,
	}
}

func (uc *knowledgeUsecase) CreateKnowledgeBase(ctx context.Context, req *serializers.CreateKnowledgeRequest) (*serializers.KnowledgeBase, error) {
	model := &models.KnowledgeBase{}
	if err := utils.Copy(model, req); err != nil {
		return nil, err
	}

	encryptedTempKey, tempAddr, err := utils.GenerateAddress(uc.secretKey)
	if err != nil {
		return nil, err
	}
	model.DepositAddress = strings.ToLower(tempAddr)
	model.DepositPrivKey = encryptedTempKey

	encryptedTempKey, tempAddr, err = utils.GenerateSolanaAddress(uc.secretKey)
	if err != nil {
		return nil, err
	}

	model.SolanaDepositAddress = strings.ToLower(tempAddr)
	model.SolanaDepositPrivKey = encryptedTempKey

	model.Status = models.KnowledgeBaseStatusWaitingPayment
	model.Fee = 1

	resp, err := uc.knowledgeBaseRepo.CreateKnowledgeBase(ctx, model)
	if err != nil {
		return nil, err
	}

	for _, f := range req.Files {
		file := &models.KnowledgeBaseFile{
			FileUrl:         f.Url,
			FileName:        f.Name,
			FileSize:        f.Size,
			KnowledgeBaseId: resp.ID,
		}
		_, err := uc.knowledgeBaseFileRepo.Create(ctx, file)
		if err != nil {
			return nil, err
		}
	}

	r, err := uc.knowledgeBaseRepo.GetKnowledgeBaseById(ctx, model.ID)
	if err != nil {
		return nil, err
	}
	result := &serializers.KnowledgeBase{}
	if err := utils.Copy(result, r); err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *knowledgeUsecase) ListKnowledgeBase(ctx context.Context, req *models.ListKnowledgeBaseRequest) ([]*serializers.KnowledgeBase, error) {
	resp, err := uc.knowledgeBaseRepo.ListKnowledgeBaseByAddress(ctx, req.UserAddress)
	if err != nil {
		return nil, err
	}
	result := []*serializers.KnowledgeBase{}
	if err := utils.Copy(&result, resp); err != nil {
		return nil, err
	}
	return result, nil
}

func (uc *knowledgeUsecase) GetKnowledgeBaseById(ctx context.Context, id uint) (*models.KnowledgeBase, error) {
	return uc.knowledgeBaseRepo.GetKnowledgeBaseById(ctx, id)
}

func (uc *knowledgeUsecase) DeleteKnowledgeBaseById(ctx context.Context, id uint) error {
	return uc.knowledgeBaseRepo.DeleteKnowledgeBaseById(ctx, id)
}

func (uc *knowledgeUsecase) GetKnowledgeBaseByStatus(ctx context.Context, status models.KnowledgeBaseStatus, offset, limit int) ([]*models.KnowledgeBase, error) {
	return uc.knowledgeBaseRepo.GetKnowledgeBaseByStatus(ctx, status, offset, limit)
}

func (uc *knowledgeUsecase) UpdateKnowledgeBaseById(ctx context.Context, id uint, updatedFields map[string]interface{}) error {
	return uc.knowledgeBaseRepo.UpdateKnowledgeBaseById(ctx, id, updatedFields)
}

func (uc *knowledgeUsecase) WatchWalletChange(ctx context.Context) error {
	offset := 0
	limit := 30
	for {
		resp, err := uc.knowledgeBaseRepo.GetKnowledgeBaseByStatus(
			ctx, models.KnowledgeBaseStatusWaitingPayment, offset, limit,
		)
		logger.Logger().Info("GetKnowledgeBaseByStatus", zap.Any("total", len(resp)))
		if err != nil {
			return err
		}

		if len(resp) == 0 {
			break
		}

		for _, k := range resp {
			if err := uc.checkBalance(ctx, k); err != nil {
				return err
			}

			if k.Status == models.KnowledgeBaseStatusPaymentReceipt {
				if _, err := uc.insertFilesToRAG(ctx, k); err != nil {
					return err
				}
			}
		}
		offset += len(resp)
	}
	return nil
}

func (uc *knowledgeUsecase) checkBalance(ctx context.Context, kn *models.KnowledgeBase) error {
	knPrice := new(big.Float).SetFloat64(kn.Fee)
	knPrice = knPrice.Mul(knPrice, big.NewFloat(1e18))
	_knPrice := new(big.Int)
	_knPrice, _ = knPrice.Int(_knPrice)

	logger.Info("WatchWalletChange", "checkOrderBalanceAndProcess start",
		zap.Any("knowledge_base", kn),
		zap.Any("knPrice", knPrice),
		zap.Any("_knPrice", _knPrice),
	)

	// balance, err := utils.GetBalanceOnSolanaChain(ctx, kn.SolanaDepositAddress)
	// if err != nil {
	// 	logger.Logger().Error("GetBalanceOnSolanaChain", zap.Error(err))
	// }

	// if balance.Cmp(_knPrice) >= 0 && _knPrice.Uint64() >= 0 {
	// 	kn.Status = models.KnowledgeBaseStatusPaymentReceipt
	// 	if err := uc.knowledgeBaseRepo.UpdateStatus(ctx, kn); err != nil {
	// 		return err
	// 	}
	// }

	for networkId, net := range uc.networks {
		nId, err := strconv.ParseUint(networkId, 10, 64)
		if err != nil {
			continue
		}

		ethClient, ok := uc.ethApiMap[nId]
		if !ok {
			if net["rpc_url"] == "" {
				continue
			}

			uc.ethApiMap[nId] = &ethapi.Client{
				BaseURL:           net["rpc_url"],
				MinGasPrice:       net["min_gas_price"],
				BTCL1:             strings.ToUpper(net["is_btc_l1"]) == "TRUE",
				BlockTimeDisabled: true,
			}
			ethClient = uc.ethApiMap[nId]
		}

		balance, err := uc.balanceOfAddress(ctx, kn.DepositAddress, ethClient, net)
		if err != nil {
			continue
		}

		if balance.Cmp(_knPrice) >= 0 && _knPrice.Uint64() > 0 {
			updatedFields := make(map[string]interface{})
			updatedFields["status"] = models.KnowledgeBaseStatusPaymentReceipt
			updatedFields["deposit_tx_hash"] = fmt.Sprintf("%s/address/%s", net["explorer_url"], kn.DepositAddress)
			updatedFields["deposit_chain_id"] = nId
			kn.Status = models.KnowledgeBaseStatusPaymentReceipt
			if err := uc.knowledgeBaseRepo.UpdateKnowledgeBaseById(ctx, kn.ID, updatedFields); err != nil {
				return err
			}
		}
	}
	return nil
}

func (uc *knowledgeUsecase) balanceOfAddress(_ context.Context, address string, client *ethapi.Client, netInfo map[string]string) (*big.Int, error) {
	chainId := client.ChainID()
	conAddress := netInfo["eai_contract_address"]
	if address == "0x0000000000000000000000000000000000000000" || chainId == 0 || conAddress == "" {
		return big.NewInt(0), nil
	}
	if strings.EqualFold(conAddress, "0x000000000000000000000000000000000000800A") {
		balance, err := client.Balance(address)
		if err != nil {
			return nil, err
		}
		return balance, nil
	}

	if chainId == models.TRON_CHAIN_ID {
		balance, err := uc.trxApi.Trc20Balance(conAddress, address)
		if err != nil {
			return nil, err
		}
		return balance, nil
	}

	balanace, err := client.Erc20Balance(conAddress, address)
	if err != nil {
		return nil, err
	}

	return balanace, nil
}

func (uc *knowledgeUsecase) insertFilesToRAG(ctx context.Context, kn *models.KnowledgeBase) (*models.InsertRagResponse, error) {
	resp := &models.InsertRagResponse{}
	if uc.webhookUrl == "" {
		uc.webhookUrl = "https://agent.api.eternalai.org/api/knowledge/webhook-file"
	}
	body := struct {
		FileUrls []string `json:"file_urls"`
		Ref      string   `json:"ref"`
		Hook     string   `json:"hook"`
	}{
		FileUrls: kn.FileUrls(),
		Ref:      fmt.Sprintf("%d", kn.ID),
		Hook:     fmt.Sprintf("%s/%d", uc.webhookUrl, kn.ID),
	}
	logger.Info("knowledgebase", "insert_file_to_rag", zap.Any("body", body))
	_, err := resty.New().R().SetContext(ctx).SetDebug(true).
		SetBody(body).
		SetResult(resp).
		Post(fmt.Sprintf("%s/api/insert", uc.ragApi))
	if err != nil {
		return nil, err
	}

	kn.Status = models.KnowledgeBaseStatusProcessing
	if err = uc.knowledgeBaseRepo.UpdateStatus(ctx, kn); err != nil {
		return nil, err
	}
	return resp, nil
}
