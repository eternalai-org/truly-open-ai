// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {Set} from "../lib/Set.sol";

interface IModelLoadBalancer {
    struct ClusterGroup {
        uint16 anchorPoint;
        Set.Uint256Set clusterIds;
        string name;
    }

    event ClusterGroupCreated(bytes32 indexed groupId, string name);
    event ClusterGroupRemoved(bytes32 indexed groupId, string name);
    event ClusterAdded(bytes32 indexed groupId, uint256 indexed clusterId);
    event ClusterRemoved(bytes32 indexed groupId, uint256 indexed clusterId);
    event InferencePerformed(
        address indexed caller,
        uint256 indexed inferenceId,
        bytes32 indexed groupId,
        uint256 clusterId,
        bytes data
    );

    error ClusterGroupAlreadyExist(string name);
    error ClusterGroupNotFound(string name);
    error InvalidData();
    error InactiveClusterGroup();

    function infer(
        string memory groupName,
        bytes calldata data,
        bool rawFlag
    ) external returns (uint256);
}
