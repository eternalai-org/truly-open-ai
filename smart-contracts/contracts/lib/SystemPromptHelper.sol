// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;
import {ISystemPromptManager} from "../interfaces/ISystemPromptManager.sol";
import {ECDSAUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";
import {ERC721EnumerableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";

library SystemPromptHelper {
    function concatSystemPrompts(
        ISystemPromptManager.TokenMetaData storage self
    ) external view returns (bytes memory) {
        bytes[] memory sysPrompts = self.sysPrompts;
        uint256 len = sysPrompts.length;
        bytes memory concatedPrompt;

        for (uint256 i = 0; i < len; i++) {
            concatedPrompt = abi.encodePacked(
                concatedPrompt,
                sysPrompts[i],
                ";"
            );
        }

        return concatedPrompt;
    }

    function recover(
        uint256 _agentId,
        string calldata _uri,
        uint256 _randomNonce,
        bytes calldata _signature
    ) external view returns (address) {
        bytes32 structHash = keccak256(
            abi.encode(
                _uri,
                _agentId,
                _randomNonce,
                address(this),
                block.chainid
            )
        );
        bytes32 hash = ECDSAUpgradeable.toEthSignedMessageHash(structHash);
        return ECDSAUpgradeable.recover(hash, _signature);
    }

    function recover(
        uint256 _agentId,
        bytes calldata _sysPrompt,
        uint256 _promptIdx,
        uint256 _randomNonce,
        bytes calldata _signature
    ) external view returns (address) {
        bytes32 structHash = keccak256(
            abi.encode(
                _sysPrompt,
                _agentId,
                _promptIdx,
                _randomNonce,
                address(this),
                block.chainid
            )
        );

        bytes32 hash = ECDSAUpgradeable.toEthSignedMessageHash(structHash);
        return ECDSAUpgradeable.recover(hash, _signature);
    }
}
