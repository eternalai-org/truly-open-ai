// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IERC2981Upgradeable} from "@openzeppelin/contracts-upgradeable/interfaces/IERC2981Upgradeable.sol";
import {IERC721Upgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol";
import {IERC721EnumerableUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721EnumerableUpgradeable.sol";
import {IERC721MetadataUpgradeable} from "@openzeppelin/contracts-upgradeable/token/ERC721/extensions/IERC721MetadataUpgradeable.sol";

interface IModelCollection is
    IERC721Upgradeable,
    IERC721MetadataUpgradeable,
    IERC721EnumerableUpgradeable,
    IERC2981Upgradeable
{
    event MintPriceUpdate(uint256 newValue);
    event RoyaltyPortionUpdate(uint16 newValue);
    event RoyaltyReceiverUpdate(address newAddress);
    event ManagerAuthorization(address indexed account);
    event ManagerDeauthorization(address indexed account);
    event NewToken(
        uint256 indexed tokenId,
        string uri,
        address model,
        address indexed minter
    );
    event TokenURIUpdate(uint256 indexed tokenId, string uri);
    event TokenModelUpdate(uint256 indexed tokenId, address model);

    error AlreadyMinted();
    error Authorized();
    error FailedTransfer();
    error InsufficientFunds();
    error InvalidModel();
    error InvalidSignature();
    error Unauthorized();

    function version() external pure returns (string memory version);
    function nextModelId() external view returns (uint256 nextModelId);
    function mintPrice() external view returns (uint256 mintPrice);
    function royaltyReceiver() external view returns (address royaltyReceiver);
    function royaltyPortion() external view returns (uint16 royaltyPortion);
    function isManager(address account) external view returns (bool isManager);

    function mint(
        address to,
        string calldata uri,
        address model
    ) external payable returns (uint256 tokenId);
    function mintBySignature(
        address to,
        string calldata uri,
        address model,
        address manager,
        uint8 v,
        bytes32 r,
        bytes32 s
    ) external returns (uint256 tokenId);
}
