// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {OwnableUpgradeable} from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import {IERC165Upgradeable} from "@openzeppelin/contracts-upgradeable/interfaces/IERC165Upgradeable.sol";
import {IERC2981Upgradeable} from "@openzeppelin/contracts-upgradeable/interfaces/IERC2981Upgradeable.sol";
import {ERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/ERC721Upgradeable.sol";
import {IERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol";
import {ERC721EnumerableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721EnumerableUpgradeable.sol";
import {ERC721PausableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721PausableUpgradeable.sol";
import {ERC721URIStorageUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/ERC721URIStorageUpgradeable.sol";
import {IERC721MetadataUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol";
import {EIP712Upgradeable, ECDSAUpgradeable} from "@openzeppelin/contracts-upgradeable/utils/cryptography/EIP712Upgradeable.sol";

import {IModel} from "./interfaces/IModel.sol";
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

    receive() external payable {}

    modifier onlyManager() {
        if (msg.sender != owner() && !isManager[msg.sender])
            revert Unauthorized();
        _;
    }

    function initialize(
        string calldata _name,
        string calldata _symbol,
        uint256 _mintPrice,
        address _royaltyReceiver,
        uint16 _royaltyPortion,
        uint256 _nextModelId
    ) external initializer {
        __ERC721_init(_name, _symbol);
        __ERC721Pausable_init();
        __Ownable_init();

        mintPrice = _mintPrice;
        royaltyReceiver = _royaltyReceiver;
        royaltyPortion = _royaltyPortion;
        nextModelId = _nextModelId;

        isManager[owner()] = true;
    }

    function version() external pure returns (string memory) {
        return VERSION;
    }

    function pause() external onlyOwner whenNotPaused {
        _pause();
    }

    function unpause() external onlyOwner whenPaused {
        _unpause();
    }

    function authorizeManager(address _account) external onlyOwner {
        if (isManager[_account]) revert Authorized();
        isManager[_account] = true;
        emit ManagerAuthorization(_account);
    }

    function deauthorizeManager(address _account) external onlyOwner {
        if (!isManager[_account]) revert Unauthorized();
        isManager[_account] = false;
        emit ManagerDeauthorization(_account);
    }

    function updateMintPrice(uint256 _mintPrice) external onlyOwner {
        mintPrice = _mintPrice;
        emit MintPriceUpdate(_mintPrice);
    }

    function updateRoyaltyReceiver(
        address _royaltyReceiver
    ) external onlyOwner {
        royaltyReceiver = _royaltyReceiver;
        emit RoyaltyReceiverUpdate(_royaltyReceiver);
    }

    function updateRoyaltyPortion(uint16 _royaltyPortion) external onlyOwner {
        royaltyPortion = _royaltyPortion;
        emit RoyaltyPortionUpdate(_royaltyPortion);
    }

    function mint_(
        address _to,
        string calldata _uri,
        address _model,
        uint256 tokenId
    ) internal returns (uint256) {
        if (_model == address(0)) revert InvalidModel();
        if (msg.value < mintPrice) revert InsufficientFunds();

        _safeMint(_to, tokenId);
        _setTokenURI(tokenId, _uri);
        models[tokenId] = _model;
        IModel(_model).setModelId(tokenId);

        emit NewToken(tokenId, _uri, _model, msg.sender);

        return tokenId;
    }

    function mint(
        address _to,
        string calldata _uri,
        address _model
    ) external payable onlyManager returns (uint256) {
        while (models[nextModelId] != address(0)) {
            nextModelId++;
        }
        uint256 tokenId = nextModelId++;

        return mint_(_to, _uri, _model, tokenId);
    }

    function mintBySignature(
        address _to,
        string calldata _uri,
        address _model,
        address _manager,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) public virtual returns (uint256) {
        bytes32 hash = getHashToSign(_to, _uri, _model, _manager);

        address signer = ECDSAUpgradeable.recover(hash, v, r, s);
        if (signer != _manager || !isManager[_manager])
            revert InvalidSignature();
        while (models[nextModelId] != address(0)) {
            nextModelId++;
        }
        uint256 tokenId = nextModelId++;
        return mint_(_to, _uri, _model, tokenId);
    }

    function getHashToSign(
        address _to,
        string calldata _uri,
        address _model,
        address _manager
    ) public view virtual returns (bytes32) {
        bytes32 structHash = keccak256(abi.encode(_to, _uri, _model, _manager));

        return _hashTypedDataV4(structHash);
    }

    function withdraw(address _to, uint _value) external onlyOwner {
        (bool success, ) = _to.call{value: _value}("");
        if (!success) revert FailedTransfer();
    }

    function updateTokenURI(
        uint256 _tokenId,
        string calldata _uri
    ) external onlyOwner {
        _setTokenURI(_tokenId, _uri);
        emit TokenURIUpdate(_tokenId, _uri);
    }

    function updateTokenModel(
        uint256 _tokenId,
        address _model
    ) external onlyOwner {
        require(_model != address(0), "invalid token model");

        models[_tokenId] = _model;
        emit TokenModelUpdate(_tokenId, _model);
    }

    function modelAddressOf(uint256 _tokenId) external view returns (address) {
        return models[_tokenId];
    }

    function royaltyInfo(
        uint256 _tokenId,
        uint256 _salePrice
    ) external view returns (address, uint256) {
        _tokenId;
        return (
            royaltyReceiver,
            (_salePrice * royaltyPortion) / PORTION_DENOMINATOR
        );
    }

    function tokenURI(
        uint256 _tokenId
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
        return super.tokenURI(_tokenId);
    }

    function supportsInterface(
        bytes4 _interfaceId
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
            _interfaceId == type(IERC2981Upgradeable).interfaceId ||
            super.supportsInterface(_interfaceId);
    }

    function _beforeTokenTransfer(
        address _from,
        address _to,
        uint256 _tokenId,
        uint256 _batchSize
    )
        internal
        override(
            ERC721Upgradeable,
            ERC721EnumerableUpgradeable,
            ERC721PausableUpgradeable
        )
    {
        super._beforeTokenTransfer(_from, _to, _tokenId, _batchSize);
    }

    function _burn(
        uint256 _tokenId
    ) internal override(ERC721Upgradeable, ERC721URIStorageUpgradeable) {
        super._burn(_tokenId);
    }

    function setModelId(
        address _model,
        uint256 tokenId
    ) internal returns (uint256) {
        if (_model == address(0)) revert InvalidModel();

        models[tokenId] = _model;
        IModel(_model).setModelId(tokenId);

        return tokenId;
    }
}
