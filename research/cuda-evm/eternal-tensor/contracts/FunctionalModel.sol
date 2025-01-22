// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/access/Ownable.sol";
import "./libraries/layers-new/_index.sol";
import { ICollectibleModel } from "./interfaces/ICollectibleModel.sol";
import { IFunctionalModel } from "./interfaces/IFunctionalModel.sol";

error LayerTypeNotSupported();
error DataTooMuch();
error ModelNotReady();
error GlobalIndexStartIncorrect();

contract FunctionalModel is IFunctionalModel, ICollectibleModel, Ownable {
    Model public model;
    uint modelId;

    function constructModel(LayerConfig[] calldata layersConfig) external onlyOwner {
        if (model.layers.length > 0) {
            delete model.layers;
        }
        model.ptrLayer = 0;
        model.requiredWeights = 0;
        model.appendedWeights = 0;
        for (uint256 i = 0; i < layersConfig.length; i++) {
            makeLayer(layersConfig[i]);
        }
    }

    function makeLayer(LayerConfig memory config) internal {
        Layers.LayerType layerType = config.layerType;
        Layer memory layer;
        layer.layerType = layerType;
        layer.inputIndices = config.inputIndices;
        layer.layerContract = ILayer(config.layerAddress);
        model.requiredWeights += layer.layerContract.getWeightCount();
        model.layers.push(layer);
    }

    function appendWeights(uint256[] calldata weights, uint globalIndexStart) external onlyOwner {
        if (globalIndexStart != model.appendedWeights + 1) {
            revert GlobalIndexStartIncorrect();
        }
        uint ptrLayer = model.ptrLayer;
        uint idx = 0;
        while (idx < weights.length && ptrLayer < model.layers.length) {            
            ILayer layer = model.layers[ptrLayer].layerContract;
            uint count = Tensors.min(layer.getRemainingWeightCount(), weights.length - idx);
            uint256[] memory toUpload = new uint[](count);
            for(uint i = 0; i < count; ++i) {
                toUpload[i] = weights[idx + i];
            }
            idx += count;
            bool isDone = layer.appendWeights(toUpload);
            if (isDone) ++ptrLayer;
        }
        if (idx < weights.length) {
            revert DataTooMuch();
        }
        model.ptrLayer = ptrLayer;
        model.appendedWeights += weights.length;
    }

    function predict(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory output) {
        if (!isReady()) {
            revert ModelNotReady();
        }

        uint inputIdx = 0;
        Tensors.Tensor[] memory outputs = new Tensors.Tensor[](model.layers.length);

        for(uint i = 0; i < model.layers.length; ++i) {
            Layer memory layer = model.layers[i];
            if (layer.layerType == Layers.LayerType.Input) {
                outputs[i] = input[inputIdx];
                ++inputIdx;
            } else {
                Tensors.Tensor[] memory layerInput = new Tensors.Tensor[](layer.inputIndices.length);
                for(uint j = 0; j < layer.inputIndices.length; ++j) {
                    layerInput[j] = outputs[layer.inputIndices[j]];
                }
                outputs[i] = layer.layerContract.forward(layerInput);

            }
        }
        
        output = outputs[outputs.length - 1];
    }

    function isReady() public view returns (bool) {
        return model.appendedWeights == model.requiredWeights;
    }
    
    function getInfo()
        public
        view
        returns (
            string memory,
            Layer[] memory,
            uint
        )
    {
        return (
            model.modelName,
            model.layers,
            model.requiredWeights
        );
    }

    function getExpectedInputDim() external view returns (uint[][] memory inputDims) {
        uint nInputLayer = 0;
        for(uint i = 0; i < model.layers.length; ++i) {
            if (model.layers[i].layerType == Layers.LayerType.Input) {
                ++nInputLayer;
            }
        }

        inputDims = new uint[][](nInputLayer);
        uint ptr = 0;
        for(uint i = 0; i < model.layers.length; ++i) {
            if (model.layers[i].layerType == Layers.LayerType.Input) {
                inputDims[ptr++] = InputLayer(address(model.layers[i].layerContract)).getInputDim();
            }
        }
        return inputDims;
    }

    function setModelId(uint256 _modelId) external {
        if (!isReady()) {
            revert ModelNotReady();
        }
        modelId = _modelId;
    }
}
