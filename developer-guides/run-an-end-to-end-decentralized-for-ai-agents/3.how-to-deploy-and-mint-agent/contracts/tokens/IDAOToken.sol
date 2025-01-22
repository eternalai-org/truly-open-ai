// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IERC20Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC20/IERC20Upgradeable.sol";

interface IDAOToken is IERC20Upgradeable {
    function mintBatch(address[] memory to, uint256[] memory amount) external;
    function mint(address to, uint256 amount) external;
    function validateSupplyIncrease(
        uint256 _amount
    ) external view returns (bool);
}
