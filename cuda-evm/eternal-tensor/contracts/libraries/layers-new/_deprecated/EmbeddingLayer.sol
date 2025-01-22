// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./../ILayer.sol";
import "../../Float32x32/Lib32x32.sol";
import "../../tensors/_deprecated/Tensor1DMethods.sol";
import "../../tensors/_deprecated/Tensor2DMethods.sol";
import "../../tensors/_deprecated/Tensor3DMethods.sol";
import "../../tensors/_deprecated/Tensor4DMethods.sol";
import "../../tensors/_deprecated/Tensor1DCuda.sol";
import "../../tensors/_deprecated/Tensor2DCuda.sol";
import "../../tensors/_deprecated/Tensor3DCuda.sol";
import "../../tensors/_deprecated/Tensor4DCuda.sol";

contract EmbeddingLayer {
    using Tensor1DMethods for Tensors.Tensor1D;
    using Tensor2DMethods for Tensors.Tensor2D;
    using Tensor3DMethods for Tensors.Tensor3D;
    using Tensor4DMethods for Tensors.Tensor4D;

    uint public inputDim;
    uint public outputDim;
    Tensors.Tensor2D public w;
    uint _ptrLayer;
    uint _ptr;

    constructor(bytes memory config) {
        (uint256 _inputDim, uint256 _outputDim) = abi.decode(
            config,
            (uint256, uint256)
        );
        inputDim = _inputDim;
        outputDim = _outputDim;        
        w = Tensor2DMethods.emptyTensor(_inputDim, _outputDim);
    }

    function getParamsCount() external view returns (uint) {
        return w.n * w.m;
    }

    function getEmbedding(Float32x32[][] memory embeddingMatrix, Float32x32[] memory tokens) internal pure returns (Tensors.Tensor2D memory) {
        Float32x32[][] memory embeddings = new Float32x32[][](tokens.length);
        for (uint i = 0; i < tokens.length; i++) {
            embeddings[i] = embeddingMatrix[uint64(toInt(tokens[i]))];
        }

        return Tensor2DMethods.from(embeddings);
    }

    function forward(Tensors.TensorData[] calldata input) external view returns (Tensors.TensorData memory) {
        if (input[0].dim.length == 0) {
            revert IncorrectTensorType();
        }        
        if (input[0].dim.length == 1) {
            Float32x32[] memory x = abi.decode(input[0].data, (Float32x32[]));
            Float32x32[][] memory embeddingMatrix = w.mat;
            Tensors.Tensor2D memory zt = getEmbedding(embeddingMatrix, x);
            return Tensors.TensorData(abi.encode(zt.mat), zt.getDim());
        }
        revert TensorTypeNotSupported(); 
    }

    function appendWeights(Float32x32[] calldata x, uint idx) external returns (uint, bool) {
        uint ptrLayer = _ptrLayer;
        uint ptr = _ptr;
        if (ptrLayer == 0) {
            uint m = w.m;
            uint cnt = w.n * w.m;
            while (idx < x.length && ptr < cnt) {
                w.mat[ptr / m].push(x[idx]);
                ptr++; idx++;
            }
            if (ptr == cnt) { ++ptrLayer; ptr = 0; }
        }
        _ptrLayer = ptrLayer;
        _ptr = ptr;
        return (idx, ptrLayer == 1);
    }

    function getWeight(uint i, uint j) external view returns (Float32x32) {
        return w.mat[i][j];
    }
}
