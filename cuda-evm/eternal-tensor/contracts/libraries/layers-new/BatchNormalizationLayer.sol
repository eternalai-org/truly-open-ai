// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract BatchNormalizationLayer is ILayer {
    using TensorMethods for Tensors.Tensor;

    uint public inputUnits;
    uint public momentum;
    uint public epsilon;
    Tensors.Tensor gamma;
    Tensors.Tensor beta;
    Tensors.Tensor movingMean;
    Tensors.Tensor movingVariance;
    uint _ptrLayer;
    uint _ptr;
    uint nLayer;
    uint[] remainingWeight;

    constructor(bytes memory config) {
        (uint256 _inputUnits, uint256 _momentum, uint256 _epsilon) = abi.decode(
            config,
            (uint256, uint256, uint256)
        );
        inputUnits = _inputUnits;
        momentum = _momentum;
        epsilon = _epsilon;
        
        TensorMethods.initStorageZerosTensor(gamma, Tensors.get1DShape(inputUnits));
        TensorMethods.initStorageZerosTensor(beta, Tensors.get1DShape(inputUnits));
        TensorMethods.initStorageZerosTensor(movingMean, Tensors.get1DShape(inputUnits));
        TensorMethods.initStorageZerosTensor(movingVariance, Tensors.get1DShape(inputUnits));

        nLayer = 4;
        remainingWeight = new uint[](nLayer);
        remainingWeight[0] = Tensors.getWeightCount(gamma.shapes);
        remainingWeight[1] = Tensors.getWeightCount(beta.shapes);
        remainingWeight[2] = Tensors.getWeightCount(movingMean.shapes);
        remainingWeight[3] = Tensors.getWeightCount(movingVariance.shapes);
        for(uint i = nLayer-1; i > 0; --i) {
            remainingWeight[i-1] += remainingWeight[i];
        }
    }

    function getWeightCount() external view returns (uint) {
        return Tensors.getWeightCount(gamma.shapes)*4;
    }

    function getRemainingWeightCount() external view returns (uint) {
        return (_ptrLayer == nLayer) ? 0 : remainingWeight[_ptrLayer] - _ptr;
    }

    function forward(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory) {
        if (input[0].shapes.length == 0) {
            revert IncorrectTensorType();
        }

        Tensors.Tensor memory yt = input[0].batchnorm(epsilon, momentum, gamma, beta, movingMean, movingVariance);
        return yt;
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        uint ptrLayer = _ptrLayer;
        uint ptr = _ptr;
        uint idx = 0;

        if (ptrLayer == 0) {
            (ptrLayer, ptr, idx) = TensorMethods.appendStorageTensor(gamma, x, ptrLayer, ptr, idx);
        }
        if (ptrLayer == 1) {
            (ptrLayer, ptr, idx) = TensorMethods.appendStorageTensor(beta, x, ptrLayer, ptr, idx);
        }
        if (ptrLayer == 2) {
            (ptrLayer, ptr, idx) = TensorMethods.appendStorageTensor(movingMean, x, ptrLayer, ptr, idx);
        }
        if (ptrLayer == 3) {
            (ptrLayer, ptr, idx) = TensorMethods.appendStorageTensor(movingVariance, x, ptrLayer, ptr, idx);
        }
        _ptrLayer = ptrLayer;
        _ptr = ptr;

        return (ptrLayer == 4); // Return true if all layers are done
    }

    // function getMomentum() external view returns (uint){
    //     return momentum;
    // }

    // function getEpsilon() external view returns (uint){
    //     return epsilon;
    // }

    function getGamma() external view returns (uint256[] memory, uint64[] memory){
        return (gamma.data, gamma.shapes);
    }
    
    function getBeta() external view returns (uint256[] memory, uint64[] memory){
        return (beta.data, beta.shapes);
    }
    
    function getMovingMean() external view returns (uint256[] memory, uint64[] memory){
        return (movingMean.data, movingMean.shapes);
    }
    
    function getMovingVariance() external view returns (uint256[] memory, uint64[] memory){
        return (movingVariance.data, movingVariance.shapes);
    }
}
