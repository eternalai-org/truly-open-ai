// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract MaxPooling2DLayer is ILayer {
    using TensorMethods for Tensors.Tensor;

    uint[2] public size;
    uint[2] public stride;
    uint public padding;

    constructor(bytes memory config) {
        (uint256[2] memory _size, uint256[2] memory _stride, uint8 _padding) = abi.decode(
            config,
            (uint256[2], uint256[2], uint8)
        );
        size = _size;
        stride = _stride;
        padding = _padding;
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

        Tensors.Tensor memory yt = input[0].maxpooling2d(size[0], size[1], stride[0], stride[1], padding);
        return yt;
    }

    function appendWeights(uint256[] calldata weights) external returns (bool) {
        return true;
    }

    function getSize() external view returns (uint256[2] memory){
        return size;
    }

    function getStride() external view returns (uint256[2] memory){
        return stride;
    }
}
