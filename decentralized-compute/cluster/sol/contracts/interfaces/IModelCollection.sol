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
    event NewModel(
        address indexed caller,
        address indexed owner,
        uint256 indexed modelId,
        string uri
    );
    event ModelURIUpdate(uint256 indexed modelId, string uri);
    event WEAITokenUpdate(address oldToken, address newToken);

    error AlreadyMinted();
    error Authorized();
    error FailedTransfer();
    error InsufficientFunds();
    error InvalidModel();
    error InvalidValue();
    error InvalidSignature();
    error Unauthorized();

    function version() external pure returns (string memory);
    function nextModelId() external view returns (uint256);
    function mintPrice() external view returns (uint256);
    function royaltyReceiver() external view returns (address);
    function royaltyPortion() external view returns (uint16);
    function checkModelExist(uint256 modelId) external view returns (bool);
    function isManager(address account) external view returns (bool);

    function mint(
        address to,
        string calldata uri
    ) external returns (uint256 modelId);
}
