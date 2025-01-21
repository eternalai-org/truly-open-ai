// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "../tensors/Tensors.sol";
import "../tensors/TensorMethods.sol";

import { Float32x32 } from "./../Float32x32/Lib32x32.sol";

error TensorTypeNotSupported();
error IncorrectTensorType();
error IncorrectTensorDim();
error WeightAppendingOutOfBound();

interface ILayer {
    function appendWeights(uint256[] calldata weights) external returns (bool);
    function getWeightCount() external view returns (uint);
    function getRemainingWeightCount() external view returns (uint);
    function forward(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory);
}
