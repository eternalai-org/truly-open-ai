// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IAI721} from "../interfaces/IAI721.sol";

abstract contract AI721Storage is IAI721 {
    mapping(uint256 nftId => TokenMetaData) internal datas;
    uint256 public nextTokenId;
    uint256 public mintPrice;
    address public royaltyReceiver;
    uint16 public royaltyPortion;

    mapping(address => bool) public isManager;
    address public workerHub;
    address public hybridModel;

    mapping(address nftOwner => uint256) internal earnedFees;
    mapping(uint256 nftId => uint256) public poolBalance;
    mapping(address nftOwner => mapping(bytes signature => bool))
        public signaturesUsed;

    mapping(uint256 agentId => bytes[]) internal missionsOf;

    uint256[47] private __gap;
}
