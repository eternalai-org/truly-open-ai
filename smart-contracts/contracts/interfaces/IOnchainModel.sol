// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import { IModel } from "./IModel.sol";

interface IOnchainModel is IModel {
    event ModelRegUpdate(address newAddress);

    event IdentifierUpdate(uint256 newValue);
    event InferenceCostUpdate(uint256 newValue);
    event NameUpdate(string newValue);
    event ImplementationUpdate(address newAddress);
    event InferResult(bytes result);

    error InsufficientFunds();
    error ModelIdAlreadySet();
    error ModelNotReady();

    function version() external pure returns (string memory version);

    function identifier() external view returns (uint256 identifier);
    function inferenceCost() external view returns (uint256 inferenceCost);
    function name() external view returns (string memory name);
    function implementation() external view returns (address implementation);

    function infer(bytes calldata _data) external payable;
}
