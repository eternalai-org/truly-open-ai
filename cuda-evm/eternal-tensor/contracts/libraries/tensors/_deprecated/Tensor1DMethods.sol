// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Float32x32, sqrt } from "./../../Float32x32/Lib32x32.sol";
import "../Tensors.sol";

library Tensor1DMethods {
    function zerosTensor(uint n) internal pure returns (Tensors.Tensor1D memory ts) {
        ts.n = n;
        ts.mat = new Float32x32[](n);
    }

    function onesTensor(uint n) internal pure returns (Tensors.Tensor1D memory ts) {
        ts.n = n;
        ts.mat = new Float32x32[](n);
        for (uint i = 0; i < n; i++) {
            ts.mat[i] = Tensors.ONE;
        }
    }

    function emptyTensor(uint n) internal pure returns (Tensors.Tensor1D memory ts) {
        ts.n = n;
    }

    function from(Float32x32[] memory mat) internal pure returns (Tensors.Tensor1D memory ts) {
        ts.n = mat.length;
        ts.mat = mat;
    }

    function count(Tensors.Tensor1D memory ts) internal pure returns (uint) {
        return ts.n;
    }

    function __apply_unary_op(
        Tensors.Tensor1D memory a,
        function(Float32x32) internal view returns (Float32x32) op
    ) internal view returns (Tensors.Tensor1D memory) {
        Tensors.Tensor1D memory res = zerosTensor(a.n);
        for (uint i = 0; i < res.n; i++) {
            res.mat[i] = op(a.mat[i]);
        }
        return res;
    }

    function __apply_binary_op(
        Tensors.Tensor1D memory a,
        Tensors.Tensor1D memory b,
        function(Float32x32, Float32x32) internal pure returns (Float32x32) op
    ) internal pure returns (Tensors.Tensor1D memory) {
        Tensors.Tensor1D memory res = zerosTensor(a.n);
        for (uint i = 0; i < res.n; i++) {
            res.mat[i] = op(a.mat[i], b.mat[i]);
        }
        return res;
    }

    function activation(Tensors.Tensor1D memory a, Tensors.ActivationFunc actv) internal view returns (Tensors.Tensor1D memory) {
        if (actv == Tensors.ActivationFunc.LeakyReLU) {
            return __apply_unary_op(a, Tensors.__leaky_relu);
        } else if (actv == Tensors.ActivationFunc.Linear) {
            return __apply_unary_op(a, Tensors.__linear);
        } else if (actv == Tensors.ActivationFunc.ReLU) {
            return __apply_unary_op(a, Tensors.__relu);
        } else if (actv == Tensors.ActivationFunc.Sigmoid) {
            return __apply_unary_op(a, Tensors.__sigmoid);
        } else if (actv == Tensors.ActivationFunc.Tanh) {
            return __apply_unary_op(a, Tensors.__tanh);
        } else if (actv == Tensors.ActivationFunc.Softmax) {
            return softmax(a);
        } else {
            revert InvalidActivationFunction();
        }
    }

    function add(Tensors.Tensor1D memory a, Tensors.Tensor1D memory b) internal pure returns (Tensors.Tensor1D memory) {
        return __apply_binary_op(a, b, Tensors.__add);
    }

    function add(Tensors.Tensor1D memory a, Float32x32 num) internal pure returns (Tensors.Tensor1D memory) {
        Tensors.Tensor1D memory res = zerosTensor(a.n);
        for (uint i = 0; i < a.n; i++) {
            res.mat[i] = a.mat[i] + num;
        }
        return res;
    }

    function mul(Tensors.Tensor1D memory a, Tensors.Tensor1D memory b) internal pure returns (Tensors.Tensor1D memory) {
        return __apply_binary_op(a, b, Tensors.__mul);
    }

    function mul(Tensors.Tensor1D memory a, Float32x32 num) internal pure returns (Tensors.Tensor1D memory) {
        Tensors.Tensor1D memory res = zerosTensor(a.n);
        for (uint i = 0; i < a.n; i++) {
            res.mat[i] = a.mat[i] * num;
        }
        return res;
    }

    function matMul(Tensors.Tensor1D memory a, Tensors.Tensor2D memory b) internal pure returns (Tensors.Tensor1D memory) {
        Tensors.Tensor1D memory res = zerosTensor(b.m);
        for (uint j = 0; j < b.m; j++) {
            for (uint k = 0; k < b.n; k++) {
                res.mat[j] = res.mat[j] + a.mat[k] * b.mat[k][j];
            }
        }
        return res;
    }

    function load(Tensors.Tensor1D memory ts, Float32x32[] memory data, uint n) internal pure {
        ts.n = n;
        ts.mat = new Float32x32[](n);
        for (uint i = 0; i < n; i++) {
            ts.mat[i] = data[i];
        }
    }

    function matMul(Tensors.Tensor1D memory a, Tensors.Tensor1D memory b) internal view returns (Tensors.Tensor1D memory) {
        if (a.n != b.n) {
            revert InvalidMatrixDimensions();
        }

        Tensors.Tensor1D memory res = zerosTensor(1);
        for (uint i = 0; i < a.n; i++) {
            res.mat[0] = res.mat[0] + a.mat[i] * b.mat[i];
        }
        return res;
    }

    function loadPartial(Tensors.Tensor1D storage ts, Float32x32[] memory data, uint ptr, uint idx) internal returns (uint, uint) {
        uint n = ts.n;
        while (idx < data.length && ptr < n) {
            ts.mat.push(data[idx]);
            ptr++;
            idx++;
        }
        return (ptr, idx);
    }

    function cloneTensor(Tensors.Tensor1D memory ts) internal pure returns (Tensors.Tensor1D memory) {
        Tensors.Tensor1D memory result;
        load(result, ts.mat, ts.n);
        return result;
    }

    function softmax(Tensors.Tensor1D memory a) internal view returns (Tensors.Tensor1D memory) {
        Tensors.Tensor1D memory res = __apply_unary_op(a, Tensors.__exp);
        Float32x32 sum_e = Float32x32.wrap(0);
        for (uint i = 0; i < res.n; i++) {
            sum_e = sum_e + res.mat[i];
        }
        for (uint i = 0; i < a.n; i++) {
            res.mat[i] = res.mat[i].div(sum_e);
        }
        return res;
    }

    function getDim(Tensors.Tensor1D memory a) internal pure returns (uint[] memory dim) {
        dim = new uint[](1);
        dim[0] = a.n;
    }

    function batchNormalize(Tensors.Tensor1D memory x,
        Tensors.Tensor1D memory movingMean,
        Tensors.Tensor1D memory movingVariance,
        Tensors.Tensor1D memory gamma,
        Tensors.Tensor1D memory beta,
        Float32x32 epsilon
    ) internal pure returns (Tensors.Tensor1D memory) {
        unchecked{
            for (uint i = 0; i < x.n; i++) {
                x.mat[i] = (x.mat[i] - movingMean.mat[i]) / sqrt(movingVariance.mat[i] + epsilon);
                x.mat[i] = x.mat[i] * gamma.mat[i] + beta.mat[i];
            }
        return x;
        } 
    }
}
