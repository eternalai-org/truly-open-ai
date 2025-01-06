// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;

import "../../core/interfaces/IUniswapV3Pool.sol";
import "./PoolAddress.sol";
import "./LiquidityAmounts.sol";
import "./PositionKey.sol";
import "../../core/libraries/FixedPoint128.sol";
import "../../core/libraries/TickMath.sol";

library UniswapV3Broker {
    /// @notice Emitted when liquidity is increased for a position NFT
    /// @dev Also emitted when a token is minted
    /// @param tokenId The ID of the token for which liquidity was increased
    /// @param liquidity The amount by which liquidity for the NFT position was increased
    /// @param amount0 The amount of token0 that was paid for the increase in liquidity
    /// @param amount1 The amount of token1 that was paid for the increase in liquidity
    event IncreaseLiquidity(
        uint256 indexed tokenId,
        uint128 liquidity,
        uint256 amount0,
        uint256 amount1
    );

    event DecreaseLiquidity(
        uint256 indexed tokenId,
        uint128 liquidity,
        uint256 amount0,
        uint256 amount1
    );

    event Collect(
        uint256 indexed tokenId,
        address recipient,
        uint256 amount0,
        uint256 amount1
    );

    struct MintCallbackData {
        PoolAddress.PoolKey poolKey;
        address payer;
    }

    struct AddLiquidityParams {
        address token0;
        address token1;
        uint24 fee;
        address recipient;
        int24 tickLower;
        int24 tickUpper;
        uint256 amount0Desired;
        uint256 amount1Desired;
        uint256 amount0Min;
        uint256 amount1Min;
    }

    struct IncreaseLiquidityParams {
        uint256 tokenId;
        uint256 amount0Desired;
        uint256 amount1Desired;
        uint256 amount0Min;
        uint256 amount1Min;
        uint256 deadline;
    }

    struct DecreaseLiquidityParams {
        uint256 tokenId;
        uint128 liquidity;
        uint256 amount0Min;
        uint256 amount1Min;
        uint256 deadline;
    }

    struct CollectParams {
        uint256 tokenId;
        address recipient;
        uint128 amount0Max;
        uint128 amount1Max;
    }

    // details about the uniswap position
    struct Position {
        // the nonce for permits
        uint96 nonce;
        // the address that is approved for spending this token
        address operator;
        // the ID of the pool with which this token is connected
        uint80 poolId;
        // the tick range of the position
        int24 tickLower;
        int24 tickUpper;
        // the liquidity of the position
        uint128 liquidity;
        // the fee growth of the aggregate position as of the last action on the individual position
        uint256 feeGrowthInside0LastX128;
        uint256 feeGrowthInside1LastX128;
        // how many uncollected tokens are owed to the position, as of the last computation
        uint128 tokensOwed0;
        uint128 tokensOwed1;
    }

    /// @notice Add liquidity to an initialized pool
    function addLiquidity(
        address factory,
        AddLiquidityParams memory params
    )
        public
        returns (
            uint128 liquidity,
            uint256 amount0,
            uint256 amount1,
            IUniswapV3Pool pool
        )
    {
        PoolAddress.PoolKey memory poolKey = PoolAddress.PoolKey({
            token0: params.token0,
            token1: params.token1,
            fee: params.fee
        });

        pool = IUniswapV3Pool(PoolAddress.computeAddress(factory, poolKey));

        // compute the liquidity amount
        {
            (uint160 sqrtPriceX96, , , , , , ) = pool.slot0();
            uint160 sqrtRatioAX96 = TickMath.getSqrtRatioAtTick(
                params.tickLower
            );
            uint160 sqrtRatioBX96 = TickMath.getSqrtRatioAtTick(
                params.tickUpper
            );

            liquidity = LiquidityAmounts.getLiquidityForAmounts(
                sqrtPriceX96,
                sqrtRatioAX96,
                sqrtRatioBX96,
                params.amount0Desired,
                params.amount1Desired
            );
        }

        (amount0, amount1) = pool.mint(
            params.recipient,
            params.tickLower,
            params.tickUpper,
            liquidity,
            abi.encode(MintCallbackData({poolKey: poolKey, payer: msg.sender}))
        );

        require(
            amount0 >= params.amount0Min && amount1 >= params.amount1Min,
            "Price slippage check"
        );
    }

    function increaseLiquidity(
        address factory,
        IncreaseLiquidityParams calldata params,
        UniswapV3Broker.Position memory positionArg,
        PoolAddress.PoolKey memory poolKey
    )
        public
        returns (
            UniswapV3Broker.Position memory position,
            uint128 liquidity,
            uint256 amount0,
            uint256 amount1
        )
    {
        position = positionArg;
        IUniswapV3Pool pool;
        (liquidity, amount0, amount1, pool) = addLiquidity(
            factory,
            AddLiquidityParams({
                token0: poolKey.token0,
                token1: poolKey.token1,
                fee: poolKey.fee,
                tickLower: position.tickLower,
                tickUpper: position.tickUpper,
                amount0Desired: params.amount0Desired,
                amount1Desired: params.amount1Desired,
                amount0Min: params.amount0Min,
                amount1Min: params.amount1Min,
                recipient: address(this)
            })
        );

        bytes32 positionKey = PositionKey.compute(
            address(this),
            position.tickLower,
            position.tickUpper
        );

        // this is now updated to the current transaction
        (
            ,
            uint256 feeGrowthInside0LastX128,
            uint256 feeGrowthInside1LastX128,
            ,

        ) = pool.positions(positionKey);

        position.tokensOwed0 += uint128(
            FullMath.mulDiv(
                feeGrowthInside0LastX128 - position.feeGrowthInside0LastX128,
                position.liquidity,
                FixedPoint128.Q128
            )
        );
        position.tokensOwed1 += uint128(
            FullMath.mulDiv(
                feeGrowthInside1LastX128 - position.feeGrowthInside1LastX128,
                position.liquidity,
                FixedPoint128.Q128
            )
        );

        position.feeGrowthInside0LastX128 = feeGrowthInside0LastX128;
        position.feeGrowthInside1LastX128 = feeGrowthInside1LastX128;
        position.liquidity += liquidity;

        emit IncreaseLiquidity(params.tokenId, liquidity, amount0, amount1);
    }

    function decreaseLiquidity(
        address factory,
        DecreaseLiquidityParams calldata params,
        UniswapV3Broker.Position memory positionArg,
        PoolAddress.PoolKey memory poolKey
    )
        public
        returns (
            UniswapV3Broker.Position memory position,
            uint256 amount0,
            uint256 amount1
        )
    {
        require(params.liquidity > 0);
        position = positionArg;

        uint128 positionLiquidity = position.liquidity;
        require(positionLiquidity >= params.liquidity);

        IUniswapV3Pool pool = IUniswapV3Pool(
            PoolAddress.computeAddress(factory, poolKey)
        );
        (amount0, amount1) = pool.burn(
            position.tickLower,
            position.tickUpper,
            params.liquidity
        );

        require(
            amount0 >= params.amount0Min && amount1 >= params.amount1Min,
            "Price slippage check"
        );

        bytes32 positionKey = PositionKey.compute(
            address(this),
            position.tickLower,
            position.tickUpper
        );
        // this is now updated to the current transaction
        (
            ,
            uint256 feeGrowthInside0LastX128,
            uint256 feeGrowthInside1LastX128,
            ,

        ) = pool.positions(positionKey);

        position.tokensOwed0 +=
            uint128(amount0) +
            uint128(
                FullMath.mulDiv(
                    feeGrowthInside0LastX128 -
                        position.feeGrowthInside0LastX128,
                    positionLiquidity,
                    FixedPoint128.Q128
                )
            );
        position.tokensOwed1 +=
            uint128(amount1) +
            uint128(
                FullMath.mulDiv(
                    feeGrowthInside1LastX128 -
                        position.feeGrowthInside1LastX128,
                    positionLiquidity,
                    FixedPoint128.Q128
                )
            );

        position.feeGrowthInside0LastX128 = feeGrowthInside0LastX128;
        position.feeGrowthInside1LastX128 = feeGrowthInside1LastX128;
        // subtraction is safe because we checked positionLiquidity is gte params.liquidity
        position.liquidity = positionLiquidity - params.liquidity;

        emit DecreaseLiquidity(
            params.tokenId,
            params.liquidity,
            amount0,
            amount1
        );
    }

    function collect(
        address factory,
        CollectParams calldata params,
        UniswapV3Broker.Position memory positionArg,
        PoolAddress.PoolKey memory poolKey
    )
        internal
        returns (
            UniswapV3Broker.Position memory position,
            address recipient,
            uint256 amount0,
            uint256 amount1
        )
    {
        require(params.amount0Max > 0 || params.amount1Max > 0);
        // allow collecting to the nft position manager address with address 0
        recipient = params.recipient == address(0)
            ? address(this)
            : params.recipient;

        position = positionArg;

        IUniswapV3Pool pool = IUniswapV3Pool(
            PoolAddress.computeAddress(factory, poolKey)
        );

        (uint128 tokensOwed0, uint128 tokensOwed1) = (
            position.tokensOwed0,
            position.tokensOwed1
        );

        // trigger an update of the position fees owed and fee growth snapshots if it has any liquidity
        if (position.liquidity > 0) {
            pool.burn(position.tickLower, position.tickUpper, 0);
            (
                ,
                uint256 feeGrowthInside0LastX128,
                uint256 feeGrowthInside1LastX128,
                ,

            ) = pool.positions(
                    PositionKey.compute(
                        address(this),
                        position.tickLower,
                        position.tickUpper
                    )
                );

            tokensOwed0 += uint128(
                FullMath.mulDiv(
                    feeGrowthInside0LastX128 -
                        position.feeGrowthInside0LastX128,
                    position.liquidity,
                    FixedPoint128.Q128
                )
            );
            tokensOwed1 += uint128(
                FullMath.mulDiv(
                    feeGrowthInside1LastX128 -
                        position.feeGrowthInside1LastX128,
                    position.liquidity,
                    FixedPoint128.Q128
                )
            );

            position.feeGrowthInside0LastX128 = feeGrowthInside0LastX128;
            position.feeGrowthInside1LastX128 = feeGrowthInside1LastX128;
        }

        // compute the arguments to give to the pool#collect method
        (uint128 amount0Collect, uint128 amount1Collect) = (
            params.amount0Max > tokensOwed0 ? tokensOwed0 : params.amount0Max,
            params.amount1Max > tokensOwed1 ? tokensOwed1 : params.amount1Max
        );

        // the actual amounts collected are returned
        (amount0, amount1) = pool.collect(
            recipient,
            position.tickLower,
            position.tickUpper,
            amount0Collect,
            amount1Collect
        );

        // sometimes there will be a few less wei than expected due to rounding down in core, but we just subtract the full amount expected
        // instead of the actual amount so we can burn the token
        (position.tokensOwed0, position.tokensOwed1) = (
            tokensOwed0 - amount0Collect,
            tokensOwed1 - amount1Collect
        );

        emit Collect(params.tokenId, recipient, amount0Collect, amount1Collect);
    }
}
