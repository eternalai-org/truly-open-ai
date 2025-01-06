// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import "../interfaces/IPeripheryImmutableState.sol";

/// @title Immutable state
/// @notice Immutable state used by periphery contracts
abstract contract PeripheryImmutableState is
    IPeripheryImmutableState,
    OwnableUpgradeable
{
    /// @inheritdoc IPeripheryImmutableState
    address public override factory;
    /// @inheritdoc IPeripheryImmutableState
    address public override WETH;

    function __PeripheryImmutableState_init(
        address _factory,
        address _WETH
    ) internal initializer {
        __Ownable_init();
        //
        factory = _factory;
        WETH = _WETH;
    }

    function setWETH(address WETHArg) external onlyOwner {
        WETH = WETHArg;
    }
}
