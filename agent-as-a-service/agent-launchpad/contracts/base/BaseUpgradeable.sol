// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {BaseMulticall} from "./BaseMulticall.sol";
import {BlockContext} from "./BlockContext.sol";

contract BaseUpgradeable is
    BlockContext,
    BaseMulticall,
    OwnableUpgradeable,
    ReentrancyGuardUpgradeable
{
    function __BaseUpgradeable_init() internal {
        __Ownable_init();
        __ReentrancyGuard_init();
    }
}
