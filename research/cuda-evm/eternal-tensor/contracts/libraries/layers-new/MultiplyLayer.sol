// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract MultiplyLayer is ILayer {
    using TensorMethods for Tensors.Tensor;

    constructor(bytes memory config) {
    }

    function getWeightCount() external view returns (uint) {
        return 0;
    }

    function getRemainingWeightCount() external view returns (uint) {
        return 0;
    }

    function isSameShape(uint64[] memory dim1, uint64[] memory dim2) internal pure returns (bool) {
        if (dim1.length != dim2.length) return false;
        uint n = dim1.length;
        for (uint i = 0; i < n; i++)
            if (dim1[i] != dim2[i]) return false;
        return true;
    }

    function forward(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory) {
        if (input[0].shapes.length == 0) {
            revert IncorrectTensorType();
        }
        Tensors.Tensor memory zt = input[0];

        for (uint i = 1; i < input.length; i++) {
            if (!isSameShape(input[0].shapes, input[i].shapes)) 
                revert IncorrectTensorDim();
            zt = zt.mul(input[i]);
        }

        return zt;
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        return true;
    }
}
