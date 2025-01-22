// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract ConcatenateLayer is ILayer {
    using TensorMethods for Tensors.Tensor;

    int8 public axis;
    constructor(bytes memory config) {
        (int8 _axis) = abi.decode(
            config,
            (int8)
        );
        axis = _axis;
    }

    function getWeightCount() external view returns (uint) {
        return 0;
    }

    function getRemainingWeightCount() external view returns (uint) {
        return 0;
    }

    uint internal constant MASK = (1 << 64) - 1;
    function getValue(Tensors.Tensor calldata x, uint idx) internal pure returns (uint) {
        return x.data[idx/4] >> ((idx % 4) << 6) & MASK;
    }

    function isSameShape(uint64[] memory dim1, uint64[] memory dim2, uint8 dim) internal pure returns (bool) {
        if (dim1.length != dim2.length) return false;
        uint n = dim1.length;
        for (uint i = 0; i < n; i++) {
            if (i==dim) continue;
            if (dim1[i] != dim2[i]) return false;
        }
        return true;
    }

    function forward(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory) {
        if (input[0].shapes.length == 0) {
            revert IncorrectTensorType();
        } 

        uint8 dim;
        if (axis==-1) dim = uint8(input[0].shapes.length - 1);
        else if (axis >= int(input[0].shapes.length)) {
            revert IncorrectTensorDim();
        } else dim = uint8(axis);

        for (uint i=1; i<input.length; i++)
            if (!isSameShape(input[0].shapes, input[i].shapes, dim))
                revert IncorrectTensorDim();

        Tensors.Tensor memory ol = TensorMethods.concat(input, dim);
        return ol;
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        return true;
    }
}
