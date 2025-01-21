// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract FlattenLayer is ILayer {
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

        uint256[] memory data = input[0].data;
        uint64[] memory shapes = new uint64[](1);
        shapes[0] = uint64(Tensors.getElementCount(input[0].shapes));

        return TensorMethods.toTensor(data,shapes);
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        return true;
    }
}
