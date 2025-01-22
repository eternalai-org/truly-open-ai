// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract GlobalAveragePooling2DLayer is ILayer { // Average of each features (cols)

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
        Tensors.Tensor memory yt = input[0].global_avgpooling2d();
        return yt;
    }

    function appendWeights(uint256[] calldata weights) external returns (bool) {
        return true;
    }
}
