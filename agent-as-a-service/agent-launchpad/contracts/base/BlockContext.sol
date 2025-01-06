// SPDX-License-Identifier: GPL-3.0-or-later
pragma solidity ^0.8.0;

abstract contract BlockContext {
    function _blockTimestamp32() internal view virtual returns (uint32) {
        return uint32(block.timestamp);
    }

    function _blockTimestamp() internal view virtual returns (uint256) {
        return uint256(block.timestamp);
    }

    function _blockNumber() internal view virtual returns (uint256) {
        return block.number;
    }

    function getChainId() public view virtual returns (uint256 chainId) {
        assembly {
            chainId := chainid()
        }
    }
}
