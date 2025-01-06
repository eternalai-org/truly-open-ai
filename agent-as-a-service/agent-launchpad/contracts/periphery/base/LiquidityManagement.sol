// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.0;

import "../../core/interfaces/IUniswapV3Factory.sol";
import "../../core/interfaces/callback/IUniswapV3MintCallback.sol";
import "../../core/libraries/TickMath.sol";

import "../libraries/PoolAddress.sol";
import "../libraries/CallbackValidation.sol";
import "../libraries/LiquidityAmounts.sol";

import "../libraries/UniswapV3Broker.sol";

import "./PeripheryPayments.sol";
import "./PeripheryImmutableState.sol";

/// @title Liquidity management functions
/// @notice Internal functions for safely managing liquidity in Uniswap V3
abstract contract LiquidityManagement is
    IUniswapV3MintCallback,
    PeripheryImmutableState,
    PeripheryPayments
{
    /// @inheritdoc IUniswapV3MintCallback
    function uniswapV3MintCallback(
        uint256 amount0Owed,
        uint256 amount1Owed,
        bytes calldata data
    ) external override {
        UniswapV3Broker.MintCallbackData memory decoded = abi.decode(
            data,
            (UniswapV3Broker.MintCallbackData)
        );
        CallbackValidation.verifyCallback(factory, decoded.poolKey);

        if (amount0Owed > 0)
            pay(decoded.poolKey.token0, decoded.payer, msg.sender, amount0Owed);
        if (amount1Owed > 0)
            pay(decoded.poolKey.token1, decoded.payer, msg.sender, amount1Owed);
    }

    /// @notice Add liquidity to an initialized pool
    function addLiquidity(
        UniswapV3Broker.AddLiquidityParams memory params
    )
        internal
        returns (
            uint128 liquidity,
            uint256 amount0,
            uint256 amount1,
            IUniswapV3Pool pool
        )
    {
        return UniswapV3Broker.addLiquidity(factory, params);
    }
}
