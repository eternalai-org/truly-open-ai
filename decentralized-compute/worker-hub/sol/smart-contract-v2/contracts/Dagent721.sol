// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {ERC721PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721PausableUpgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import {AI721Upgradeable, IERC20} from "./standardUpgradeable/AI721Upgradeable.sol";

contract Dagent721 is
    ERC721PausableUpgradeable,
    OwnableUpgradeable,
    AI721Upgradeable
{
    // storage
    mapping(uint256 nftId => bytes[]) private _missionsOf;
    uint256[50] private __gap;

    // event
    event AgentMissionAddNew(uint256 indexed agentId, bytes[] missions);
    event AgentMissionUpdate(
        uint256 indexed agentId,
        uint256 missionIndex,
        bytes oldSysMission,
        bytes newSysMission
    );

    function initialize(
        string calldata name_,
        string calldata symbol_,
        uint256 mintPrice_,
        address royaltyReceiver_,
        uint16 royaltyPortion_,
        uint256 nextTokenId_,
        address gpuManager_,
        IERC20 tokenFee_
    ) external initializer {
        __ERC721_init(name_, symbol_);
        __ERC721Pausable_init();
        __Ownable_init();

        _AI721_init(
            mintPrice_,
            royaltyReceiver_,
            royaltyPortion_,
            nextTokenId_,
            gpuManager_,
            tokenFee_
        );
    }

    function pause() external onlyOwner whenNotPaused {
        _pause();
    }

    function unpause() external onlyOwner whenPaused {
        _unpause();
    }

    function updateMintPrice(uint256 mintPrice) external onlyOwner {
        _setMintPrice(mintPrice);
    }

    function updateRoyaltyReceiver(address royaltyReceiver) external onlyOwner {
        _setRoyaltyReceiver(royaltyReceiver);
    }

    function updateRoyaltyPortion(uint16 royaltyPortion) external onlyOwner {
        _setRoyaltyPortion(royaltyPortion);
    }

    function updateGPUManager(address gpuManager) external onlyOwner {
        _setGPUManager(gpuManager);
    }

    function mint(
        address to,
        string calldata uri,
        bytes calldata data,
        uint256 fee,
        string calldata promptKey,
        address promptScheduler,
        uint32 modelId
    ) external returns (uint256) {
        return
            _wrapMint(to, uri, data, fee, promptKey, promptScheduler, modelId);
    }

    //
    function _beforeTokenTransfer(
        address _from,
        address _to,
        uint256 _agentId,
        uint256 _batchSize
    ) internal override(AI721Upgradeable, ERC721PausableUpgradeable) {
        super._beforeTokenTransfer(_from, _to, _agentId, _batchSize);
    }

    function _burn(
        uint256 agentId
    ) internal override(ERC721Upgradeable, AI721Upgradeable) {
        super._burn(agentId);
    }

    function tokenURI(
        uint256 _agentId
    )
        public
        view
        override(ERC721Upgradeable, AI721Upgradeable)
        returns (string memory)
    {
        return super.tokenURI(_agentId);
    }

    function createMission(
        uint256 agentId,
        bytes calldata missionData
    ) public onlyAgentOwner(agentId) {
        if (missionData.length == 0) revert InvalidAgentData();
        _missionsOf[agentId].push(missionData);

        emit AgentMissionAddNew(agentId, _missionsOf[agentId]);
    }

    function getMissionIdsByAgentId(
        uint256 agentId
    ) public view returns (bytes[] memory) {
        return _missionsOf[agentId];
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721Upgradeable, AI721Upgradeable) returns (bool) {
        return super.supportsInterface(interfaceId);
    }
}
