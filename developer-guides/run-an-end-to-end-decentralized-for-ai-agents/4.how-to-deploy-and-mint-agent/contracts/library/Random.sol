// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library Random {
    struct Randomizer {
        uint256 nonce;
    }

    function randomUint256(Randomizer storage _randomizer) internal returns (uint256) {
        uint256 newNonce = uint256(keccak256(abi.encodePacked(
            _randomizer.nonce,
            block.timestamp,
            blockhash(block.number)
        )));
        return _randomizer.nonce = newNonce;
    }
}
