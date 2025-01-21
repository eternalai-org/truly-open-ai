// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract ReLULayer is ILayer {
    using TensorMethods for Tensors.Tensor;

    constructor(bytes memory config) {
    }

    function getWeightCount() external view returns (uint) {
        return 0;
    }

    function getRemainingWeightCount() external view returns (uint) {
        return 0;
    }

    function forward(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory) {
        if (input[0].shapes.length == 0) {
            revert IncorrectTensorType();
        }
        Tensors.Tensor memory zt = input[0].activation(Tensors.ActivationFunc.ReLU);
        return zt;
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        return true;
    }
}
