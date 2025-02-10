package eth

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"errors"
	"math/big"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// NewEthClient creates a new Ethereum client for HTTP connections
func NewEthClient(rpc string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// NewEthWsClient creates a new Ethereum client for WebSocket connections
func NewEthWsClient(ws string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(ws)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// WaitForTx waits for a transaction to be mined
func WaitForTx(client *ethclient.Client, tx common.Hash) error {
	for i := 0; i < 30; i++ {
		time.Sleep(2 * time.Second)
		_, isPending, err := client.TransactionByHash(context.Background(), tx)
		if err != nil {
			continue
		}
		if !isPending {
			time.Sleep(2 * time.Second)
			return nil
		}
	}
	return errors.New("timeout waiting for transaction")
}

// WaitForTxReceipt waits for a transaction receipt to be available
func WaitForTxReceipt(client *ethclient.Client, tx common.Hash) (*types.Receipt, error) {
	for i := 0; i < 20; i++ {
		time.Sleep(2 * time.Second)
		txReceipt, err := client.TransactionReceipt(context.Background(), tx)
		if err != nil {
			continue
		}
		if txReceipt != nil {
			time.Sleep(2 * time.Second)
			return txReceipt, nil
		}
	}
	return nil, errors.New("timeout waiting for transaction receipt")
}

// WalletAddressFromCompressedPublicKey converts a compressed public key to an Ethereum address
func WalletAddressFromCompressedPublicKey(publicKeyStr string) (string, error) {
	pubBytes, err := hex.DecodeString(publicKeyStr)
	if err != nil {
		return "", err
	}

	pubkey, err := crypto.DecompressPubkey(pubBytes)
	if err != nil {
		return "", err
	}

	return crypto.PubkeyToAddress(*pubkey).Hex(), nil
}

// GetAccountInfo returns the private key and Ethereum address for a given private key
func GetAccountInfo(privKey string) (*ecdsa.PrivateKey, *common.Address, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, nil, err
	}

	publicKeyECDSA, _ := privateKey.Public().(*ecdsa.PublicKey)
	publicKeyAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, &publicKeyAddress, nil
}

// GenerateKeyFromSeed generates a private key, public key, and address from a seed
func GenerateKeyFromSeed(seed string) (string, string, string, error) {
	seedHex := hex.EncodeToString([]byte(seed))
	master, ch := hd.ComputeMastersFromSeed([]byte(seedHex))
	path := "m/44'/1022'/0'/0/0'"
	priv, err := hd.DerivePrivateKeyForPath(master, ch, path)
	if err != nil {
		return "", "", "", err
	}

	privateKey, err := crypto.ToECDSA(priv)
	if err != nil {
		return "", "", "", err
	}

	publicKeyECDSA, _ := privateKey.Public().(*ecdsa.PublicKey)
	privKey := hex.EncodeToString(priv)
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey := hex.EncodeToString(publicKeyBytes)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privKey, pubKey, strings.ToLower(address), nil
}

// GenerateAddress generates a new Ethereum address, public key, and private key
func GenerateAddress() (privKey, pubKey, address string, err error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "", "", "", err
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	privKey = hexutil.Encode(privateKeyBytes)[2:]

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", "", "", errors.New("failed to cast public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return privKey, pubKey, address, nil
}

// GenerateAddressFromPrivKey generates a public key and address from a private key
func GenerateAddressFromPrivKey(privKey string) (pubKey, address string, err error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", "", err
	}

	publicKeyECDSA, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return "", "", errors.New("failed to cast public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)[4:]
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return pubKey, address, nil
}

// Client wraps an Ethereum client for convenience
type Client struct {
	eth *ethclient.Client
}

// NewClient creates a new Client instance
func NewClient(eth *ethclient.Client) *Client {
	return &Client{eth}
}

// PendingNonceAt retrieves the nonce of an address
func (c *Client) PendingNonceAt(ctx context.Context, address common.Address) (uint64, error) {
	return c.eth.PendingNonceAt(ctx, address)
}

// SuggestGasPrice retrieves the suggested gas price
func (c *Client) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return c.eth.SuggestGasPrice(ctx)
}

// NetworkID retrieves the network ID
func (c *Client) NetworkID(ctx context.Context) (*big.Int, error) {
	return c.eth.NetworkID(ctx)
}

// SendTransaction sends a transaction to the network
func (c *Client) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	return c.eth.SendTransaction(ctx, tx)
}

// Transfer creates and sends a transaction from one address to another
func (c *Client) Transfer(senderPrivKey, receiverAddress string, amount, gasPrice *big.Int, gasLimit, nonce uint64) (string, error) {
	privateKey, err := crypto.HexToECDSA(senderPrivKey)
	if err != nil {
		return "", err
	}

	publicKeyECDSA, _ := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	if nonce <= 0 {
		nonce, err = c.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			return "", err
		}
	}

	if gasLimit == 0 {
		gasLimit = uint64(21000)
	}

	if gasPrice == nil {
		gasPrice, err = c.SuggestGasPrice(context.Background())
		if err != nil {
			return "", err
		}
	}

	// fee := new(big.Int).Mul(big.NewInt(int64(gasLimit)), gasPrice)
	toAddress := common.HexToAddress(receiverAddress)
	tx := types.NewTransaction(nonce, toAddress, amount, gasLimit, gasPrice, nil)

	chainID, err := c.NetworkID(context.Background())
	if err != nil {
		return "", err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return "", err
	}

	err = c.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", err
	}

	return signedTx.Hash().Hex(), nil
}
