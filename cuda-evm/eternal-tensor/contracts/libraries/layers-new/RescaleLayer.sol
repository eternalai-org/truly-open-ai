// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract RescaleLayer is ILayer {
    using TensorMethods for Tensors.Tensor;

    uint public scale;
    uint public offset;

    constructor(bytes memory config) {
        (uint _scale, uint _offset) = abi.decode(
            config,
            (uint, uint)
        );
        scale = _scale;
        offset = _offset;
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
        uint len = Tensors.getElementCount(input[0].shapes);
        uint weightLen = (len + 3) / 4;

        Tensors.Tensor memory scaleTensor;
        scaleTensor.shapes = input[0].shapes;
        scaleTensor.data = new uint[](weightLen);

        Tensors.Tensor memory offsetTensor;
        offsetTensor.shapes = input[0].shapes;
        offsetTensor.data = new uint[](weightLen);

        for (uint i=0; i<len; i++) {
            uint bit = (3 - i%4) * 64;
            scaleTensor.data[i/4] |= scale << bit;
            offsetTensor.data[i/4] |= offset << bit;
        }

        Tensors.Tensor memory zt = input[0].mul(scaleTensor).add(offsetTensor);

        return zt;
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        return true;
    }
}
