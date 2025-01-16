package services

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/daos"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/errs"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
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
							lp, err := s.dao.FirstLaunchpadByID(
								tx,
								lp.ID,
								map[string][]interface{}{},
								true,
							)
							if err != nil {
								return errs.NewError(err)
							}
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
							lpm, err = s.dao.FirstLaunchpadMemberByID(
								tx,
								lpm.ID,
								map[string][]interface{}{},
								true,
							)
							if err != nil {
								return errs.NewError(err)
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
								maxFundAmount := models.SubBigFloats(&lp.MaxFundBalance.Float, &lp.FundBalance.Float)
								if maxFundAmount.Cmp(models.SubBigFloats(&lpm.MaxFundBalance.Float, &lpm.FundBalance.Float)) > 0 {
									maxFundAmount = models.SubBigFloats(&lpm.MaxFundBalance.Float, &lpm.FundBalance.Float)
								}
								if maxFundAmount.Cmp(big.NewFloat(0)) > 0 {
									fundBalance := &lptx.Amount.Float
									refundBalance := big.NewFloat(0)
									if fundBalance.Cmp(maxFundAmount) > 0 {
										fundBalance = maxFundAmount
										refundBalance = models.SubBigFloats(&lptx.Amount.Float, fundBalance)
									}
									err = tx.Model(lpm).Updates(
										map[string]interface{}{
											"fund_balance":   gorm.Expr("fund_balance + ?", numeric.NewBigFloatFromFloat(fundBalance)),
											"refund_balance": gorm.Expr("refund_balance + ?", numeric.NewBigFloatFromFloat(refundBalance)),
											"total_balance":  gorm.Expr("total_balance + ?", lptx.Amount),
										},
									).Error
									if err != nil {
										return errs.NewError(err)
									}
									err = tx.Model(lp).Updates(
										map[string]interface{}{
											"fund_balance":   gorm.Expr("fund_balance + ?", numeric.NewBigFloatFromFloat(fundBalance)),
											"refund_balance": gorm.Expr("refund_balance + ?", numeric.NewBigFloatFromFloat(refundBalance)),
											"total_balance":  gorm.Expr("total_balance + ?", lptx.Amount),
										},
									).Error
									if err != nil {
										return errs.NewError(err)
									}
									// update token balance
									lpm, err = s.dao.FirstLaunchpadMemberByID(
										tx,
										lpm.ID,
										map[string][]interface{}{},
										true,
									)
									if err != nil {
										return errs.NewError(err)
									}
									lp, err := s.dao.FirstLaunchpadByID(
										tx,
										lp.ID,
										map[string][]interface{}{},
										true,
									)
									if err != nil {
										return errs.NewError(err)
									}
									tokenBalance := models.QuoBigFloats(
										models.MulBigFloats(&lpm.FundBalance.Float, &lp.TgeBalance.Float),
										&lp.MaxFundBalance.Float,
									)
									lpm.TokenBalance = numeric.NewBigFloatFromFloat(tokenBalance)
									err = s.dao.Save(tx, lpm)
									if err != nil {
										return errs.NewError(err)
									}
								} else {
									err = tx.Model(lpm).Updates(
										map[string]interface{}{
											"refund_balance": gorm.Expr("refund_balance + ?", lptx.Amount),
											"total_balance":  gorm.Expr("total_balance + ?", lptx.Amount),
										},
									).Error
									if err != nil {
										return errs.NewError(err)
									}
									err = tx.Model(lp).Updates(
										map[string]interface{}{
											"refund_balance": gorm.Expr("refund_balance + ?", lptx.Amount),
											"total_balance":  gorm.Expr("total_balance + ?", lptx.Amount),
										},
									).Error
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
				tokenAddress, txHash, err := func() (string, string, error) {
					tokenAddr, tokenHash, err := s.GetEthereumClient(ctx, m.NetworkID).
						DeployDAOTToken(
							s.GetAddressPrk(daoPoolAddress),
							m.TokenName,
							m.TokenSymbol,
						)
					if err != nil {
						return tokenAddr, tokenHash, errs.NewError(err)
					}
					time.Sleep(5 * time.Second)
					err = s.GetEthereumClient(ctx, m.NetworkID).TransactionConfirmed(tokenHash)
					if err != nil {
						return tokenAddr, tokenHash, errs.NewError(err)
					}
					return tokenAddr, tokenHash, nil
				}()
				if err != nil {
					err = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"token_address":        strings.ToLower(tokenAddress),
								"status":               models.LaunchpadStatusTokenError,
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
								"token_address":        strings.ToLower(tokenAddress),
								"total_supply":         numeric.NewBigFloatFromString("1000000000"),
								"deploy_token_tx_hash": txHash,
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

func (s *Service) AgentSettleDAOToken(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentSettleDAOToken_%d", memeID),
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
			if m.Status == models.LaunchpadStatusEnd && m.SettleFundTxHash == "" {
				daoPoolAddress := strings.ToLower(s.conf.GetConfigKeyString(m.NetworkID, "dao_pool_address"))
				txHash, err := func() (string, error) {
					hash, err := s.GetEthereumClient(ctx, m.NetworkID).
						DAOTreasurySettleFund(
							m.Address,
							s.GetAddressPrk(daoPoolAddress),
							models.ConvertBigFloatToWei(&m.FundBalance.Float, 18),
						)
					if err != nil {
						return hash, errs.NewError(err)
					}
					_ = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"settle_fund_tx_hash": hash,
							},
						).Error
					time.Sleep(5 * time.Second)
					err = s.GetEthereumClient(ctx, m.NetworkID).TransactionConfirmed(hash)
					if err != nil {
						return hash, errs.NewError(err)
					}
					return hash, nil
				}()
				if err != nil {
					_ = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"settle_fund_tx_hash": txHash,
								"status":              models.LaunchpadStatusSettleError,
							},
						).Error
					return errs.NewError(err)
				} else {
					err = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"status":              models.LaunchpadStatusSettled,
								"settle_fund_tx_hash": txHash,
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

func (s *Service) AgentAddLiquidityDAOToken(ctx context.Context, memeID uint) error {
	err := s.JobRunCheck(
		ctx,
		fmt.Sprintf("AgentAddLiquidityDAOToken_%d", memeID),
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
			if m.Status == models.LaunchpadStatusTge && m.AddLiquidityTxHash == "" {
				daoPoolAddress := strings.ToLower(s.conf.GetConfigKeyString(m.NetworkID, "dao_pool_address"))
				{
					hash, err := s.GetEthereumClient(ctx, m.NetworkID).
						Erc20ApproveMaxCheck(
							m.TokenAddress,
							s.GetAddressPrk(daoPoolAddress),
							helpers.HexToAddress(m.Address),
						)
					if err != nil {
						return errs.NewError(err)
					}
					if hash != "" {
						time.Sleep(5 * time.Second)
						err = s.GetEthereumClient(ctx, m.NetworkID).TransactionConfirmed(hash)
						if err != nil {
							return errs.NewError(err)
						}
					}
				}
				txHash, err := func() (string, error) {
					hash, err := s.GetEthereumClient(ctx, m.NetworkID).
						DAOTreasuryAddLiquidity(
							m.Address,
							s.GetAddressPrk(daoPoolAddress),
							helpers.HexToAddress(m.TokenAddress),
						)
					if err != nil {
						return hash, errs.NewError(err)
					}
					_ = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"add_liquidity_tx_hash": hash,
							},
						).Error
					time.Sleep(5 * time.Second)
					err = s.GetEthereumClient(ctx, m.NetworkID).TransactionConfirmed(hash)
					if err != nil {
						return hash, errs.NewError(err)
					}
					return hash, nil
				}()
				if err != nil {
					_ = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"add_liquidity_tx_hash": txHash,
								"status":                models.LaunchpadStatusPoolError,
							},
						).Error
					return errs.NewError(err)
				} else {
					err = daos.GetDBMainCtx(ctx).Model(&m).
						Updates(
							map[string]interface{}{
								"status":                models.LaunchpadStatusDone,
								"add_liquidity_tx_hash": txHash,
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
