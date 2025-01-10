// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {IERC165Upgradeable} from "@openzeppelin/contracts-upgradeable/interfaces/IERC165Upgradeable.sol";
import {IERC2981Upgradeable} from "@openzeppelin/contracts-upgradeable/interfaces/IERC2981Upgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import {ERC721EnumerableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import {ERC721PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721PausableUpgradeable.sol";
import {ERC721URIStorageUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {IERC721MetadataUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol";
import {EIP712Upgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";

import {TransferHelper} from "./lib/TransferHelper.sol";
import {ModelCollectionStorage} from "./storages/ModelCollectionStorage.sol";

contract ModelCollection is
    ModelCollectionStorage,
    EIP712Upgradeable,
    ERC721EnumerableUpgradeable,
    ERC721PausableUpgradeable,
    ERC721URIStorageUpgradeable,
    OwnableUpgradeable
{
    string private constant VERSION = "v0.0.1";
    uint256 private constant PORTION_DENOMINATOR = 10000;

    modifier onlyManager() {
        if (msg.sender != owner() && !_isManager[msg.sender])
            revert Unauthorized();
        _;
    }

    function initialize(
        string calldata name_,
        string calldata symbol_,
        uint256 mintPrice_,
        address royaltyReceiver_,
        uint16 royaltyPortion_,
        uint256 nextModelId_,
        address wEAIToken_
    ) external initializer {
        __ERC721_init(name_, symbol_);
        __ERC721Pausable_init();
        __Ownable_init();

        if (royaltyReceiver_ == address(0) || wEAIToken_ == address(0))
            revert InvalidValue();
        if (nextModelId_ >= type(uint32).max) revert InvalidValue();

        _mintPrice = mintPrice_;
        _royaltyReceiver = royaltyReceiver_;
        _royaltyPortion = royaltyPortion_;
        _nextModelId = nextModelId_;
        _wEAIToken = wEAIToken_;

        _isManager[owner()] = true;
    }

    receive() external payable {}

    function version() external pure returns (string memory) {
        return VERSION;
    }

    function pause() external onlyOwner whenNotPaused {
        _pause();
    }

    function unpause() external onlyOwner whenPaused {
        _unpause();
    }

    function authorizeManager(address account) external onlyOwner {
        if (_isManager[account]) revert Authorized();
        _isManager[account] = true;
        emit ManagerAuthorization(account);
    }

    function deauthorizeManager(address account) external onlyOwner {
        if (!_isManager[account]) revert Unauthorized();
        _isManager[account] = false;
        emit ManagerDeauthorization(account);
    }

    function isManager(address account) external view returns (bool) {
        return _isManager[account];
    }

    function updateWEAIToken(address newToken) external onlyOwner {
        if (newToken == address(0)) revert InvalidValue();

        emit WEAITokenUpdate(_wEAIToken, newToken);
        _wEAIToken = newToken;
    }

    function wEAIToken() external view returns (address) {
        return _wEAIToken;
    }

    function updateMintPrice(uint256 newPrice) external onlyOwner {
        _mintPrice = newPrice;
        emit MintPriceUpdate(newPrice);
    }

    function mintPrice() external view returns (uint256) {
        return _mintPrice;
    }

    function updateRoyaltyReceiver(address newReceiver) external onlyOwner {
        _royaltyReceiver = newReceiver;
        emit RoyaltyReceiverUpdate(newReceiver);
    }

    function royaltyReceiver() external view returns (address) {
        return _royaltyReceiver;
    }

    function updateRoyaltyPortion(uint16 newPortion) external onlyOwner {
        _royaltyPortion = newPortion;
        emit RoyaltyPortionUpdate(newPortion);
    }

    function royaltyPortion() external view returns (uint16) {
        return _royaltyPortion;
    }

    function mint(
        address to,
        string calldata uri
    ) external onlyManager returns (uint256) {
        uint256 modelId = _nextModelId++;

        while (_exists(modelId)) {
            modelId++;
        }
        if (modelId >= type(uint32).max) revert InvalidValue();

        if (_mintPrice > 0) {
            TransferHelper.safeTransferFrom(
                _wEAIToken,
                msg.sender,
                address(this),
                _mintPrice
            );
        }

        return _mint(to, uri, modelId);
    }

    function _mint(
        address to,
        string calldata uri,
        uint256 modelId
    ) internal returns (uint256) {
        _safeMint(to, modelId);
        _setTokenURI(modelId, uri);

        emit NewModel(msg.sender, to, modelId, uri);

        return modelId;
    }

    function nextModelId() external view returns (uint256) {
        return _nextModelId;
    }

    function checkModelExist(uint256 modelId) external view returns (bool) {
        return _exists(modelId);
    }

    function _burn(
        uint256 modelId
    ) internal override(ERC721Upgradeable, ERC721URIStorageUpgradeable) {
        super._burn(modelId);
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 modelId,
        uint256 batchSize
    )
        internal
        override(
            ERC721Upgradeable,
            ERC721EnumerableUpgradeable,
            ERC721PausableUpgradeable
        )
    {
        super._beforeTokenTransfer(from, to, modelId, batchSize);
    }

    function updateModelURI(
        uint256 modelId,
        string calldata uri
    ) external onlyOwner {
        _setTokenURI(modelId, uri);
        emit ModelURIUpdate(modelId, uri);
    }

    function tokenURI(
        uint256 modelId
    )
        public
        view
        override(
            ERC721Upgradeable,
            ERC721URIStorageUpgradeable,
            IERC721MetadataUpgradeable
        )
        returns (string memory)
    {
        return super.tokenURI(modelId);
    }

    function royaltyInfo(
        uint256 modelId,
        uint256 salePrice
    ) external view returns (address receiver, uint256 royaltyAmount) {
        modelId;

        receiver = _royaltyReceiver;
        royaltyAmount = (salePrice * _royaltyPortion) / PORTION_DENOMINATOR;
    }

    function supportsInterface(
        bytes4 interfaceId
    )
        public
        view
        override(
            ERC721Upgradeable,
            ERC721EnumerableUpgradeable,
            ERC721URIStorageUpgradeable,
            IERC165Upgradeable
        )
        returns (bool)
    {
        return
            interfaceId == type(IERC2981Upgradeable).interfaceId ||
            super.supportsInterface(interfaceId);
    }

    function withdraw(address to, uint256 value) external onlyOwner {
        (bool success, ) = to.call{value: value}("");
        if (!success) revert FailedTransfer();
    }
}
