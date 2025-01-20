package ethapi

import (
	"math/big"

	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/helpers"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/arbitrumnonfungiblepositionmanager"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/arbitrumpool"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/bscpool"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/erc721"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/memenonfungiblepositionmanager"
	"github.com/eternalai-org/truly-open-ai/agent-as-a-service/agent-orchestration/backend/services/3rd/binds/uniswapv3pool"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type UniV3SwapPositionInfo struct {
	Nonce                    *big.Int `json:"nonce"`
	Operator                 string   `json:"opperator"`
	Token0                   string   `json:"token0"`
	Token1                   string   `json:"token1"`
	Fee                      *big.Int `json:"fee"`
	TickLower                *big.Int `json:"tick_lower"`
	TickUpper                *big.Int `json:"tick_upper"`
	Liquidity                *big.Int `json:"liquidity"`
	FeeGrowthInside0LastX128 *big.Int `json:"fee_growth_inside0_last_x128"`
	FeeGrowthInside1LastX128 *big.Int `json:"fee_growth_inside1_last_x128"`
	TokensOwed0              *big.Int `json:"tokens_owed0"`
	TokensOwed1              *big.Int `json:"tokens_owed1"`
	Owner                    string   `json:"owner"`
	NetworkID                uint64   `json:"network_id"`
}

type SwapNonfungiblePositionManagerMintParams struct {
	Token0         string   `json:"token0"`
	Token1         string   `json:"token1"`
	Fee            *big.Int `json:"fee"`
	TickLower      *big.Int `json:"tick_lower"`
	TickUpper      *big.Int `json:"tick_upper"`
	Amount0Desired *big.Int `json:"amount0_desired"`
	Amount1Desired *big.Int `json:"amount1_desired"`
	Amount0Min     *big.Int `json:"amount0_min"`
	Amount1Min     *big.Int `json:"amount1_min"`
	Deadline       *big.Int `json:"deadline"`
	SqrtPriceX96   *big.Int `json:"sqrt_price_x96"`
	NetworkID      uint64   `json:"network_id"`
}

type UniswapPositionLiquidity struct {
	TxHash          string   `json:"tx_hash"`
	ContractAddress string   `json:"contract_address"`
	Timestamp       uint64   `json:"timestamp"`
	Index           uint     `json:"log_index"`
	TxIndex         uint     `json:"tx_index"`
	BlockNumber     uint64   `json:"block_number"`
	TokenId         *big.Int `json:"token_id"`
	Liquidity       *big.Int `json:"liquidity"`
	Amount0         *big.Int `json:"amount0"`
	Amount1         *big.Int `json:"amount1"`
	NetworkID       uint64   `json:"network_id"`
}

func (c *Client) MemeNonfungiblePositionManagerPositionInfo(positionAddr string, tokenId *big.Int) (*UniV3SwapPositionInfo, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	position, err := memenonfungiblepositionmanager.NewNonfungiblePositionManager(helpers.HexToAddress(positionAddr), client)
	if err != nil {
		return nil, err
	}
	info, err := position.Positions(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, err
	}
	resp := &UniV3SwapPositionInfo{
		Nonce:                    info.Nonce,
		Operator:                 info.Operator.Hex(),
		Token0:                   info.Token0.Hex(),
		Token1:                   info.Token1.Hex(),
		Fee:                      info.Fee,
		TickLower:                info.TickLower,
		TickUpper:                info.TickUpper,
		Liquidity:                info.Liquidity,
		FeeGrowthInside0LastX128: info.FeeGrowthInside0LastX128,
		FeeGrowthInside1LastX128: info.FeeGrowthInside1LastX128,
		TokensOwed0:              info.TokensOwed0,
		TokensOwed1:              info.TokensOwed1,
	}
	nft, err := erc721.NewErc721(helpers.HexToAddress(positionAddr), client)
	if err != nil {
		return nil, err
	}
	ownerAddres, err := nft.OwnerOf(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, err
	}
	resp.Owner = ownerAddres.Hex()
	return resp, nil
}

func (c *Client) CamelotNonfungiblePositionManagerPositionInfo(positionAddr string, tokenId *big.Int) (*UniV3SwapPositionInfo, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	position, err := arbitrumnonfungiblepositionmanager.NewNonfungiblePositionManager(helpers.HexToAddress(positionAddr), client)
	if err != nil {
		return nil, err
	}
	info, err := position.Positions(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, err
	}
	resp := &UniV3SwapPositionInfo{
		Nonce:                    info.Nonce,
		Operator:                 info.Operator.Hex(),
		Token0:                   info.Token0.Hex(),
		Token1:                   info.Token1.Hex(),
		TickLower:                info.TickLower,
		TickUpper:                info.TickUpper,
		Liquidity:                info.Liquidity,
		FeeGrowthInside0LastX128: info.FeeGrowthInside0LastX128,
		FeeGrowthInside1LastX128: info.FeeGrowthInside1LastX128,
		TokensOwed0:              info.TokensOwed0,
		TokensOwed1:              info.TokensOwed1,
	}
	nft, err := erc721.NewErc721(helpers.HexToAddress(positionAddr), client)
	if err != nil {
		return nil, err
	}
	ownerAddres, err := nft.OwnerOf(&bind.CallOpts{}, tokenId)
	if err != nil {
		return nil, err
	}
	resp.Owner = ownerAddres.Hex()
	return resp, nil
}

func (c *Client) UniswapV3PoolSlot0(poolAddr string) (*struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint8
	Unlocked                   bool
}, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	pool, err := uniswapv3pool.NewUniswapv3pool(helpers.HexToAddress(poolAddr), client)
	if err != nil {
		return nil, err
	}
	resp, err := pool.Slot0(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) PancakeswapV3PoolSlot0(poolAddr string) (*struct {
	SqrtPriceX96               *big.Int
	Tick                       *big.Int
	ObservationIndex           uint16
	ObservationCardinality     uint16
	ObservationCardinalityNext uint16
	FeeProtocol                uint32
	Unlocked                   bool
}, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	pool, err := bscpool.NewPool(helpers.HexToAddress(poolAddr), client)
	if err != nil {
		return nil, err
	}
	resp, err := pool.Slot0(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *Client) UniswapV3PoolSlot0Arb(poolAddr string) (*struct {
	Price              *big.Int
	Tick               *big.Int
	FeeZto             uint16
	FeeOtz             uint16
	TimepointIndex     uint16
	CommunityFeeToken0 uint8
	CommunityFeeToken1 uint8
	Unlocked           bool
}, error) {
	client, err := c.getClient()
	if err != nil {
		return nil, err
	}
	pool, err := arbitrumpool.NewPool(helpers.HexToAddress(poolAddr), client)
	if err != nil {
		return nil, err
	}
	resp, err := pool.GlobalState(&bind.CallOpts{})
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
