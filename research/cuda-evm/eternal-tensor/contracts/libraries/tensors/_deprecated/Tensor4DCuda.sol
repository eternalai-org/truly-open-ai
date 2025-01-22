// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Float32x32 } from "./../../Float32x32/Lib32x32.sol";
import "./../Tensors.sol";
import "./Tensor4DMethods.sol";

library Tensor4DCuda {
    function avgPooling2D_CUDA(
        Tensors.Tensor4D memory x,
        uint[2] memory size,
        uint[2] memory stride,
        Tensors.PaddingType padding
    ) internal view returns (Tensors.Tensor4D memory) {
        int64 params = CUDAParams.encodePooling2D(x.m, x.p, x.q, size[0], stride[0], uint(padding));
        for (uint i = 0; i < x.n; ++i){
            int outH = 0; int outW = 0;
            (x.mat[i], outH, outW) = CUDA.avgPoolingFloat32x32(x.mat[i], params);
        }
        return x;
    }
}
