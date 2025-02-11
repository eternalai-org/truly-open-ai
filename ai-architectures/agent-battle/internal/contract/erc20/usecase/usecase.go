package usecase

import (

"agent-battle/pkg/constants"
"context"
"crypto/ecdsa"
"errors"
"fmt"
"github.com/ethereum/go-ethereum"
"github.com/ethereum/go-ethereum/core/types"
"golang.org/x/crypto/sha3"
"math/big"
"strings"
"time"

	"agent-battle/internal/contract/erc20"
	"agent-battle/internal/core/port"
	"agent-battle/pkg/eth"
	"agent-battle/pkg/logger"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var Module = fx.Module("contract_erc20_usecase", fx.Provide(NewContractErc20Usecase))

type contractErc20Usecase struct {
	contractInstance *erc20.Erc20
	contractClient   *ethclient.Client

	contractAddress common.Address
	contractRpc     string
}

func NewContractErc20Usecase() (port.IContractErc20Usecase, error) {
	uc := &contractErc20Usecase{
		contractRpc:     viper.GetString("CONTRACT_RPC"),
		contractAddress: common.HexToAddress(viper.GetString("CONTRACT_ERC20_ADDRESS")),
	}
	client, err := newClient(uc.contractRpc)
	if err != nil {
		return nil, err
	}

	uc.contractClient = client
	ct, err := erc20.NewErc20(uc.contractAddress, client)
	if err != nil {
		return nil, err
	}

	uc.contractInstance = ct
	ctx := context.Background()
	if err := uc.healthCheck(ctx); err != nil {
		return nil, err
	}
	return uc, nil
}

func (uc *contractErc20Usecase) FilterTransfer(ctx context.Context, startBlock, endBlock uint64, from, to []common.Address) (*erc20.Erc20TransferIterator, error) {
	return uc.contractInstance.FilterTransfer(&bind.FilterOpts{
		Start:   startBlock,
		End:     &endBlock,
		Context: ctx,
	}, from, to)
}

func (uc *contractErc20Usecase) BalanceOfAddress(ctx context.Context, address string) (*big.Int, error) {
	balance, err := uc.contractInstance.BalanceOf(nil, common.HexToAddress(address))
	if err != nil {
		balance, err = uc.contractClient.BalanceAt(ctx, common.HexToAddress(address), nil)
		if err != nil {
			return nil, err
		}
	}
	return balance, nil
}

func newClient(contractRpc string) (*ethclient.Client, error) {
	client, err := eth.NewEthClient(contractRpc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (uc *contractErc20Usecase) getChainID(ctx context.Context) (*big.Int, error) {
	chainID, err := uc.contractClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	if chainID == nil {
		return nil, ethereum.NotFound
	}

	return chainID, nil
}

func (uc *contractErc20Usecase) healthCheck(ctx context.Context) error {
	chainId, err := uc.getChainID(ctx)
	if err != nil {
		return err
	}

	logger.AtLog.Infof("contractErc20Usecase#ContractAddress %s - chainId: %d - version: %s", uc.contractAddress, chainId, "")
	return nil
}

func (uc *contractErc20Usecase) CurrentBlockNumber(ctx context.Context) (uint64, error) {
	block, err := uc.contractClient.BlockNumber(ctx)
	if err != nil {
		return 0, err
	}
	return block, nil
}

func (uc *contractErc20Usecase) TransferToken(ctx context.Context, toAddress string, amount *big.Int, privateKey string) (string, error) {
	// Get private key
	privateKey = strings.TrimPrefix(privateKey, "0x") // remove 0x prefix
	privateKeyBytes, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey := privateKeyBytes.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	chainId, err := uc.contractClient.ChainID(ctx)
	if err != nil {
		return "", err
	}

	txParams, err := uc.createTokenTxParams(ctx, fromAddress, common.HexToAddress(toAddress), amount)
	if err != nil {
		return "", err
	}
	tx := types.NewTx(txParams)

	// Sign transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), privateKeyBytes)
	if err != nil {
		return "", err
	}

	// Send transaction to the network
	err = uc.contractClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}

	// Wait for transaction receipt (optional, for confirmation)
	receipt, err := bind.WaitMined(context.Background(), uc.contractClient, signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction receipt: %w", err)
	}

	// Check if the transaction was successful
	if receipt.Status != 1 {
		return "", fmt.Errorf("transaction failed: %s", receipt.TxHash.Hex())
	}

	logger.GetLoggerInstanceFromContext(ctx).Info("TransferToken", zap.Any("from_address", fromAddress), zap.Any("receipt", receipt))

	return receipt.TxHash.Hex(), nil
}

// TransferETH transfers ETH from the fromAddress to the toAddress
func (uc *contractErc20Usecase) TransferETH(ctx context.Context, toAddress string, amount *big.Int, privateKey string) (string, error) {
	// Get private key
	privateKey = strings.TrimPrefix(privateKey, "0x") // remove 0x prefix
	privateKeyBytes, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", fmt.Errorf("failed to parse private key: %w", err)
	}

	publicKey := privateKeyBytes.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	txParams, err := uc.createNativeTxParams(ctx, fromAddress, common.HexToAddress(toAddress), amount)
	if err != nil {
		return "", err
	}
	tx := types.NewTx(txParams)
	return uc.sendTransaction(ctx, tx, fromAddress, privateKeyBytes)
}

// sendTransaction signs and sends the transaction to the network
// returns the transaction hash if successful
func (uc *contractErc20Usecase) sendTransaction(
	ctx context.Context,
	tx *types.Transaction,
	fromAddress common.Address,
	privateKeyBytes *ecdsa.PrivateKey,
) (string, error) {
	chainId, err := uc.contractClient.ChainID(ctx)
	if err != nil {
		return "", err
	}

	// Sign transaction
	signer := types.NewEIP155Signer(chainId)
	signedTx, err := types.SignTx(tx, signer, privateKeyBytes)
	if err != nil {
		return "", err
	}

	// Send transaction to the network
	err = uc.contractClient.SendTransaction(ctx, signedTx)
	if err != nil {
		return "", err
	}

	// Wait for transaction receipt (optional, for confirmation)
	receipt, err := bind.WaitMined(context.Background(), uc.contractClient, signedTx)
	if err != nil {
		return "", fmt.Errorf("failed to wait for transaction receipt: %w", err)
	}

	// wait around 1 second for the funds to be indexed by the node
	time.Sleep(1 * time.Second)
	latestReceipt, err := uc.contractClient.TransactionReceipt(ctx, receipt.TxHash)
	if err != nil {
		return "", fmt.Errorf("failed to get transaction receipt: %w", err)
	}

	// Check if the transaction receipt is the latest
	if latestReceipt.TxHash.Hex() != receipt.TxHash.Hex() {
		return "", fmt.Errorf("transaction receipt not found: %s", receipt.TxHash.Hex())
	}

	// Check if the transaction was successful
	if latestReceipt.Status != 1 {
		return "", fmt.Errorf("transaction failed: %s", latestReceipt.TxHash.Hex())
	}

	logger.GetLoggerInstanceFromContext(ctx).Info("Transfer", zap.Any("from_address", fromAddress), zap.Any("receipt", latestReceipt))

	return receipt.TxHash.Hex(), nil
}

// EstimateGasFee estimates the gas fee for a token transfer
func (uc *contractErc20Usecase) EstimateGasFee(
	ctx context.Context,
	fromAddress string,
	toAddress string,
	amount *big.Int,
) (*big.Int, error) {
	txParams, err := uc.createTokenTxParams(
		ctx,
		common.HexToAddress(fromAddress),
		common.HexToAddress(toAddress),
		amount,
		true /* doNotNeedCheckBalance */,
	)
	if err != nil {
		return nil, err
	}

	gasFee := txParams.Gas * txParams.GasPrice.Uint64()
	return new(big.Int).SetUint64(gasFee), nil
}

// createNativeTxParams creates transaction parameters for native token transfer
func (uc *contractErc20Usecase) createNativeTxParams(
	ctx context.Context,
	fromAddress,
	toAddress common.Address,
	amount *big.Int,
) (*types.LegacyTx, error) {
	// get nonce of fromAddress
	nonce, err := uc.contractClient.NonceAt(ctx, fromAddress, nil)
	if err != nil {
		return nil, err
	}

	// get suggested gas price
	gasPrice, tipCap, err := uc.estimateGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	// create call message
	callMsg := ethereum.CallMsg{
		From:      fromAddress,
		To:        &toAddress,
		Gas:       constants.DefaultEthereumGasLimit,
		GasFeeCap: gasPrice, // maxFeePerGas (base fee + priority fee)
		GasTipCap: tipCap,   // maxPriorityFeePerGas (priority fee)
		Data:      nil,
		Value:     amount,
	}

	// estimate gas limit
	gas, err := uc.contractClient.EstimateGas(ctx, callMsg)
	if err != nil {
		// if failed to estimate gas, use default gas limit
		logger.GetLoggerInstanceFromContext(ctx).Error(
			"failed to estimate gas",
			zap.String("from_address", fromAddress.String()),
			zap.String("to_address", toAddress.String()),
			zap.Error(err),
		)
	}

	// set default gas limit if the estimated gas is too low
	if gas < constants.DefaultEthereumGasLimit {
		gas = constants.DefaultEthereumGasLimit
	}

	// check min eth balance
	gasFee := gas * gasPrice.Uint64()
	err = uc.checkEthBalanceOfFromAddress(ctx, fromAddress, amount, gasFee)
	if err != nil {
		return nil, err
	}

	txParams := &types.LegacyTx{
		Nonce:    nonce,
		Gas:      gas,
		GasPrice: gasPrice,
		To:       &toAddress,
		Value:    callMsg.Value,
	}

	return txParams, nil
}

// createTokenTxParams creates transaction parameters for token transfer
func (uc *contractErc20Usecase) createTokenTxParams(
	ctx context.Context,
	fromAddress,
	toAddress common.Address,
	amount *big.Int,
	doNotNeedCheckBalance ...bool,
) (*types.LegacyTx, error) {
	// estimate gas
	value := big.NewInt(0) // for token transfer, the value always equals 0
	nonce, err := uc.contractClient.NonceAt(ctx, fromAddress, nil)
	if err != nil {
		return nil, err
	}

	// get suggested gas price
	gasPrice, tipCap, err := uc.estimateGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	// create call message
	callMsg := ethereum.CallMsg{
		From:      fromAddress,
		To:        &uc.contractAddress, // for token transfer, to is the contract address
		Gas:       constants.DefaultEthereumTokenGasLimit,
		GasFeeCap: gasPrice, // maxFeePerGas (base fee + priority fee)
		GasTipCap: tipCap,   // maxPriorityFeePerGas (priority fee)
		Data:      uc.generateTransferData(toAddress, amount),
		Value:     value,
	}

	// estimate gas limit
	gas, err := uc.contractClient.EstimateGas(ctx, callMsg)
	if err != nil {
		// if failed to estimate gas, use default gas limit
		logger.GetLoggerInstanceFromContext(ctx).Error(
			"failed to estimate gas for token transfer",
			zap.Error(err),
			zap.String("from_address", fromAddress.String()),
			zap.String("to_address", toAddress.String()),
		)
	}
	// set default gas limit if the estimated gas is too low
	if gas < constants.DefaultEthereumTokenGasLimit {
		gas = constants.DefaultEthereumTokenGasLimit
	}

	// check min eth balance
	if len(doNotNeedCheckBalance) == 0 {
		minEthBalance := gas * gasPrice.Uint64()
		err = uc.checkBalanceOfFromAddress(ctx, fromAddress, amount, minEthBalance)
		if err != nil {
			return nil, err
		}
	}

	txParams := &types.LegacyTx{
		Nonce:    nonce,
		Gas:      gas,
		GasPrice: gasPrice,
		To:       &uc.contractAddress, // for token transfer, to is the contract address
		Value:    value,
		Data:     callMsg.Data,
	}

	return txParams, nil
}

func (uc *contractErc20Usecase) estimateGasPrice(ctx context.Context) (*big.Int, *big.Int, error) {
	// get suggested gas tip cap
	tipCap, err := uc.contractClient.SuggestGasTipCap(ctx)
	if err != nil {
		return nil, nil, err
	}

	// get suggested gas price
	baseFee, err := uc.contractClient.SuggestGasPrice(ctx)
	if err != nil {
		return nil, nil, err
	}

	gasPrice := baseFee.Add(baseFee, tipCap)

	// failed with 219190 gas: max fee per gas less than block base fee
	buffer := big.NewInt(1e8) // Add a 0.1 Gwei buffer to the gas price
	gasPrice = gasPrice.Add(gasPrice, buffer)

	// failed with 219190 gas: max fee per gas less than block base fee
	return gasPrice, tipCap, nil
}

func (uc *contractErc20Usecase) checkEthBalanceOfFromAddress(
	ctx context.Context,
	fromAddress common.Address,
	amount *big.Int,
	estimatedFee uint64,
) error {
	// check native balance
	ethBalance, err := uc.contractClient.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return err
	}

	// Check if native ethBalance is sufficient
	zero := big.NewInt(0)
	if ethBalance.Cmp(zero) <= 0 {
		return fmt.Errorf("insufficient eth balance")
	}

	// Check if native ethBalance is sufficient to pay f
	amountWithFee := new(big.Int).Add(amount, new(big.Int).SetUint64(estimatedFee))
	if ethBalance.Cmp(amountWithFee) < 0 {
		return fmt.Errorf("insufficient eth balance")
	}

	return nil
}

