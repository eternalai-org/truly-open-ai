package services

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/models"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/ethapi"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/types/numeric"
	"github.com/jinzhu/gorm"
)

func (s *Service) LaunchpadErc20TokenTransferEvent(tx *gorm.DB, networkID uint64, event *ethapi.Erc20TokenTransferEventResp) error {
	switch networkID {
	case models.BASE_CHAIN_ID:
		{
			if s.conf.ExistsedConfigKey(networkID, "eai_contract_address") {
				contractAddress := strings.ToLower(event.ContractAddress)
				eaiAddress := s.conf.GetConfigKeyString(networkID, "eai_contract_address")
				toAddress := strings.ToLower(event.To)
				if !strings.EqualFold(toAddress, models.ETH_ZERO_ADDRESS) && strings.EqualFold(contractAddress, eaiAddress) {
					eventId := fmt.Sprintf("%d_%s_%d", networkID, event.TxHash, event.Index)
					lptx, err := s.dao.FirstLaunchpadTransaction(
						tx,
						map[string][]interface{}{
							"event_id = ?": {eventId},
						},
						map[string][]interface{}{},
						[]string{},
					)
					if err != nil {
						return errs.NewError(err)
					}
					if lptx == nil {
						lp, err := s.dao.FirstLaunchpad(
							tx,
							map[string][]interface{}{
								"address = ?": {toAddress},
							},
							map[string][]interface{}{},
							[]string{},
						)
						if err != nil {
							return errs.NewError(err)
						}
						if lp != nil {
							lpm, err := s.dao.FirstLaunchpadMember(
								tx,
								map[string][]interface{}{
									"launchpad_id = ?": {lp.ID},
								},
								map[string][]interface{}{},
								[]string{},
							)
							if err != nil {
								return errs.NewError(err)
							}
							if lpm == nil {
								lpm = &models.LaunchpadMember{
									UserAddress: toAddress,
									LaunchpadID: lp.ID,
									Tier:        string(models.LaunchpadTier3),
								}
								err = s.dao.Create(tx, lpm)
								if err != nil {
									return errs.NewError(err)
								}
							}
							lptx = &models.LaunchpadTransaction{
								NetworkID:   networkID,
								LaunchpadID: lp.ID,
								EventId:     eventId,
								TxHash:      event.TxHash,
								Type:        models.LaunchpadTransactionTypeDeposit,
								UserAddress: strings.ToLower(event.From),
								Amount:      numeric.NewBigFloatFromFloat(models.ConvertWeiToBigFloat(event.Value, 18)),
								Status:      models.LaunchpadTransactionStatusDone,
							}
							err = s.dao.Create(tx, lptx)
							if err != nil {
								return errs.NewError(err)
							}
							if lptx.Status == models.LaunchpadTransactionStatusDone {
								err = tx.Model(lpm).
									UpdateColumn("fund_balance", gorm.Expr("fund_balance + ?", lptx.Amount)).
									UpdateColumn("total_balance", gorm.Expr("total_balance + ?", lptx.Amount)).
									Error
								if err != nil {
									return errs.NewError(err)
								}
								err = tx.Model(lp).
									UpdateColumn("fund_balance", gorm.Expr("fund_balance + ?", lptx.Amount)).
									UpdateColumn("total_balance", gorm.Expr("total_balance + ?", lptx.Amount)).
									Error
								if err != nil {
									return errs.NewError(err)
								}
							}
						}
					}
				}
			}
		}
	}
	return nil
}

func (s *Service) AgentDeployDAOToken(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentDeployDAOToken_%d", memeID),
		func() error {
			m, err := s.dao.FirstLaunchpadByID(
				daos.GetDBMainCtx(ctx),
				memeID,
				map[string][]interface{}{},
				false,
			)
			if err != nil {
				return errs.NewError(err)
			}
			if m == nil {
				return errs.NewError(errs.ErrBadRequest)
			}
			if m.TokenSymbol == "" {
				for i := 0; i < 3; i++ {
					tokenInfo, err := s.GenerateTokenInfoFromSystemPrompt(ctx, "", m.Name)
					if err != nil {
						continue
					}
					m.TokenSymbol = tokenInfo.TokenSymbol
					m.TokenName = tokenInfo.TokenName
					m.TokenImageUrl = tokenInfo.TokenImageUrl
					err = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"token_name":      m.TokenName,
								"token_symbol":    m.TokenSymbol,
								"token_image_url": m.TokenImageUrl,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
					break
				}
				if m.TokenSymbol == "" {
					_ = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"status": models.LaunchpadStatusTokenError,
							},
						).Error
					return errs.NewError(err)
				}
			}
			if m.Status == models.LaunchpadStatusEnd && m.TokenSymbol != "" && m.TokenAddress == "" {
				daoPoolAddress := strings.ToLower(s.conf.GetConfigKeyString(m.NetworkID, "dao_pool_address"))
				var tokenAddress, txHash string
				for i := 0; i < 3; i++ {
					tokenAddress, txHash, err = s.GetEthereumClient(ctx, m.NetworkID).
						DeployDAOTToken(
							s.GetAddressPrk(daoPoolAddress),
							m.TokenName,
							m.TokenSymbol,
						)
					if err != nil {
						time.Sleep(10 * time.Second)
						continue
					}
					time.Sleep(10 * time.Second)
					err = s.GetEthereumClient(ctx, m.NetworkID).TransactionConfirmed(txHash)
					if err != nil {
						continue
					}
					break
				}
				if tokenAddress != "" {
					err = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"token_address":        strings.ToLower(tokenAddress),
								"total_supply":         numeric.NewBigFloatFromString("1000000000"),
								"deploy_token_tx_hash": txHash,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				} else {
					err = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"status": models.LaunchpadStatusTokenError,
							},
						).Error
					if err != nil {
						return errs.NewError(err)
					}
				}
			}
			return nil
		},
	)
	if err != nil {
		return errs.NewError(err)
	}
	return nil
}
