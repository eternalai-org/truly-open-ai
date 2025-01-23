package ethapi

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/eternal-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/memenonfungiblepositionmanager"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (c *Client) MemeNonfungiblePositionManagerMint(contractAddr string, prkHex string, weth9 common.Address, sqrtPriceX96 *big.Int, params *memenonfungiblepositionmanager.INonfungiblePositionManagerMintParams) (string, error) {
	pbkHex, prk, err := c.parsePrkAuth(prkHex)
	if err != nil {
		return "", err
	}
	gasPrice, gasTipCap, err := c.GetCachedGasPriceAndTipCap()
	if err != nil {
		return "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	value := big.NewInt(0)
	if !strings.EqualFold(weth9.Hex(), params.Token0.Hex()) {
		if params.Amount0Desired.Cmp(big.NewInt(0)) > 0 {
			allowance, err := c.Erc20Allowance(params.Token0.Hex(), pbkHex.Hex(), contractAddr)
			if err != nil {
				return "", err
			}
			if allowance.Cmp(big.NewInt(0)) <= 0 {
				approveHash, err := c.Erc20ApproveMax(
					params.Token0.Hex(),
					prkHex,
					contractAddr,
				)
				if err != nil {
					return "", err
				}
				time.Sleep(10 * time.Second)
				err = c.WaitMined(approveHash)
				if err != nil {
					return "", err
				}
			}
		}
	} else {
		value = params.Amount0Desired
	}
	if !strings.EqualFold(weth9.Hex(), params.Token1.Hex()) {
		if params.Amount1Desired.Cmp(big.NewInt(0)) > 0 {
			allowance, err := c.Erc20Allowance(params.Token1.Hex(), pbkHex.Hex(), contractAddr)
			if err != nil {
				return "", err
			}
			if allowance.Cmp(big.NewInt(0)) <= 0 {
				approveHash, err := c.Erc20ApproveMax(
					params.Token1.Hex(),
					prkHex,
					contractAddr,
				)
				if err != nil {
					return "", err
				}
				time.Sleep(10 * time.Second)
				err = c.WaitMined(approveHash)
				if err != nil {
					return "", err
				}
			}
		}
	} else {
		value = params.Amount1Desired
	}
	contractAddress := helpers.HexToAddress(contractAddr)
	// EstimateGas
	instanceABI, err := abi.JSON(strings.NewReader(memenonfungiblepositionmanager.NonfungiblePositionManagerABI))
	if err != nil {
		return "", err
	}
	multicallBytes := [][]byte{}
	{
		multicallData, err := instanceABI.Pack(
			"createAndInitializePoolIfNecessary",
			params.Token0,
			params.Token1,
			params.Fee,
			sqrtPriceX96,
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	{
		multicallData, err := instanceABI.Pack(
			"mint",
			memenonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
				Token0:         params.Token0,
				Token1:         params.Token1,
				Fee:            params.Fee,
				TickLower:      params.TickLower,
				TickUpper:      params.TickUpper,
				Amount0Desired: params.Amount0Desired,
				Amount1Desired: params.Amount1Desired,
				Amount0Min:     params.Amount0Min,
				Amount1Min:     params.Amount1Min,
				Recipient:      pbkHex,
				Deadline:       params.Deadline,
			},
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	if strings.EqualFold(weth9.Hex(), params.Token0.Hex()) || strings.EqualFold(weth9.Hex(), params.Token1.Hex()) {
		multicallData, err := instanceABI.Pack(
			"refundETH",
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	dataBytes, err := instanceABI.Pack(
		"multicall",
		multicallBytes,
	)
	if err != nil {
		return "", err
	}
	gasNumber, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  pbkHex,
		To:    &contractAddress,
		Data:  dataBytes,
		Value: value,
	})
	if err != nil {
		return "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}
	nonceAt, err := client.PendingNonceAt(context.Background(), pbkHex)
	if err != nil {
		return "", err
	}
	rawTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainID)),
		Nonce:     nonceAt,
		GasFeeCap: gasPrice,
		GasTipCap: gasTipCap,
		Gas:       (gasNumber * 12 / 10),
		To:        &contractAddress,
		Value:     value,
		Data:      dataBytes,
	})
	signedTx, err := types.SignTx(rawTx, types.NewLondonSigner(big.NewInt(int64(chainID))), prk)
	if err != nil {
		return "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}

func (c *Client) MemeNonfungiblePositionManagerMintData(recipient string, sqrtPriceX96 *big.Int, params *memenonfungiblepositionmanager.INonfungiblePositionManagerMintParams) ([]byte, error) {
	instanceABI, err := abi.JSON(strings.NewReader(memenonfungiblepositionmanager.NonfungiblePositionManagerABI))
	if err != nil {
		return nil, err
	}
	multicallBytes := [][]byte{}
	{
		multicallData, err := instanceABI.Pack(
			"createAndInitializePoolIfNecessary",
			params.Token0,
			params.Token1,
			params.Fee,
			sqrtPriceX96,
		)
		if err != nil {
			return nil, err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	{
		multicallData, err := instanceABI.Pack(
			"mint",
			memenonfungiblepositionmanager.INonfungiblePositionManagerMintParams{
				Token0:         params.Token0,
				Token1:         params.Token1,
				Fee:            params.Fee,
				TickLower:      params.TickLower,
				TickUpper:      params.TickUpper,
				Amount0Desired: params.Amount0Desired,
				Amount1Desired: params.Amount1Desired,
				Amount0Min:     params.Amount0Min,
				Amount1Min:     params.Amount1Min,
				Recipient:      helpers.HexToAddress(recipient),
				Deadline:       params.Deadline,
			},
		)
		if err != nil {
			return nil, err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	{
		multicallData, err := instanceABI.Pack(
			"refundETH",
		)
		if err != nil {
			return nil, err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	dataBytes, err := instanceABI.Pack(
		"multicall",
		multicallBytes,
	)
	if err != nil {
		return nil, err
	}
	return dataBytes, nil
}

func (c *Client) MemeNonfungiblePositionManagerBurn(contractAddr string, adminPrk string, weth9 common.Address, tokenId *big.Int) (string, error) {
	pbkHex, prk, err := c.parsePrkAuth(adminPrk)
	if err != nil {
		return "", err
	}
	gasPrice, gasTipCap, err := c.GetCachedGasPriceAndTipCap()
	if err != nil {
		return "", err
	}
	client, err := c.getClient()
	if err != nil {
		return "", err
	}
	nonceAt, err := client.PendingNonceAt(context.Background(), pbkHex)
	if err != nil {
		return "", err
	}
	contractAddress := helpers.HexToAddress(contractAddr)
	// EstimateGas
	instanceABI, err := abi.JSON(strings.NewReader(memenonfungiblepositionmanager.NonfungiblePositionManagerMetaData.ABI))
	if err != nil {
		return "", err
	}
	positionmanager, err := memenonfungiblepositionmanager.NewNonfungiblePositionManager(helpers.HexToAddress(contractAddr), client)
	if err != nil {
		return "", err
	}
	pos, err := positionmanager.Positions(&bind.CallOpts{}, tokenId)
	if err != nil {
		return "", err
	}
	multicallBytes := [][]byte{}
	{
		multicallData, err := instanceABI.Pack(
			"decreaseLiquidity",
			memenonfungiblepositionmanager.UniswapV3BrokerDecreaseLiquidityParams{
				TokenId:    tokenId,
				Liquidity:  pos.Liquidity,
				Amount0Min: big.NewInt(0),
				Amount1Min: big.NewInt(0),
				Deadline:   big.NewInt(time.Now().Add(60 * time.Second).Unix()),
			},
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	isNativeReceive := strings.EqualFold(weth9.Hex(), pos.Token0.Hex()) || strings.EqualFold(weth9.Hex(), pos.Token1.Hex())
	{
		amountMax := big.NewInt(0).Sub(new(big.Int).Exp(big.NewInt(2), big.NewInt(128), nil), big.NewInt(1))
		multicallData, err := instanceABI.Pack(
			"collect",
			memenonfungiblepositionmanager.UniswapV3BrokerCollectParams{
				TokenId:    tokenId,
				Recipient:  helpers.HexToAddress("0x0000000000000000000000000000000000000000"),
				Amount0Max: amountMax,
				Amount1Max: amountMax,
			},
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	if isNativeReceive {
		multicallData, err := instanceABI.Pack(
			"unwrapWETH",
			big.NewInt(0),
			pbkHex,
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	if !strings.EqualFold(weth9.Hex(), pos.Token0.Hex()) {
		multicallData, err := instanceABI.Pack(
			"sweepToken",
			pos.Token0,
			big.NewInt(0),
			pbkHex,
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	if !strings.EqualFold(weth9.Hex(), pos.Token1.Hex()) {
		multicallData, err := instanceABI.Pack(
			"sweepToken",
			pos.Token1,
			big.NewInt(0),
			pbkHex,
		)
		if err != nil {
			return "", err
		}
		multicallBytes = append(
			multicallBytes,
			multicallData,
		)
	}
	dataBytes, err := instanceABI.Pack(
		"multicall",
		multicallBytes,
	)
	if err != nil {
		return "", err
	}
	gasNumber, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: pbkHex,
		To:   &contractAddress,
		Data: dataBytes,
	})
	if err != nil {
		return "", err
	}
	chainID, err := c.GetChainID()
	if err != nil {
		return "", err
	}
	rawTx := types.NewTx(&types.DynamicFeeTx{
		ChainID:   big.NewInt(int64(chainID)),
		Nonce:     nonceAt,
		GasFeeCap: gasPrice,
		GasTipCap: gasTipCap,
		Gas:       (gasNumber * 12 / 10),
		To:        &contractAddress,
		Value:     big.NewInt(0),
		Data:      dataBytes,
	})
	signedTx, err := types.SignTx(rawTx, types.NewLondonSigner(big.NewInt(int64(chainID))), prk)
	if err != nil {
		return "", err
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}
	return signedTx.Hash().Hex(), nil
}
