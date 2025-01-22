// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {IDAOToken} from "./tokens/IDAOToken.sol";
import {TransferHelper} from "./library/TransferHelper.sol";

contract Treasury is OwnableUpgradeable, ReentrancyGuardUpgradeable {
    address public daoToken;
    uint256[100] private __gap;

    event Receive(uint256 _amount, address _from);

    function initialize(address _daoToken) external initializer {
        require(
            _daoToken != address(0),
            "Treasury: daoToken is the zero address"
        );
        __Ownable_init();
        __ReentrancyGuard_init();

        daoToken = _daoToken;
    }

    receive() external payable {
        emit Receive(msg.value, msg.sender);
    }
}
