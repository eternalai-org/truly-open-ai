// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface ICallBack {
    function resultReceived(bytes calldata result) external;
    function resultReceived(uint originInferId, bytes calldata result) external;
}