// checkBalanceOfFromAddress checks if the fromAddress has enough balance to pay for gas and transfer amount
func (uc *contractErc20Usecase) checkBalanceOfFromAddress(
	ctx context.Context,
	fromAddress common.Address,
	amount *big.Int,
	minEthBalance uint64,
) error {
	// check native balance
	ethBalance, err := uc.contractClient.BalanceAt(ctx, fromAddress, nil)
	if err != nil {
		return err
	}

	// Check if native ethBalance is sufficient
	zero := big.NewInt(0)
	if ethBalance.Cmp(zero) <= 0 {
		return fmt.Errorf("insufficient eth balance")
	}

	// Check if native ethBalance is sufficient to pay for gas
	if ethBalance.Uint64() < minEthBalance {
		logger.GetLoggerInstanceFromContext(ctx).Error(
			"insufficient eth balance",
			zap.String("from_address", fromAddress.String()),
			zap.Uint64("eth_balance", ethBalance.Uint64()),
			zap.Uint64("min_eth_balance", minEthBalance),
		)
		return fmt.Errorf("insufficient eth balance to pay for gas")
	}

	tokenBalance, err := uc.contractInstance.BalanceOf(nil, fromAddress)
	if err != nil {
		return err
	}

	if tokenBalance.Cmp(amount) < 0 {
		return fmt.Errorf("insufficient token balance")
	}

	return nil
}

func (uc *contractErc20Usecase) generateTransferData(toAddress common.Address, amount *big.Int) []byte {
	transferFnSignature := []byte("transfer(address,uint256)")
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	methodID := hash.Sum(nil)[:4]

	// pad address and amount to 32 bytes
	// which is the length of a word in the EVM
	// and the length of the data expected by the transfer function
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)

	var data []byte

	// append methodID, paddedAddress, and paddedAmount
	data = append(data, methodID...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)

	return data
}
