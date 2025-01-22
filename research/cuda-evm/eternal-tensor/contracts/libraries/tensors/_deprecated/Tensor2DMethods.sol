// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Float32x32, sqrt } from "./../../Float32x32/Lib32x32.sol";
import "./../Tensors.sol";
import "./Tensor1DMethods.sol";

library Tensor2DMethods {
    function zerosTensor(uint n, uint m) internal pure returns (Tensors.Tensor2D memory ts) {
        ts.n = n;
        ts.m = m;
        ts.mat = new Float32x32[][](n);
        for (uint i = 0; i < n; i++) {
            ts.mat[i] = new Float32x32[](m);
        }
    }
    
    function onesTensor(uint n, uint m) internal pure returns (Tensors.Tensor2D memory ts) {
        ts.n = n;
        ts.m = m;
        ts.mat = new Float32x32[][](n);
        for (uint i = 0; i < n; i++) {
            ts.mat[i] = new Float32x32[](m);
            for (uint j = 0; j < m; j++) {
                ts.mat[i][j] = Tensors.ONE;
            }
        }
    }

    function emptyTensor(uint n, uint m) internal pure returns (Tensors.Tensor2D memory ts) {
        ts.n = n;
        ts.m = m;
        ts.mat = new Float32x32[][](n);
    }
    
    function from(Float32x32[][] memory mat) internal pure returns (Tensors.Tensor2D memory ts) {
        ts.n = mat.length;
        ts.m = mat[0].length;
        ts.mat = mat;
    }

    function flat(Tensors.Tensor2D memory ts) internal pure returns (Tensors.Tensor1D memory) {
        return Tensor1DMethods.from(flat(ts.mat));
    }

    function flat(Float32x32[][] memory mat) internal pure returns (Float32x32[] memory) {
        Float32x32[] memory result = new Float32x32[](mat.length * mat[0].length);
        uint ptr = 0;
        for (uint i = 0; i < mat.length; i++) {
            for (uint j = 0; j < mat[i].length; j++) {
                result[ptr] = mat[i][j];
                ptr += 1;
            }
        }
        return result;
    }

    function load(Tensors.Tensor2D memory ts, Float32x32[] memory data, uint n, uint m) internal pure {
        ts.n = n;
        ts.m = m;
        ts.mat = new Float32x32[][](n);

        uint ptr = 0;
        for (uint i = 0; i < n; i++) {
            ts.mat[i] = new Float32x32[](m);
            for (uint j = 0; j < m; j++) {
                ts.mat[i][j] = ptr < data.length ? data[ptr] : Float32x32.wrap(0);
                ptr += 1;
            }
        }
    }

    function count(Tensors.Tensor2D memory ts) internal pure returns (uint) {
        return ts.n * ts.m;
    }

    function loadPartial(Tensors.Tensor2D storage ts, Float32x32[] memory data, uint ptr, uint idx) internal returns (uint, uint) {
        uint m = ts.m;
        uint cnt = count(ts);
        while (idx < data.length && ptr < cnt) {
            ts.mat[ptr / m].push(data[idx]);
            ptr++;
            idx++;
        }
        return (ptr, idx);
    }

    function cloneTensor(Tensors.Tensor2D memory ts) internal pure returns (Tensors.Tensor2D memory) {
        Tensors.Tensor2D memory result;
        load(result, flat(ts.mat), ts.n, ts.m);
        return result;
    }

    function __apply_unary_op(
        Tensors.Tensor2D memory a,
        function(Float32x32) internal view returns (Float32x32) op
    ) internal view returns (Tensors.Tensor2D memory) {
        Tensors.Tensor2D memory res = zerosTensor(a.n, a.m);
        for (uint i = 0; i < res.n; i++) {
            for (uint j = 0; j < res.m; j++) {
                res.mat[i][j] = op(a.mat[i][j]);
            }
        }
        return res;
    }

    function activation(Tensors.Tensor2D memory a, Tensors.ActivationFunc actv) internal view returns (Tensors.Tensor2D memory) {
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
        } else {
            revert InvalidActivationFunction();
        }
    }

    function __apply_binary_op(
        Tensors.Tensor2D memory a, 
        Tensors.Tensor1D memory b, 
        function(Float32x32, Float32x32) internal pure returns (Float32x32) op
    ) internal pure returns (Tensors.Tensor2D memory) {
        Tensors.Tensor2D memory res = zerosTensor(a.n, a.m);
        for (uint i = 0; i < res.n; i++) {
            for (uint j = 0; j < res.m; j++) {
                res.mat[i][j] = op(a.mat[i][j], b.mat[j]);
            }
        }
        return res;
    }

    function __apply_binary_op(
        Tensors.Tensor2D memory a, 
        Tensors.Tensor2D memory b, 
        function(Float32x32, Float32x32) internal pure returns (Float32x32) op
    ) internal pure returns (Tensors.Tensor2D memory) {
        Tensors.Tensor2D memory res = zerosTensor(a.n, a.m);
        for (uint i = 0; i < res.n; i++) {
            for (uint j = 0; j < res.m; j++) {
                res.mat[i][j] = op(a.mat[i][j], b.mat[i % b.n][j % b.m]);
            }
        }
        return res;
    }

    function mul(Tensors.Tensor2D memory a, Tensors.Tensor2D memory b) internal pure returns (Tensors.Tensor2D memory) {
        return __apply_binary_op(a, b, Tensors.__mul);
    }
    
    function mul(Tensors.Tensor2D memory a, Tensors.Tensor1D memory b) internal pure returns (Tensors.Tensor2D memory) {
        return __apply_binary_op(a, b, Tensors.__mul);
    }

    function mul(Tensors.Tensor2D memory a, Float32x32 num) internal pure returns (Tensors.Tensor2D memory) {
        Tensors.Tensor2D memory res = zerosTensor(a.n, a.m);
        for (uint i = 0; i < a.n; i++) {
            for(uint j = 0; j < a.m; j++) {
                res.mat[i][j] = a.mat[i][j] * num;
            }
        }
        return res;
    }

    function add(Tensors.Tensor2D memory a, Tensors.Tensor2D memory b) internal pure returns (Tensors.Tensor2D memory) {
        return __apply_binary_op(a, b, Tensors.__add);
    }

    function add(Tensors.Tensor2D memory a, Tensors.Tensor1D memory b) internal pure returns (Tensors.Tensor2D memory) {
        return __apply_binary_op(a, b, Tensors.__add);
    }

    function add(Tensors.Tensor2D memory a, Float32x32 num) internal pure returns (Tensors.Tensor2D memory) {
        Tensors.Tensor2D memory res = zerosTensor(a.n, a.m);
        for (uint i = 0; i < a.n; i++) {
            for(uint j = 0; j < a.m; j++) {
                res.mat[i][j] = a.mat[i][j] + num;
            }
        }
        return res;
    }

    function matMul(Tensors.Tensor2D memory a, Tensors.Tensor2D memory b) internal pure returns (Tensors.Tensor2D memory) {
        Tensors.Tensor2D memory res;
        res.n = a.n;
        res.m = b.m;
        res.mat = new Float32x32[][](res.n);

        for (uint i = 0; i < res.n; i++) {
            res.mat[i] = new Float32x32[](res.m);
            for (uint j = 0; j < res.m; j++) {
                for (uint k = 0; k < a.m; k++) {
                    res.mat[i][j] = res.mat[i][j] + a.mat[i][k].mul(b.mat[k][j]);
                }
            }
        }
        return res;
    }

    function softmax(Tensors.Tensor2D memory a) internal view returns (Tensors.Tensor2D memory) {
        Tensors.Tensor2D memory res = __apply_unary_op(a, Tensors.__exp);
        Float32x32 sum_e = Float32x32.wrap(0);
        for (uint i = 0; i < res.n; i++) {
            for (uint j = 0; j < res.m; j++) {
                sum_e = sum_e + res.mat[i][j];
            }
        }
        for (uint i = 0; i < a.n; i++) {
            for (uint j = 0; j < a.m; j++) {
                res.mat[i][j] = res.mat[i][j].div(sum_e);
            }
        }
        return res;
    }
    
    function getDim(Tensors.Tensor2D memory a) internal pure returns (uint[] memory dim) {
        dim = new uint[](2);
        dim[0] = a.n;
        dim[1] = a.m;
    }

    function batchNormalize(Tensors.Tensor2D memory x,
        Tensors.Tensor1D memory movingMean,
        Tensors.Tensor1D memory movingVariance,
        Tensors.Tensor1D memory gamma,
        Tensors.Tensor1D memory beta,
        Float32x32 epsilon
    ) internal pure returns (Tensors.Tensor2D memory) {
        unchecked{
            for (uint i = 0; i < x.m; i++) {
                for (uint j = 0; j < x.mat.length; j++) {
                    x.mat[j][i] = (x.mat[j][i] - movingMean.mat[i]) / sqrt(movingVariance.mat[i] + epsilon);
                    x.mat[j][i] = x.mat[j][i] * gamma.mat[i] + beta.mat[i];
                }
            }
        return x;
        } 
    }
}
