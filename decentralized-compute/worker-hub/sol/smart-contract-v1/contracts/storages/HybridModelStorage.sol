// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IHybridModel} from "../interfaces/IHybridModel.sol";

abstract contract HybridModelStorage is IHybridModel {
    uint256 public identifier;

    string public name;
    string public metadata;

    address public workerHub;
    address public modelCollection;

    uint256[49] private __gap;
}
