// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./../libraries/layers-new/_index.sol";

interface IFunctionalModel {
    struct LayerConfig {
        Layers.LayerType layerType;
        address layerAddress;
        uint[] inputIndices;
    }

    struct Layer {
        Layers.LayerType layerType;
        ILayer layerContract;
        uint[] inputIndices;
    }

    struct Model {
        string modelName;
        uint256 requiredWeights;
        uint256 appendedWeights;
        uint256 ptrLayer;
        Layer[] layers;
    }

    function constructModel(LayerConfig[] calldata layersConfig) external;
    function appendWeights(uint[] calldata weights, uint globalIndexStart) external;
    function predict(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor calldata output);
    function isReady() external view returns (bool);
    function getExpectedInputDim() external view returns (uint[][] memory inputDims);

    function getInfo()
        external
        view
        returns (
            string memory modelName,
            Layer[] memory layers,
            uint256 totalWeights
        );
}