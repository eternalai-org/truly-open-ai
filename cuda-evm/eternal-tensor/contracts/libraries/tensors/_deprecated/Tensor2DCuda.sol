// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Float32x32 } from "./../../Float32x32/Lib32x32.sol";
import "./../Tensors.sol";
import "./Tensor2DMethods.sol";

library Tensor2DCuda {
    function matMul_CUDA(Tensors.Tensor2D memory a, Tensors.Tensor2D memory b) internal view returns (Tensors.Tensor2D memory) {
        Float32x32[][] memory c = CUDA.gemmFloat32x32(a.mat, b.mat);        
        return Tensor2DMethods.from(c);
    }

    function add_CUDA(Tensors.Tensor2D memory a, Tensors.Tensor2D memory b) internal view returns (Tensors.Tensor2D memory) {
        Float32x32[][] memory c = CUDA.addFloat32x32(a.mat, b.mat);
        return Tensor2DMethods.from(c);
    }

    function mul_CUDA(Tensors.Tensor2D memory a, Tensors.Tensor2D memory b) internal view returns (Tensors.Tensor2D memory) {
        Float32x32[][] memory c = CUDA.mulFloat32x32(a.mat, b.mat);
        return Tensor2DMethods.from(c);
    }

    function activation_CUDA(Tensors.Tensor2D memory a, Tensors.ActivationFunc actv) internal view returns (Tensors.Tensor2D memory) {
        if (actv == Tensors.ActivationFunc.LeakyReLU) {
            return Tensor2DMethods.__apply_unary_op(a, Tensors.__leaky_relu);
        } else if (actv == Tensors.ActivationFunc.Linear) {
            return a;
        } else if (actv == Tensors.ActivationFunc.ReLU) {
            Float32x32[][] memory res = CUDA.reluFloat32x32(a.mat);
            return Tensor2DMethods.from(res);
        } else if (actv == Tensors.ActivationFunc.Sigmoid) {
            Float32x32[][] memory res = CUDA.sigmoidFloat32x32(a.mat);
            return Tensor2DMethods.from(res);
        } else if (actv == Tensors.ActivationFunc.Tanh) {
            Float32x32[][] memory res = CUDA.tanhFloat32x32(a.mat);
            return Tensor2DMethods.from(res);
        } else if (actv == Tensors.ActivationFunc.Softmax) {
            Float32x32[][] memory res = CUDA.softmaxFloat32x32(a.mat);
            return Tensor2DMethods.from(res);
        } else {
            revert InvalidActivationFunction();
        }
    }
}
