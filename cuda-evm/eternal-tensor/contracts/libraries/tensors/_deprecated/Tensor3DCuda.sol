// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Float32x32 } from "./../../Float32x32/Lib32x32.sol";
import "./../Tensors.sol";
import "./Tensor3DMethods.sol";

library Tensor3DCuda {
    function conv2D_CUDA(
        Tensors.Tensor3D memory x,
        Tensors.Tensor4D memory kernel,
        Tensors.Tensor1D memory bias,
        uint[2] memory stride,
        Tensors.PaddingType padding
    ) internal view returns (Tensors.Tensor3D memory) {
        int64 params = CUDAParams.encodeConv2D(x.n, x.m, x.p, kernel.q, kernel.n, stride[0], uint(padding));
        Float32x32[][] memory t_bias = new Float32x32[][](1);
        t_bias[0] = bias.mat;
        Float32x32[][][] memory res = CUDA.conv2dFloat32x32(x.mat, kernel.mat, t_bias, params);
        return Tensor3DMethods.from(res);
    }

    function maxPooling2D_CUDA(
        Tensors.Tensor3D memory x,
        uint[2] memory size,
        uint[2] memory stride,
        Tensors.PaddingType padding
    ) internal view returns (Tensors.Tensor3D memory) {
        int64 params = CUDAParams.encodePooling2D(x.n, x.m, x.p, size[0], stride[0], uint(padding));
        Float32x32[][][] memory res = CUDA.maxPoolingFloat32x32(x.mat, params);
        return Tensor3DMethods.from(res);
    }
    
    function avgPooling2D_CUDA(
        Tensors.Tensor3D memory x,
        uint[2] memory size,
        uint[2] memory stride,
        Tensors.PaddingType padding
    ) internal view returns (Tensors.Tensor3D memory) {
        int64 params = CUDAParams.encodePooling2D(x.n, x.m, x.p,  size[0], stride[0], uint(padding));
        int outH = 1 ; int outW = 1;
        Float32x32[][][] memory res = new Float32x32[][][](1);
        (res, outH, outW) = CUDA.avgPoolingFloat32x32(x.mat, params);
        return Tensor3DMethods.from(res);
    }

    function globalAvgPooling2D_CUDA(
        Tensors.Tensor3D memory x
    ) internal view returns (Tensors.Tensor1D memory){
        Tensors.Tensor1D memory res = Tensor1DMethods.zerosTensor(x.p);

        Float32x32[][] memory sumMat = new Float32x32[][](x.m);
        for (uint i = 0; i < x.m; ++i) {
            sumMat[i] = new Float32x32[](x.p);
            for (uint j = 0; j < x.p; ++j) {
                sumMat[i][j] = Float32x32.wrap(0);
            }
        }

        for (uint i = 0; i < x.n; ++i) {
            sumMat = CUDA.addFloat32x32(sumMat, x.mat[i]);
        }

        for (uint j = 0; j < x.p; ++j) {
            Float32x32 sum = Float32x32.wrap(0);
            for (uint k = 0; k < x.m; ++k) {
                sum = sum + sumMat[k][j];
            }
            res.mat[j] = sum / (fromInt(int256(x.n*x.m)));
        }

        return res;
    }
    
    function activation_CUDA(Tensors.Tensor3D memory a, Tensors.ActivationFunc actv) internal view returns (Tensors.Tensor3D memory) {
        if (actv == Tensors.ActivationFunc.LeakyReLU) {
            return Tensor3DMethods.__apply_unary_op(a, Tensors.__leaky_relu);
        } else if (actv == Tensors.ActivationFunc.Linear) {
            return a;
        } else if (actv == Tensors.ActivationFunc.ReLU) {
            Float32x32[][][] memory res = CUDA.relu3DFloat32x32View(a.mat, 32, 32);
            return Tensor3DMethods.from(res);
        } else if (actv == Tensors.ActivationFunc.Sigmoid) {
            Float32x32[][][] memory res = CUDA.sigmoid3DFloat32x32View(a.mat, 32, 32);
            return Tensor3DMethods.from(res);
        } else if (actv == Tensors.ActivationFunc.Tanh) {
            Float32x32[][][] memory res = CUDA.tanh3DFloat32x32View(a.mat, 32, 32);
            return Tensor3DMethods.from(res);
        } else {
            revert InvalidActivationFunction();
        }
    }

    function batchNormalization_CUDA(
        Tensors.Tensor3D memory x,
        Tensors.Tensor1D memory movingMean,
        Tensors.Tensor1D memory movingVariance,
        Tensors.Tensor1D memory gamma,
        Tensors.Tensor1D memory beta,
        Float32x32 epsilon
    ) internal view returns (Tensors.Tensor3D memory) {
        int64 params = CUDAParams.encodeBatchNormalization(x.n, x.m, x.p);
        Float32x32[][] memory t_params = new Float32x32[][](4);
        t_params[0] = movingMean.mat;
        t_params[1] = movingVariance.mat;
        t_params[2] = gamma.mat;
        t_params[3] = beta.mat;
        Float32x32[][][] memory res = CUDA.batchNormFloat32x32(x.mat, t_params, epsilon, params, 32, 32);
        return Tensor3DMethods.from(res);
    }

    function concatenate_CUDA(
        Tensors.Tensor3D[] memory x,
        uint8 axis,
        Float32x32[][] memory shapes
    ) internal view returns (Tensors.Tensor3D memory) {
        uint8 len = uint8(x.length);
        Float32x32[][][][] memory mat = new Float32x32[][][][](len);
        for (uint i = 0; i < len; i++) {
            mat[i] = x[i].mat;
        }
        Float32x32[][][] memory res = CUDA.concatFloat32x32(mat, shapes, axis, 3, len, 32, 32);
        return Tensor3DMethods.from(res);
    }
}
