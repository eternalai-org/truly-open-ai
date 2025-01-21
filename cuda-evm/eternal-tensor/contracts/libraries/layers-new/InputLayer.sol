// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./Layers.sol";
import "./ILayer.sol";

contract InputLayer is ILayer {
    Layers.InputType inputType;
    uint[] inputDim;

    constructor(bytes memory config) {
        (uint8 _inputType, uint[] memory _inputDim) = abi.decode(
            config,
            (uint8, uint[])
        );
        inputDim = _inputDim;
        inputType = Layers.InputType(_inputType);    
    }

    function getWeightCount() external view returns (uint) {
        return 0;
    }

    function getRemainingWeightCount() external view returns (uint) {
        return 0;
    }

    function forward(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory) {
        return input[0];
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        return true;
    }

    function getInputDim() external view returns (uint[] memory) {
        return inputDim;
    }
}
