// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {IKB721, IGPUManager, IInferable} from "./interfaces/IKB721.sol";
import {ERC721URIStorage, ERC721} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import {ERC721Enumerable} from "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import {ECDSA} from "@openzeppelin/contracts/utils/cryptography/EIP712.sol";
import {IERC2981} from "@openzeppelin/contracts/interfaces/IERC2981.sol";
import {SafeERC20, IERC20} from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

/**
 * @title KB721
 * @dev Implementation of decentralized storage standard KB721.
 */
contract KB721 is ERC721Enumerable, ERC721URIStorage, IKB721 {
    /// @dev Constants
    uint256 private constant PORTION_DENOMINATOR = 10000;

    /// @dev Storage
    mapping(uint256 nftId => TokenMetaData) private _datas;
    uint256 private _nextTokenId;
    uint256 private _mintPrice;
    address private _royaltyReceiver;
    uint16 private _royaltyPortion;
    address public _gpuManager;
    IERC20 private immutable _tokenFee;

    mapping(uint256 nftId => uint256) public _poolBalance;
    mapping(address nftId => mapping(bytes32 signature => bool))
        public _signaturesUsed;

    /// @dev Modifiers
    modifier onlyAgentOwner(uint256 nftId) {
        _checkAgentOwner(msg.sender, nftId);
        _;
    }

    /// @dev constructor
    constructor(
        string memory name_,
        string memory symbol_,
        uint256 mintPrice_,
        address royaltyReceiver_,
        uint16 royaltyPortion_,
        uint256 nextTokenId_,
        address gpuManager_,
        IERC20 tokenFee_
    ) ERC721(name_, symbol_) {
        _mintPrice = mintPrice_;
        _royaltyReceiver = royaltyReceiver_;
        _royaltyPortion = royaltyPortion_;
        _nextTokenId = nextTokenId_;
        _gpuManager = gpuManager_;
        _tokenFee = tokenFee_;
    }

    function _setMintPrice(uint256 mintPrice) internal virtual {
        _mintPrice = mintPrice;

        emit MintPriceUpdate(mintPrice);
    }

    function _setRoyaltyReceiver(address royaltyReceiver_) internal virtual {
        _royaltyReceiver = royaltyReceiver_;

        emit RoyaltyReceiverUpdate(royaltyReceiver_);
    }

    function _setRoyaltyPortion(uint16 royaltyPortion_) internal virtual {
        _royaltyPortion = royaltyPortion_;

        emit RoyaltyPortionUpdate(royaltyPortion_);
    }

    function _setGPUManager(address gpuManager) internal virtual {
        if (gpuManager == address(0)) revert InvalidData();

        _gpuManager = gpuManager;
        emit GPUManagerUpdate(gpuManager);
    }

    function _mint(
        address to,
        string calldata uri,
        bytes calldata data,
        uint fee,
        uint256 agentId,
        string calldata promptKey,
        address promptScheduler,
        uint32 modelId
    ) internal virtual returns (uint256) {
        if (data.length == 0) revert InvalidAgentData();

        _safeMint(to, agentId);
        _setTokenURI(agentId, uri);

        _datas[agentId].fee = uint128(fee);
        _datas[agentId].sysPrompts[promptKey].push(data);
        _datas[agentId].isUsed = true;
        _datas[agentId].promptScheduler = promptScheduler;
        _datas[agentId].modelId = modelId;

        emit NewToken(agentId, uri, data, fee, to);

        return agentId;
    }

    function _wrapMint(
        address to,
        string calldata uri,
        bytes calldata data,
        uint fee,
        string calldata promptKey,
        address promptScheduler,
        uint32 modelId
    ) internal virtual returns (uint256) {
        SafeERC20.safeTransferFrom(
            _tokenFee,
            msg.sender,
            address(this),
            _mintPrice
        );

        while (_datas[_nextTokenId].isUsed) {
            _nextTokenId++;
        }
        uint256 agentId = _nextTokenId++;

        _mint(to, uri, data, fee, agentId, promptKey, promptScheduler, modelId);

        return agentId;
    }

    function _validateURI(string calldata uri) internal pure virtual {
        if (bytes(uri).length == 0) revert InvalidAgentData();
    }

    function updateAgentURI(
        uint256 agentId,
        string calldata uri
    ) public virtual override onlyAgentOwner(agentId) {
        _validateURI(uri);

        _setTokenURI(agentId, uri);
        emit AgentURIUpdate(agentId, uri);
    }

    function updateAgentData(
        uint256 agentId,
        bytes calldata sysPrompt,
        string calldata promptKey,
        uint256 promptIdx
    ) public virtual override onlyAgentOwner(agentId) {
        _validateAgentData(agentId, sysPrompt, promptIdx, promptKey);

        emit AgentDataUpdate(
            agentId,
            promptIdx,
            _datas[agentId].sysPrompts[promptKey][promptIdx],
            sysPrompt
        );

        _datas[agentId].sysPrompts[promptKey][promptIdx] = sysPrompt;
    }

    function updateAgentModelId(
        uint256 agentId,
        uint32 newModelId
    ) public virtual override onlyAgentOwner(agentId) {
        emit AgentModelIdUpdate(agentId, _datas[agentId].modelId, newModelId);

        _datas[agentId].modelId = newModelId;
    }

    function updateSchedulePrompt(
        uint256 agentId,
        address newPromptScheduler
    ) public virtual onlyAgentOwner(agentId) {
        emit AgentPromptSchedulerdUpdate(
            agentId,
            _datas[agentId].promptScheduler,
            newPromptScheduler
        );

        _datas[agentId].promptScheduler = newPromptScheduler;
    }

    function _checkUpdatePromptPermission(
        uint256 agentId,
        bytes calldata sysPrompt,
        uint256 promptIdx,
        uint256 randomNonce,
        bytes calldata signature
    ) internal virtual {
        address agentOwner = _ownerOf(agentId);
        (address signer, bytes32 signHash) = _recover(
            keccak256(
                abi.encode(
                    sysPrompt,
                    agentId,
                    promptIdx,
                    randomNonce,
                    address(this),
                    block.chainid
                )
            ),
            signature
        );
        if (_signaturesUsed[agentOwner][signHash]) revert SignatureUsed();
        _signaturesUsed[agentOwner][signHash] = true;

        _checkAgentOwner(signer, agentId);
    }

    function _validateAgentData(
        uint256 agentId,
        bytes calldata sysPrompt,
        uint256 promptIdx,
        string calldata promptKey
    ) internal view virtual {
        if (sysPrompt.length == 0) revert InvalidAgentData();
        uint256 len = _datas[agentId].sysPrompts[promptKey].length;
        if (promptIdx >= len) revert InvalidAgentPromptIndex();
    }

    function updateAgentDataWithSignature(
        uint256 agentId,
        bytes calldata sysPrompt,
        string calldata promptKey,
        uint256 promptIdx,
        uint256 randomNonce,
        bytes calldata signature
    ) public virtual override {
        _validateAgentData(agentId, sysPrompt, promptIdx, promptKey);
        _checkUpdatePromptPermission(
            agentId,
            sysPrompt,
            promptIdx,
            randomNonce,
            signature
        );

        emit AgentDataUpdate(
            agentId,
            promptIdx,
            _datas[agentId].sysPrompts[promptKey][promptIdx],
            sysPrompt
        );

        _datas[agentId].sysPrompts[promptKey][promptIdx] = sysPrompt;
    }

    function _checkUpdateUriPermission(
        uint256 agentId,
        string calldata uri,
        uint256 randomNonce,
        bytes calldata signature
    ) internal virtual {
        address agentOwner = _ownerOf(agentId);
        (address signer, bytes32 signHash) = _recover(
            keccak256(
                abi.encode(
                    agentId,
                    uri,
                    randomNonce,
                    address(this),
                    block.chainid
                )
            ),
            signature
        );
        if (_signaturesUsed[agentOwner][signHash]) revert SignatureUsed();
        _signaturesUsed[agentOwner][signHash] = true;
        _checkAgentOwner(signer, agentId);
    }

    function updateAgentUriWithSignature(
        uint256 agentId,
        string calldata uri,
        uint256 randomNonce,
        bytes calldata signature
    ) public virtual override {
        _validateURI(uri);

        _checkUpdateUriPermission(agentId, uri, randomNonce, signature);
        _setTokenURI(agentId, uri);
        emit AgentURIUpdate(agentId, uri);
    }

    function addNewAgentData(
        uint256 agentId,
        string calldata promptKey,
        bytes calldata sysPrompt
    ) public virtual override onlyAgentOwner(agentId) {
        if (sysPrompt.length == 0) revert InvalidAgentData();

        _datas[agentId].sysPrompts[promptKey].push(sysPrompt);

        emit AgentDataAddNew(agentId, _datas[agentId].sysPrompts[promptKey]);
    }

    function updateAgentFee(
        uint256 agentId,
        uint fee
    ) public virtual override onlyAgentOwner(agentId) {
        if (_datas[agentId].fee != fee) {
            _datas[agentId].fee = uint128(fee);
        }

        emit AgentFeeUpdate(agentId, fee);
    }

    function topUpPoolBalance(
        uint256 agentId,
        uint256 amount
    ) public virtual override {
        SafeERC20.safeTransferFrom(
            _tokenFee,
            msg.sender,
            address(this),
            amount
        );
        _poolBalance[agentId] += amount;

        emit TopUpPoolBalance(agentId, msg.sender, amount);
    }

    function getAgentFee(
        uint256 agentId
    ) public view virtual returns (uint256) {
        return _datas[agentId].fee;
    }

    function getAgentSystemPrompt(
        uint256 agentId,
        string calldata promptKey
    ) public view virtual returns (bytes[] memory) {
        return _datas[agentId].sysPrompts[promptKey];
    }

    function infer(
        uint256 agentId,
        bytes calldata fwdCalldata,
        string calldata externalData,
        string calldata promptKey,
        bool flag,
        uint feeAmount
    ) public virtual override {
        (, bytes memory fwdData) = _infer(
            agentId,
            fwdCalldata,
            promptKey,
            feeAmount
        );

        uint256 inferId = IInferable(_datas[agentId].promptScheduler).infer(
            _datas[agentId].modelId,
            fwdData,
            msg.sender,
            flag
        );

        emit InferencePerformed(
            agentId,
            msg.sender,
            fwdData,
            _datas[agentId].fee,
            externalData,
            inferId
        );
    }

    function infer(
        uint256 agentId,
        bytes calldata fwdCalldata,
        string calldata externalData,
        string calldata promptKey,
        uint256 feeAmount
    ) public virtual override {
        (, bytes memory fwdData) = _infer(
            agentId,
            fwdCalldata,
            promptKey,
            feeAmount
        );

        uint256 inferId = IInferable(_datas[agentId].promptScheduler).infer(
            _datas[agentId].modelId,
            fwdData,
            msg.sender
        );

        emit InferencePerformed(
            agentId,
            msg.sender,
            fwdData,
            _datas[agentId].fee,
            externalData,
            inferId
        );
    }

    function _infer(
        uint256 agentId,
        bytes calldata fwdCalldata,
        string calldata promptKey,
        uint256 feeAmount
    ) internal virtual returns (uint256, bytes memory) {
        if (_datas[agentId].sysPrompts[promptKey].length == 0)
            revert InvalidAgentData();
        if (feeAmount < _datas[agentId].fee) revert InvalidAgentFee();
        SafeERC20.safeTransferFrom(
            _tokenFee,
            msg.sender,
            address(this),
            feeAmount
        );

        bytes memory fwdData = abi.encodePacked(
            _concatSystemPrompts(_datas[agentId].sysPrompts[promptKey]),
            fwdCalldata
        );
        uint256 estFeeWH = IGPUManager(_gpuManager).getMinFeeToUse(
            _datas[agentId].modelId
        );

        if (feeAmount < estFeeWH && _poolBalance[agentId] >= estFeeWH) {
            unchecked {
                _poolBalance[agentId] -= estFeeWH;
            }

            if (feeAmount > 0) {
                SafeERC20.safeTransfer(_tokenFee, _ownerOf(agentId), feeAmount);
            }
        } else if (feeAmount >= estFeeWH) {
            uint256 remain = feeAmount - estFeeWH;
            if (remain > 0) {
                SafeERC20.safeTransfer(_tokenFee, _ownerOf(agentId), remain);
            }
        } else {
            revert InsufficientFunds();
        }

        SafeERC20.safeApprove(
            _tokenFee,
            _datas[agentId].promptScheduler,
            estFeeWH
        );

        return (estFeeWH, fwdData);
    }

    function dataOf(
        uint256 agentId
    ) public view virtual returns (uint128, bool) {
        return (_datas[agentId].fee, _datas[agentId].isUsed);
    }

    function royaltyInfo(
        uint256 agentId,
        uint256 salePrice
    ) public view virtual returns (address, uint256) {
        agentId;
        return (
            _royaltyReceiver,
            (salePrice * _royaltyPortion) / PORTION_DENOMINATOR
        );
    }

    function tokenURI(
        uint256 agentId
    ) public view override(ERC721, ERC721URIStorage) returns (string memory) {
        return super.tokenURI(agentId);
    }

    function _checkAgentOwner(
        address user,
        uint256 agentId
    ) internal view virtual {
        if (user != _ownerOf(agentId)) revert Unauthorized();
    }

    function getAgentIdByOwner(
        address owner
    ) external view returns (uint256[] memory) {
        uint256 len = balanceOf(owner);
        uint256[] memory agentIds = new uint256[](len);

        for (uint256 i = 0; i < len; i++) {
            agentIds[i] = tokenOfOwnerByIndex(owner, i);
        }

        return agentIds;
    }

    function nextTokenId() external view returns (uint256) {
        return _nextTokenId;
    }

    function royaltyReceiver() external view returns (address) {
        return _royaltyReceiver;
    }

    function royaltyPortion() external view returns (uint16) {
        return _royaltyPortion;
    }

    function _concatSystemPrompts(
        bytes[] memory sysPrompts
    ) internal pure virtual returns (bytes memory) {
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

    function _recover(
        bytes32 structHash,
        bytes calldata signature
    ) internal pure returns (address, bytes32) {
        bytes32 hash = ECDSA.toEthSignedMessageHash(structHash);
        return (ECDSA.recover(hash, signature), hash);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 firstTokenId,
        uint256 batchSize
    ) internal virtual override(ERC721, ERC721Enumerable) {
        super._beforeTokenTransfer(from, to, firstTokenId, batchSize);
    }

    function _burn(
        uint256 agentId
    ) internal override(ERC721, ERC721URIStorage) {
        super._burn(agentId);
    }

    function supportsInterface(
        bytes4 interfaceId
    ) public view override(ERC721Enumerable, ERC721URIStorage) returns (bool) {
        return
            interfaceId == type(IERC2981).interfaceId ||
            super.supportsInterface(interfaceId);
    }
}
