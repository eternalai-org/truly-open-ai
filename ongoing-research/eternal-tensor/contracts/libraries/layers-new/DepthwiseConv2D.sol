// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract DepthwiseConv2DLayer is ILayer {

    using TensorMethods for Tensors.Tensor;
    
    Tensors.ActivationFunc public activation;
    uint public filters;
    bool public useBias;
    uint[2] public stride;
    uint public padding;
    Tensors.Tensor w;
    Tensors.Tensor b;

    uint _ptrLayer;
    uint _ptr;
    uint nLayer;
    uint[] remainingWeight;

    constructor(bytes memory config) {
        (uint8 _actv, uint256 _inFilters, uint256[2] memory _size, uint256[2] memory _stride, uint8 _padding, bool _useBias) = abi.decode(
            config,
            (uint8, uint256, uint256[2], uint256[2], uint8, bool)
        );
        activation = Tensors.ActivationFunc(_actv);
        filters = _inFilters;
        stride = _stride;
        padding = _padding;
        useBias = _useBias;

        nLayer = 2;
        TensorMethods.initStorageZerosTensor(w, Tensors.get4DShape(_size[0], _size[1], _inFilters, 1));
        TensorMethods.initStorageZerosTensor(b, Tensors.get1DShape(_inFilters));

        remainingWeight = new uint[](nLayer);
        remainingWeight[0] = Tensors.getWeightCount(w.shapes);
        remainingWeight[1] = useBias ? Tensors.getWeightCount(b.shapes) : 0;
        for(uint i = nLayer-1; i > 0; --i) {
            remainingWeight[i-1] += remainingWeight[i];
        }
    }
    function getWeightCount() external view returns (uint) {
        return remainingWeight[0];
    }

    function getRemainingWeightCount() external view returns (uint) {
        return (_ptrLayer == nLayer) ? 0 : remainingWeight[_ptrLayer] - _ptr;
    }

    function forward(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory) {
        if (input[0].shapes.length == 0) {
            revert IncorrectTensorType();
        }        

        Tensors.Tensor memory yt = input[0].depthwiseConv2D(w,b,stride[0],stride[1],padding);
        Tensors.Tensor memory zt = yt.activation(activation);
        return zt;
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        uint ptrLayer = _ptrLayer;
        uint ptr = _ptr;
        uint idx = 0;
        if (ptrLayer == 0) {
            (ptrLayer, ptr, idx) = TensorMethods.appendStorageTensor(w, x, ptrLayer, ptr, idx);
        }
        if (ptrLayer == 1) {
            if (useBias) {
                (ptrLayer, ptr, idx) = TensorMethods.appendStorageTensor(b, x, ptrLayer, ptr, idx);
            }
            else {
                ++ptrLayer;
            }
        }
        _ptrLayer = ptrLayer;
        _ptr = ptr;
        return (ptrLayer == 2);
    }

    function getFilters() external view returns (uint) {
        return filters;
    }

    function getStride() external view returns (uint[2] memory) {
        return stride;
    }

    function getW() external view returns (uint256[] memory, uint64[] memory) {
        return (w.data, w.shapes);
    }

    function getB() external view returns (uint256[] memory, uint64[] memory) {
        return (b.data, b.shapes);
    }
}
