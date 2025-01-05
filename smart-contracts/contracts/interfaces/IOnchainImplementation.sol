// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IOnchainImplementation {
    error Unauthorized();

    function isReady() external returns (bool);
    function setModelInterface(address _interface) external;
    function infer(bytes calldata _data) external returns (bytes memory result);
}
