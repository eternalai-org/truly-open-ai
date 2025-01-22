// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Float32x32} from "./Float32x32/Lib32x32.sol";

// @deprecated
library CUDAParams {
    function encodeConv2D(uint h, uint w, uint in_c, uint out_c, uint kernel_size, uint stride, uint padding) internal pure returns (int64) {
        return int64(uint64(h | (w << 12) | (in_c << 24) | (out_c << 36) | (kernel_size << 48) | (stride << 56) | (padding << 63)));
    }

    function encodePooling2D(uint h, uint w, uint c, uint pool_size, uint stride, uint padding) internal pure returns (int64) {
        return int64(uint64(h | (w << 12) | (c << 24) | (pool_size << 48) | (stride << 56) | (padding << 63)));
    }

    function encodeBatchNormalization(uint h, uint w, uint c) internal pure returns (int64) {
        return (int64(uint64(h | (w << 12) | (c << 24))));
    }

    function encodeAxis(uint axis) internal pure returns (int64) {
        return (int64(uint64(axis)));
    }
}