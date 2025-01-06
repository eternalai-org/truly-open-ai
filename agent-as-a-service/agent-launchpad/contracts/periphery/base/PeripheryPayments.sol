// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";

import "../interfaces/IPeripheryPayments.sol";
import "../../interfaces/IWETH.sol";

import "../libraries/TransferHelper.sol";

import "./PeripheryImmutableState.sol";

abstract contract PeripheryPayments is
    IPeripheryPayments,
    PeripheryImmutableState
{
    receive() external payable {
        require(msg.sender == WETH, "Not WETH");
    }

    /// @inheritdoc IPeripheryPayments
    function unwrapWETH(
        uint256 amountMinimum,
        address recipient
    ) public payable override {
        uint256 balanceWETH = IWETH(WETH).balanceOf(address(this));
        require(balanceWETH >= amountMinimum, "Insufficient WETH");

        if (balanceWETH > 0) {
            IWETH(WETH).withdraw(balanceWETH);
            TransferHelper.safeTransferETH(recipient, balanceWETH);
        }
    }

    /// @inheritdoc IPeripheryPayments
    function sweepToken(
        address token,
        uint256 amountMinimum,
        address recipient
    ) public payable override {
        uint256 balanceToken = IERC20(token).balanceOf(address(this));
        require(balanceToken >= amountMinimum, "Insufficient token");

        if (balanceToken > 0) {
            TransferHelper.safeTransfer(token, recipient, balanceToken);
        }
    }

    /// @inheritdoc IPeripheryPayments
    function refundETH() external payable override {
        if (address(this).balance > 0)
            TransferHelper.safeTransferETH(msg.sender, address(this).balance);
    }

    /// @param token The token to pay
    /// @param payer The entity that must pay
    /// @param recipient The entity that will receive payment
    /// @param value The amount to pay
    function pay(
        address token,
        address payer,
        address recipient,
        uint256 value
    ) internal {
        if (token == WETH && address(this).balance >= value) {
            // pay with WETH
            IWETH(WETH).deposit{value: value}(); // wrap only what is needed to pay
            IWETH(WETH).transfer(recipient, value);
        } else if (payer == address(this)) {
            // pay with tokens already in the contract (for the exact input multihop case)
            TransferHelper.safeTransfer(token, recipient, value);
        } else {
            // pull payment
            TransferHelper.safeTransferFrom(token, payer, recipient, value);
        }
    }
}
