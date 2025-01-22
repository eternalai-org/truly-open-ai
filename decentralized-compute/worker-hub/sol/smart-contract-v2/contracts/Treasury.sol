// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";

contract Treasury is OwnableUpgradeable, ReentrancyGuardUpgradeable {
    address public wEAIToken;
    uint256[100] private __gap;

    event Receive(uint256 _amount, address _from);

    function initialize(address _wEAIToken) external initializer {
        require(
            _wEAIToken != address(0),
            "Treasury: wEAIToken is the zero address"
        );
        __Ownable_init();
        __ReentrancyGuard_init();

        wEAIToken = _wEAIToken;
    }

    receive() external payable {
        emit Receive(msg.value, msg.sender);
    }
}
