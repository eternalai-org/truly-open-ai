// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.0;

import './IPeripheryPayments.sol';

/// @title Periphery Payments
/// @notice Functions to ease deposits and withdrawals of TC
interface IPeripheryPaymentsWithFee is IPeripheryPayments {
    /// @notice Unwraps the contract's WETH balance and sends it to recipient as TC, with a percentage between
    /// 0 (exclusive), and 1 (inclusive) going to feeRecipient
    /// @dev The amountMinimum parameter prevents malicious contracts from stealing WETH from users.
    function unwrapWETHWithFee(
        uint256 amountMinimum,
        address recipient,
        uint256 feeBips,
        address feeRecipient
    ) external payable;

    /// @notice Transfers the full amount of a token held by this contract to recipient, with a percentage between
    /// 0 (exclusive) and 1 (inclusive) going to feeRecipient
    /// @dev The amountMinimum parameter prevents malicious contracts from stealing the token from users
    function sweepTokenWithFee(
        address token,
        uint256 amountMinimum,
        address recipient,
        uint256 feeBips,
        address feeRecipient
    ) external payable;
}
