// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract OnesLikeLayer is ILayer {
    using TensorMethods for Tensors.Tensor;

    uint public constant ONE = 1 << 32;
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
        Tensors.Tensor memory zt;
        zt.shapes = input[0].shapes;
        uint len = Tensors.getElementCount(input[0].shapes);
        uint weightLen = (len + 3) / 4;
        zt.data = new uint[](weightLen);


        for (uint i=0; i<len; i++) {
            uint bit = (3 - i%4) * 64;
            zt.data[i/4] |= ONE << bit;
        }
        return zt;
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        return true;
    }
}
