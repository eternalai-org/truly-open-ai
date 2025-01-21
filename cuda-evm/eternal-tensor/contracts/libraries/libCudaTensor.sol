// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Float32x32} from "./Float32x32/Lib32x32.sol";

library CUDA {
    address constant CUDA_ADD = 0x0000000000000000000000000000000000000021;
    address constant CUDA_SUB = 0x0000000000000000000000000000000000000022;
    address constant CUDA_MUL = 0x0000000000000000000000000000000000000023;
    address constant CUDA_DIV = 0x0000000000000000000000000000000000000024;
    address constant CUDA_GEMM = 0x0000000000000000000000000000000000000025;
    address constant CUDA_ABS = 0x0000000000000000000000000000000000000026;
    address constant CUDA_BITWISE_AND = 0x0000000000000000000000000000000000000027;
    address constant CUDA_BITWISE_NOT = 0x0000000000000000000000000000000000000028;
    address constant CUDA_BITWISE_OR = 0x0000000000000000000000000000000000000029;
    address constant CUDA_BITWISE_XOR = 0x000000000000000000000000000000000000002A;
    address constant CUDA_EXP = 0x000000000000000000000000000000000000002b;
    address constant CUDA_LOG = 0x000000000000000000000000000000000000002c;
    address constant CUDA_MAX = 0x000000000000000000000000000000000000002D;
    address constant CUDA_MIN = 0x000000000000000000000000000000000000002E;
    address constant CUDA_SQRT = 0x000000000000000000000000000000000000002F;

    address constant CUDA_LAYERNORM = 0x0000000000000000000000000000000000000049;
    address constant CUDA_SOFTMAX = 0x000000000000000000000000000000000000004A;
    address constant CUDA_SIGMOID = 0x000000000000000000000000000000000000004B;
    address constant CUDA_TANH = 0x000000000000000000000000000000000000004C;
    address constant CUDA_RELU = 0x000000000000000000000000000000000000004D;
    address constant CUDA_CONV2 = 0x000000000000000000000000000000000000004e;
    address constant CUDA_MAX_POOLING = 0x000000000000000000000000000000000000004f;
    address constant CUDA_AVG_POOLING = 0x0000000000000000000000000000000000000050;

    address constant CUDA_BATCHNORM = 0x0000000000000000000000000000000000000054;
    address constant CUDA_MERGING_CONCAT = 0x0000000000000000000000000000000000000055;

    address constant CUDA_SIGMOID_3D = 0x0000000000000000000000000000000000000051;
    address constant CUDA_TANH_3D = 0x0000000000000000000000000000000000000052;
    address constant CUDA_RELU_3D = 0x0000000000000000000000000000000000000053;

    address constant CudaProcessBytesArray = 0x0000000000000000000000000000000000000057;

    uint8 constant _16BIT_SIGN_INT_MAT_TYPE = 3;
    uint8 constant _32BIT_SIGN_INT_MAT_TYPE = 4;
    uint8 constant _32BIT_FLOAT_MAT_TYPE = 5;
    uint8 constant _64BIT_FLOAT_MAT_TYPE = 6;
    /////
    uint8 constant _64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32 = 7;

    function add(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_ADD.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function addFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_ADD.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function sub(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_SUB.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function subFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_SUB.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function mul(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_MUL.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function mulFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_MUL.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function div(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_DIV.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function divFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_DIV.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function gemm(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_GEMM.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function gemmFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_GEMM.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function abs(int[][] memory mat, uint8 matTypeInBit) internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_ABS.staticcall(abi.encode(matTypeInBit, 32, 32, mat));
        (result) = abi.decode(matData, (int[][]));
    }

    function absFloat32x32(Float32x32[][] memory mat) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_ABS.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function bitwiseAnd(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_BITWISE_AND.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function bitwiseAndFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_BITWISE_AND.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function bitwiseNot(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_BITWISE_NOT.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function bitwiseNotFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_BITWISE_NOT.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function bitwiseOr(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_BITWISE_OR.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function bitwiseOrFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_BITWISE_OR.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function bitwiseXor(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_BITWISE_XOR.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function bitwiseXorFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_BITWISE_XOR.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function exp(int[][] memory mat, uint8 matTypeInBit) internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_EXP.staticcall(abi.encode(matTypeInBit, 32, 32, mat));
        (result) = abi.decode(matData, (int[][]));
    }

    function expFloat32x32(Float32x32[][] memory mat) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_EXP.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function log(int[][] memory mat, uint8 matTypeInBit) internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_LOG.staticcall(abi.encode(matTypeInBit, 32, 32, mat));
        (result) = abi.decode(matData, (int[][]));
    }

    function logFloat32x32(Float32x32[][] memory mat) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_LOG.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function max(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_MAX.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function maxFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_MAX.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function min(int[][] memory mat1, int[][] memory mat2, uint8 matTypeInBit)
    internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_MIN.staticcall(abi.encode(matTypeInBit, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (int[][]));
    }

    function minFloat32x32(Float32x32[][] memory mat1, Float32x32[][] memory mat2)
    internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_MIN.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat1, mat2));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function sqrt(int[][] memory mat, uint8 matTypeInBit) internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_SQRT.staticcall(abi.encode(matTypeInBit, 32, 32, mat));
        (result) = abi.decode(matData, (int[][]));
    }

    function sqrtFloat32x32(Float32x32[][] memory mat) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_SQRT.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function layerNorm(int[][] memory mat, uint8 matTypeInBit, int epsilon) internal view
    returns (int[][] memory result) {
        (, bytes memory matData) = CUDA_LAYERNORM.staticcall(abi.encode(matTypeInBit, 32, 32, mat, epsilon));
        (result) = abi.decode(matData, (int[][]));
    }

    function layerNormFloat32x32(Float32x32[][] memory mat, Float32x32 epsilon) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_LAYERNORM.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat, epsilon));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function softmaxFloat32x32(Float32x32[][] memory mat) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_SOFTMAX.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function sigmoidFloat32x32(Float32x32[][] memory mat) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_SIGMOID.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function tanhFloat32x32(Float32x32[][] memory mat) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_TANH.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function reluFloat32x32(Float32x32[][] memory mat) internal view
    returns (Float32x32[][] memory result) {
        (, bytes memory matData) = CUDA_RELU.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32, mat));
        (result) = abi.decode(matData, (Float32x32[][]));
    }

    function conv2dFloat32x32(Float32x32[][][] memory mat,
        Float32x32[][][][] memory kernel, Float32x32[][] memory bias, int64 params) internal view
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_CONV2.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32,
            params, mat, kernel, bias));
        (result, , ) = abi.decode(matData, (Float32x32[][][], int, int));
    }

    function maxPoolingFloat32x32(Float32x32[][][] memory mat, int64 params) internal view
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_MAX_POOLING.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32,
            params, mat));
        (result, , ) = abi.decode(matData, (Float32x32[][][], int, int));
    }

    function avgPoolingFloat32x32(Float32x32[][][] memory mat, int64 params) internal view
    returns (Float32x32[][][] memory result, int outH, int outW) {
        (, bytes memory matData) = CUDA_AVG_POOLING.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, 32, 32,
            params, mat));
        (result, outH, outW) = abi.decode(matData, (Float32x32[][][], int, int));
    }

    function sigmoid3DFloat32x32(Float32x32[][][] memory mat, uint8 matInputTypeSize, uint8 matOutputTypeSize) internal
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_SIGMOID_3D.call(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, matInputTypeSize, matOutputTypeSize, mat));
        (result) = abi.decode(matData, (Float32x32[][][]));
    }

    function sigmoid3DFloat32x32View(Float32x32[][][] memory mat, uint8 matInputTypeSize, uint8 matOutputTypeSize) internal view
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_SIGMOID_3D.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, matInputTypeSize, matOutputTypeSize, mat));
        (result) = abi.decode(matData, (Float32x32[][][]));
    }

    function tanh3DFloat32x32(Float32x32[][][] memory mat, uint8 matInputTypeSize, uint8 matOutputTypeSize) internal
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_TANH_3D.call(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, matInputTypeSize, matOutputTypeSize, mat));
        (result) = abi.decode(matData, (Float32x32[][][]));
    }

    function tanh3DFloat32x32View(Float32x32[][][] memory mat, uint8 matInputTypeSize, uint8 matOutputTypeSize) internal view
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_TANH_3D.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, matInputTypeSize, matOutputTypeSize, mat));
        (result) = abi.decode(matData, (Float32x32[][][]));
    }

    function relu3DFloat32x32(Float32x32[][][] memory mat, uint8 matInputTypeSize, uint8 matOutputTypeSize) internal
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_RELU_3D.call(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, matInputTypeSize, matOutputTypeSize, mat));
        (result) = abi.decode(matData, (Float32x32[][][]));
    }

    function relu3DFloat32x32View(Float32x32[][][] memory mat, uint8 matInputTypeSize, uint8 matOutputTypeSize) internal view
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_RELU_3D.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, matInputTypeSize, matOutputTypeSize, mat));
        (result) = abi.decode(matData, (Float32x32[][][]));
    }

    function concatFloat32x32(Float32x32[][][][] memory mat, Float32x32[][] memory shapes, uint8 axis, uint8 nDims, uint8 n, uint8 matInputTypeSize, uint8 matOutputTypeSize) internal view
        returns (Float32x32[][][] memory result) {
            (, bytes memory matData) = CUDA_MERGING_CONCAT.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, matInputTypeSize, matOutputTypeSize, axis, nDims, n,  mat, shapes));
            (result) = abi.decode(matData, (Float32x32[][][]));
    }

    function batchNormFloat32x32(Float32x32[][][] memory mat, Float32x32[][] memory W, Float32x32 epsilon, int64 params, uint8 matInputTypeSize, uint8 matOutputTypeSize) internal view
    returns (Float32x32[][][] memory result) {
        (, bytes memory matData) = CUDA_BATCHNORM.staticcall(abi.encode(_64BIT_SIGN_INT_MAT_TYPE_FIXED_FLOAT32X32, matInputTypeSize, matOutputTypeSize, epsilon, params, mat, W));
        (result) = abi.decode(matData, (Float32x32[][][]));
    }

    function processBytesArray(bytes memory data) internal view returns (bytes memory) {
        (, bytes memory result) = CudaProcessBytesArray.staticcall(data);
        return result;
    }
}
