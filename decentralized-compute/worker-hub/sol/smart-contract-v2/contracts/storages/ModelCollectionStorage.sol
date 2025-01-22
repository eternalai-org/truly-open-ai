// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IModelCollection} from "../interfaces/IModelCollection.sol";

abstract contract ModelCollectionStorage is IModelCollection {
    uint256 public _nextModelId;
    uint256 public _mintPrice;
    address public _royaltyReceiver;
    uint16 public _royaltyPortion;

    mapping(address => bool) public _isManager;
    address public _wEAIToken;

    uint256[50] private __gap;
}
