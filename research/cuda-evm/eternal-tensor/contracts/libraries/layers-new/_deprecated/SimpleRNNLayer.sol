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

contract SimpleRNNLayer {
    using Tensor1DMethods for Tensors.Tensor1D;
    using Tensor2DMethods for Tensors.Tensor2D;
    using Tensor3DMethods for Tensors.Tensor3D;
    using Tensor4DMethods for Tensors.Tensor4D;
    using Tensor1DCuda for Tensors.Tensor1D;
    using Tensor2DCuda for Tensors.Tensor2D;

    Tensors.ActivationFunc public activation;
    uint public inputDim;
    uint public outputDim;
    Tensors.Tensor2D public wx;
    Tensors.Tensor2D public wh;
    Tensors.Tensor1D public b;
    uint _ptrLayer;
    uint _ptr;

    constructor(bytes memory config) {
        (uint8 _actv, uint256 _units, uint256 _inputDim) = abi.decode(
            config,
            (uint8, uint256, uint256)
        );
        activation = Tensors.ActivationFunc(_actv);
        inputDim = _inputDim;
        outputDim = _units;        
        wx = Tensor2DMethods.emptyTensor(_inputDim, _units);
        wh = Tensor2DMethods.emptyTensor(_units, _units);
        b = Tensor1DMethods.emptyTensor(_units);
    }

    function getParamsCount() external view returns (uint) {
        return wx.n * wx.m + wh.n * wh.m + b.n;
    }

    function forward(Tensors.TensorData[] calldata input) external view returns (Tensors.TensorData memory) {
        if (input[0].dim.length == 0) {
            revert IncorrectTensorType();
        }        
        if (input[0].dim.length == 2) {
            Float32x32[][] memory x = abi.decode(input[0].data, (Float32x32[][]));
            Tensors.Tensor2D memory xt = Tensor2DMethods.from(x);
            Tensors.Tensor1D memory ht = Tensor1DMethods.zerosTensor(outputDim);

            for (uint256 t=0; t < xt.n; t++){
                Tensors.Tensor1D memory xt_t = Tensor1DMethods.from(xt.mat[t]);
                Tensors.Tensor1D memory yx_t = xt_t.matMul_CUDA(wx);
                Tensors.Tensor1D memory yh_t = ht.matMul_CUDA(wh);
                Tensors.Tensor1D memory yt = (yx_t.add_CUDA(yh_t)).add_CUDA(b);
                ht = yt.activation_CUDA(activation);
            }
            
            return Tensors.TensorData(abi.encode(ht.mat), ht.getDim());
        }
        revert TensorTypeNotSupported();
    }

    function appendWeights(Float32x32[] calldata x, uint idx) external returns (uint, bool) {
        uint ptrLayer = _ptrLayer;
        uint ptr = _ptr;
        if (ptrLayer == 0) {
            uint m = wx.m;
            uint cnt = wx.n * wx.m;
            while (idx < x.length && ptr < cnt) {
                wx.mat[ptr / m].push(x[idx]);
                ptr++; idx++;
            }
            if (ptr == cnt) { ++ptrLayer; ptr = 0; }
        }
        if (ptrLayer == 1) {
            uint m = wh.m;
            uint cnt = wh.n * wh.m;
            while (idx < x.length && ptr < cnt) {
                wh.mat[ptr / m].push(x[idx]);
                ptr++; idx++;
            }
            if (ptr == cnt) { ++ptrLayer; ptr = 0; }
        }
        if (ptrLayer == 2) {
            uint n = b.n; 
            while (idx < x.length && ptr < n) {
                b.mat.push(x[idx]);
                ptr++; idx++;
            }
            if (ptr == n) { ++ptrLayer; ptr = 0; }
        }        
        _ptrLayer = ptrLayer;
        _ptr = ptr;
        return (idx, ptrLayer == 3);
    }

    function getWeight(uint i, uint j) external view returns (Float32x32) {
        return wx.mat[i][j];
    }

    function getRecurrentWeight(uint i, uint j) external view returns (Float32x32) {
        return wh.mat[i][j];
    }

    function getBias(uint i) external view returns (Float32x32) {
        return b.mat[i];
    }
}
