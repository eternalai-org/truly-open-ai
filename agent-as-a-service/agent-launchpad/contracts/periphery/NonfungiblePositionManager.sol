// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.0;

import "../core/interfaces/IUniswapV3Pool.sol";
import "../core/libraries/FixedPoint128.sol";
import "../core/libraries/FullMath.sol";

import "./interfaces/INonfungiblePositionManager.sol";
import "./interfaces/INonfungibleTokenPositionDescriptor.sol";
import "./libraries/PositionKey.sol";
import "./libraries/PoolAddress.sol";
import "./base/LiquidityManagement.sol";
import "./base/PeripheryImmutableState.sol";
import "./base/Multicall.sol";
import "./base/ERC721Permit.sol";
import "./base/PeripheryValidation.sol";
import "./base/SelfPermit.sol";
import "./base/PoolInitializer.sol";

import "./libraries/UniswapV3Broker.sol";

/// @title NFT positions
/// @notice Wraps Uniswap V3 positions in the ERC721 non-fungible token interface
contract NonfungiblePositionManager is
    INonfungiblePositionManager,
    Multicall,
    ERC721Permit,
    PeripheryImmutableState,
    PoolInitializer,
    LiquidityManagement,
    PeripheryValidation,
    SelfPermit
{
    /// @dev IDs of pools assigned by this contract
    mapping(address => uint80) private _poolIds;

    /// @dev Pool keys by pool ID, to save on SSTOREs for position data
    mapping(uint80 => PoolAddress.PoolKey) private _poolIdToPoolKey;

    /// @dev The token ID position data
    mapping(uint256 => UniswapV3Broker.Position) private _positions;

    /// @dev The ID of the next token that will be minted. Skips 0
    uint176 private _nextId;
    /// @dev The ID of the next pool that is used for the first time. Skips 0
    uint80 private _nextPoolId;

    /// @dev The address of the token descriptor contract, which handles generating token URIs for position tokens
    address private _tokenDescriptor;

    function initialize(
        address _factory,
        address _WETH,
        address _tokenDescriptor_
    ) external initializer {
        __Ownable_init();
        __ERC721Permit_init("Uniswap V3 Positions NFT-V1", "UNI-V3-POS", "1");
        __PeripheryImmutableState_init(_factory, _WETH);
        //
        _tokenDescriptor = _tokenDescriptor_;
        //
        _nextId = 1;
        /// @dev The ID of the next pool that is used for the first time. Skips 0
        _nextPoolId = 1;
    }

    /// @inheritdoc INonfungiblePositionManager
    function positions(
        uint256 tokenId
    )
        external
        view
        override
        returns (
            uint96 nonce,
            address operator,
            address token0,
            address token1,
            uint24 fee,
            int24 tickLower,
            int24 tickUpper,
            uint128 liquidity,
            uint256 feeGrowthInside0LastX128,
            uint256 feeGrowthInside1LastX128,
            uint128 tokensOwed0,
            uint128 tokensOwed1
        )
    {
        UniswapV3Broker.Position memory position = _positions[tokenId];
        require(position.poolId != 0, "Invalid token ID");
        PoolAddress.PoolKey memory poolKey = _poolIdToPoolKey[position.poolId];
        return (
            position.nonce,
            position.operator,
            poolKey.token0,
            poolKey.token1,
            poolKey.fee,
            position.tickLower,
            position.tickUpper,
            position.liquidity,
            position.feeGrowthInside0LastX128,
            position.feeGrowthInside1LastX128,
            position.tokensOwed0,
            position.tokensOwed1
        );
    }

    /// @dev Caches a pool key
    function cachePoolKey(
        address pool,
        PoolAddress.PoolKey memory poolKey
    ) private returns (uint80 poolId) {
        poolId = _poolIds[pool];
        if (poolId == 0) {
            _poolIds[pool] = (poolId = _nextPoolId++);
            _poolIdToPoolKey[poolId] = poolKey;
        }
    }

    /// @inheritdoc INonfungiblePositionManager
    function mint(
        MintParams calldata params
    )
        external
        payable
        override
        checkDeadline(params.deadline)
        returns (
            uint256 tokenId,
            uint128 liquidity,
            uint256 amount0,
            uint256 amount1
        )
    {
        IUniswapV3Pool pool;
        (liquidity, amount0, amount1, pool) = addLiquidity(
            UniswapV3Broker.AddLiquidityParams({
                token0: params.token0,
                token1: params.token1,
                fee: params.fee,
                recipient: address(this),
                tickLower: params.tickLower,
                tickUpper: params.tickUpper,
                amount0Desired: params.amount0Desired,
                amount1Desired: params.amount1Desired,
                amount0Min: params.amount0Min,
                amount1Min: params.amount1Min
            })
        );

        _mint(params.recipient, (tokenId = _nextId++));

        bytes32 positionKey = PositionKey.compute(
            address(this),
            params.tickLower,
            params.tickUpper
        );
        (
            ,
            uint256 feeGrowthInside0LastX128,
            uint256 feeGrowthInside1LastX128,
            ,

        ) = pool.positions(positionKey);

        // idempotent set
        uint80 poolId = cachePoolKey(
            address(pool),
            PoolAddress.PoolKey({
                token0: params.token0,
                token1: params.token1,
                fee: params.fee
            })
        );

        _positions[tokenId] = UniswapV3Broker.Position({
            nonce: 0,
            operator: address(0),
            poolId: poolId,
            tickLower: params.tickLower,
            tickUpper: params.tickUpper,
            liquidity: liquidity,
            feeGrowthInside0LastX128: feeGrowthInside0LastX128,
            feeGrowthInside1LastX128: feeGrowthInside1LastX128,
            tokensOwed0: 0,
            tokensOwed1: 0
        });

        emit IncreaseLiquidity(tokenId, liquidity, amount0, amount1);
    }

    modifier isAuthorizedForToken(uint256 tokenId) {
        require(_isApprovedOrOwner(msg.sender, tokenId), "Not approved");
        _;
    }

    function tokenURI(
        uint256 tokenId
    )
        public
        view
        override(ERC721Upgradeable, IERC721MetadataUpgradeable)
        returns (string memory)
    {
        require(_exists(tokenId));
        return
            INonfungibleTokenPositionDescriptor(_tokenDescriptor).tokenURI(
                this,
                tokenId
            );
    }

    // save bytecode by removing implementation of unused method
    function baseURI() public pure returns (string memory) {}

    /// @inheritdoc INonfungiblePositionManager
    function increaseLiquidity(
        UniswapV3Broker.IncreaseLiquidityParams calldata params
    )
        external
        payable
        override
        checkDeadline(params.deadline)
        returns (uint128 liquidity, uint256 amount0, uint256 amount1)
    {
        UniswapV3Broker.Position storage position = _positions[params.tokenId];
        PoolAddress.PoolKey memory poolKey = _poolIdToPoolKey[position.poolId];
        (
            UniswapV3Broker.Position memory positionRes,
            uint128 liquidityRes,
            uint256 amount0Res,
            uint256 amount1Res
        ) = UniswapV3Broker.increaseLiquidity(
                factory,
                params,
                position,
                poolKey
            );
        // save position
        _positions[params.tokenId] = positionRes;
        // return
        liquidity = liquidityRes;
        amount0 = amount0Res;
        amount1 = amount1Res;
    }

    /// @inheritdoc INonfungiblePositionManager
    function decreaseLiquidity(
        UniswapV3Broker.DecreaseLiquidityParams calldata params
    )
        external
        payable
        override
        isAuthorizedForToken(params.tokenId)
        checkDeadline(params.deadline)
        returns (uint256 amount0, uint256 amount1)
    {
        UniswapV3Broker.Position storage position = _positions[params.tokenId];
        PoolAddress.PoolKey memory poolKey = _poolIdToPoolKey[position.poolId];
        (
            UniswapV3Broker.Position memory positionRes,
            uint256 amount0Res,
            uint256 amount1Res
        ) = UniswapV3Broker.decreaseLiquidity(
                factory,
                params,
                position,
                poolKey
            );
        // save position
        _positions[params.tokenId] = positionRes;
        // return
        amount0 = amount0Res;
        amount1 = amount1Res;
    }

    /// @inheritdoc INonfungiblePositionManager
    function collect(
        UniswapV3Broker.CollectParams calldata params
    )
        external
        payable
        override
        isAuthorizedForToken(params.tokenId)
        returns (uint256 amount0, uint256 amount1)
    {
        UniswapV3Broker.Position storage position = _positions[params.tokenId];
        PoolAddress.PoolKey memory poolKey = _poolIdToPoolKey[position.poolId];
        (
            UniswapV3Broker.Position memory positionRes,
            ,
            uint256 amount0Res,
            uint256 amount1Res
        ) = UniswapV3Broker.collect(factory, params, position, poolKey);
        require(params.amount0Max > 0 || params.amount1Max > 0);
        // save position
        _positions[params.tokenId] = positionRes;
        // return
        amount0 = amount0Res;
        amount1 = amount1Res;
    }

    /// @inheritdoc INonfungiblePositionManager
    function burn(
        uint256 tokenId
    ) external payable override isAuthorizedForToken(tokenId) {
        UniswapV3Broker.Position storage position = _positions[tokenId];
        require(
            position.liquidity == 0 &&
                position.tokensOwed0 == 0 &&
                position.tokensOwed1 == 0,
            "Not cleared"
        );
        delete _positions[tokenId];
        _burn(tokenId);
    }

    function _getAndIncrementNonce(
        uint256 tokenId
    ) internal override returns (uint256) {
        return uint256(_positions[tokenId].nonce++);
    }

    /// @inheritdoc IERC721Upgradeable
    function getApproved(
        uint256 tokenId
    )
        public
        view
        override(ERC721Upgradeable, IERC721Upgradeable)
        returns (address)
    {
        require(
            _exists(tokenId),
            "ERC721: approved query for nonexistent token"
        );

        return _positions[tokenId].operator;
    }

    /// @dev Overrides _approve to use the operator in the position, which is packed with the position permit nonce
    function _approve(
        address to,
        uint256 tokenId
    ) internal override(ERC721Upgradeable) {
        _positions[tokenId].operator = to;
        emit Approval(ownerOf(tokenId), to, tokenId);
    }
}
