// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Float32x32 } from "./../../Float32x32/Lib32x32.sol";
import "./../Tensors.sol";
import "./Tensor1DMethods.sol";

library Tensor1DCuda {
    function matMul_CUDA(Tensors.Tensor1D memory a, Tensors.Tensor2D memory b) internal view returns (Tensors.Tensor1D memory) {
        Float32x32[][] memory t_a = new Float32x32[][](1);
        t_a[0] = a.mat;
        Float32x32[][] memory c = CUDA.gemmFloat32x32(t_a, b.mat);        
        return Tensor1DMethods.from(c[0]);
    }

    function add_CUDA(Tensors.Tensor1D memory a, Tensors.Tensor1D memory b) internal view returns (Tensors.Tensor1D memory) {
        Float32x32[][] memory t_a = new Float32x32[][](1);
        t_a[0] = a.mat;
        Float32x32[][] memory t_b = new Float32x32[][](1);
        t_b[0] = b.mat;
        Float32x32[][] memory c = CUDA.addFloat32x32(t_a, t_b);
        return Tensor1DMethods.from(c[0]);
    }

    function mul_CUDA(Tensors.Tensor1D memory a, Tensors.Tensor1D memory b) internal view returns (Tensors.Tensor1D memory) {
        Float32x32[][] memory t_a = new Float32x32[][](1);
        t_a[0] = a.mat;
        Float32x32[][] memory t_b = new Float32x32[][](1);
        t_b[0] = b.mat;
        Float32x32[][] memory c = CUDA.mulFloat32x32(t_a, t_b);
        return Tensor1DMethods.from(c[0]);
    }

    function activation_CUDA(Tensors.Tensor1D memory a, Tensors.ActivationFunc actv) internal view returns (Tensors.Tensor1D memory) {
        Float32x32[][] memory t_a = new Float32x32[][](1);
        t_a[0] = a.mat;
        if (actv == Tensors.ActivationFunc.LeakyReLU) {
            return Tensor1DMethods.__apply_unary_op(a, Tensors.__leaky_relu);
        } else if (actv == Tensors.ActivationFunc.Linear) {
            return a;
        } else if (actv == Tensors.ActivationFunc.ReLU) {
            Float32x32[][] memory res = CUDA.reluFloat32x32(t_a);
            return Tensor1DMethods.from(res[0]);
        } else if (actv == Tensors.ActivationFunc.Sigmoid) {
            Float32x32[][] memory res = CUDA.sigmoidFloat32x32(t_a);
            return Tensor1DMethods.from(res[0]);
        } else if (actv == Tensors.ActivationFunc.Tanh) {
            Float32x32[][] memory res = CUDA.tanhFloat32x32(t_a);
            return Tensor1DMethods.from(res[0]);
        } else if (actv == Tensors.ActivationFunc.Softmax) {
            Float32x32[][] memory res = CUDA.softmaxFloat32x32(t_a);
            return Tensor1DMethods.from(res[0]);
        } else {
            revert InvalidActivationFunction();
        }
    }
}
