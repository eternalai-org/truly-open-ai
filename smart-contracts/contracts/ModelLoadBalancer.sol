// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;
import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/security/PausableUpgradeable.sol";
import {ReentrancyGuardUpgradeable} from "@openzeppelin/contracts-upgradeable/security/ReentrancyGuardUpgradeable.sol";
import {ModelLoadBalancerStorage, Set} from "./storages/ModelLoadBalancerStorage.sol";
import {TransferHelper} from "./lib/TransferHelper.sol";
import {IGPUManager} from "./interfaces/IGPUManager.sol";
import {IInferable} from "./interfaces/IInferable.sol";

contract ModelLoadBalancer is
    OwnableUpgradeable,
    PausableUpgradeable,
    ReentrancyGuardUpgradeable,
    ModelLoadBalancerStorage
{
    using Set for Set.Uint256Set;
    string private constant VERSION = "v0.0.1";

    function initialize(
        address gpuManager_,
        address promptScheduler_,
        address wEAIToken_
    ) public initializer {
        if (
            gpuManager_ == address(0) ||
            promptScheduler_ == address(0) ||
            wEAIToken_ == address(0)
        ) revert InvalidData();

        __Ownable_init();
        __Pausable_init();
        __ReentrancyGuard_init();

        _gpuManager = gpuManager_;
        _promptScheduler = promptScheduler_;
        _wEAIToken = wEAIToken_;
    }

    function version() external pure returns (string memory) {
        return VERSION;
    }

    function pause() external onlyOwner {
        _pause();
    }

    function unpause() external onlyOwner {
        _unpause();
    }

    function createGroup(
        string memory name,
        uint256[] calldata clusterIds
    ) external onlyOwner {
        bytes32 groupId = keccak256(abi.encodePacked(name));

        if (bytes(_groups[groupId].name).length != 0)
            revert ClusterGroupAlreadyExist(name);

        _groups[groupId].anchorPoint = 0;
        _groups[groupId].name = name;

        uint256 clusterLen = clusterIds.length;

        if (clusterLen > 0) {
            for (uint256 i = 0; i < clusterLen; i++) {
                uint256 clusterId = clusterIds[i];
                if (clusterId == 0) revert InvalidData();

                _groups[groupId].clusterIds.insert(clusterId);
                emit ClusterAdded(groupId, clusterId);
            }
        }

        emit ClusterGroupCreated(groupId, name);
    }

    function removeGroup(string calldata name) external onlyOwner {
        bytes32 id = _findGroup(name);
        delete _groups[id];
        emit ClusterGroupRemoved(id, name);
    }

    function addClustersToGroup(
        string calldata groupName,
        uint256[] calldata clusterIds
    ) external onlyOwner {
        uint256 clusterLen = clusterIds.length;
        if (clusterLen == 0) revert InvalidData();

        bytes32 groupId = _findGroup(groupName);

        for (uint256 i = 0; i < clusterLen; i++) {
            uint256 clusterId = clusterIds[i];
            if (clusterId > type(uint32).max) revert InvalidData();

            _groups[groupId].clusterIds.insert(clusterId);
            emit ClusterAdded(groupId, clusterId);
        }
    }

    function removeClustersFromGroup(
        string calldata groupName,
        uint256[] calldata clusterIds
    ) external onlyOwner {
        uint256 clusterLen = clusterIds.length;
        if (clusterLen == 0) revert InvalidData();

        bytes32 id = _findGroup(groupName);

        for (uint256 i = 0; i < clusterLen; i++) {
            uint256 clusterId = clusterIds[i];

            _groups[id].clusterIds.erase(clusterId);
            emit ClusterRemoved(id, clusterId);
        }
    }

    function getClusterIdsOfGroup(
        string memory name
    ) external view returns (uint256[] memory) {
        return _groups[keccak256(abi.encodePacked(name))].clusterIds.values;
    }

    function getClustersGroupInfo(
        string memory name
    ) external view returns (string memory, uint16, uint256[] memory) {
        bytes32 id = keccak256(abi.encodePacked(name));
        return (
            _groups[id].name,
            _groups[id].anchorPoint,
            _groups[id].clusterIds.values
        );
    }

    function infer(
        string calldata groupName,
        bytes calldata data,
        bool rawFlag
    ) external nonReentrant whenNotPaused returns (uint256) {
        if (data.length == 0) revert InvalidData();

        bytes32 groupId = _findGroup(groupName);
        uint256 numOfCluster = _groups[groupId].clusterIds.size();
        if (numOfCluster == 0) revert InvalidData();

        uint16 anchorPoint = _groups[groupId].anchorPoint;
        uint16 originPoint = anchorPoint;
        uint32 clusterId;

        while (true) {
            anchorPoint++;

            if (anchorPoint >= numOfCluster) anchorPoint = 0;

            clusterId = uint32(_groups[groupId].clusterIds.values[anchorPoint]);

            if (IGPUManager(_gpuManager).isActiveModel(clusterId)) {
                break;
            }

            if (anchorPoint == originPoint) {
                revert InactiveClusterGroup();
            }
        }
        _groups[groupId].anchorPoint = anchorPoint;

        uint256 fee = IGPUManager(_gpuManager).getMinFeeToUse(clusterId);
        if (fee > 0) {
            TransferHelper.safeTransferFrom(
                _wEAIToken,
                msg.sender,
                address(this),
                fee
            );
        }
        TransferHelper.safeApprove(_wEAIToken, _promptScheduler, fee);

        uint256 inferId;
        if (rawFlag) {
            inferId = IInferable(_promptScheduler).infer(
                clusterId,
                data,
                msg.sender,
                rawFlag
            );
        } else {
            inferId = IInferable(_promptScheduler).infer(
                clusterId,
                data,
                msg.sender
            );
        }

        emit InferencePerformed(msg.sender, inferId, groupId, clusterId, data);
        return inferId;
    }

    function _findGroup(string calldata name) internal view returns (bytes32) {
        bytes32 id = keccak256(abi.encodePacked(name));
        if (bytes(_groups[id].name).length == 0)
            revert ClusterGroupNotFound(name);

        return id;
    }

    function getGPUManagerAddress() external view returns (address) {
        return _gpuManager;
    }

    function getPromptSchedulerAddress() external view returns (address) {
        return _promptScheduler;
    }

    function getWEAITokenAddress() external view returns (address) {
        return _wEAIToken;
    }
}
