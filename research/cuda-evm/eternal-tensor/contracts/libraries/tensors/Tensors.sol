// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

import { Float32x32, fromInt, toInt } from "./../Float32x32/Lib32x32.sol";
import "../libCudaTensor.sol";
import "../libCudaParams.sol";

error InvalidActivationFunction();
error InvalidPaddingType();
error InvalidMatrixDimensions();
error NotAScalar();

library Tensors {
    Float32x32 public constant ONE = Float32x32.wrap(1 << 32);
    Float32x32 public constant TWO = Float32x32.wrap(2 << 32);
    Float32x32 public constant TEN = Float32x32.wrap(10 << 32);

    /*
        * @dev Tensor struct
        * @param data: compressed flattened tensor data
        * @param shapes: shape of the tensor in uncompressed form, mul(shapes) / 4 is the size of data    
    */
    struct Tensor {
        uint256[] data;
        uint64[] shapes;
    }

    // @deprecated
    struct TensorData {
        bytes data;
        uint[] dim;
    }

    // @deprecated
    struct Tensor1D {
        Float32x32[] mat;
        uint n;
    }

    // @deprecated
    struct Tensor2D {
        Float32x32[][] mat;
        uint n;
        uint m;
    }

    // @deprecated
    struct Tensor3D {
        Float32x32[][][] mat;
        uint n;
        uint m;
        uint p;
    }

    // @deprecated
    struct Tensor4D {
        Float32x32[][][][] mat;
        uint n;
        uint m;
        uint p;
        uint q;
    }

    enum ActivationFunc {
        LeakyReLU,
        Linear,
        ReLU,
        Sigmoid,
        Tanh,
        Softmax
    }

    enum PaddingType {
        Valid,
        Same
    }

    enum ZeroPaddingFormat {
        ChannelsFirst,
        ChannelsLast
    }

    function __linear(Float32x32 x) internal pure returns (Float32x32) {
        return x;
    }

    function __relu(Float32x32 x) internal pure returns (Float32x32) {
        return Float32x32.unwrap(x) > 0 ? x : Float32x32.wrap(0);
    }

    function __leaky_relu(Float32x32 x) internal pure returns (Float32x32) {
        return Float32x32.unwrap(x) > 0 ? x : x.mul(TWO).div(TEN);
    }

    function __sigmoid(Float32x32 x) internal pure returns (Float32x32) {
        // Exp will fail if input is greater than 21
        if (Float32x32.unwrap(x) < -21 << 32) {
            x = fromInt(-21);
        }
        return ONE.div(ONE + (-x).exp());
    }

    function __tanh(Float32x32 x) internal pure returns (Float32x32) {
        // Exp will fail if input is greater than 21
        if (Float32x32.unwrap(x) > 10 << 32) {
            x = fromInt(10);
        }
        Float32x32 t = TWO.mul(x).exp();
        return (t - ONE).div(t + ONE);
    }

    function __add(Float32x32 a, Float32x32 b) internal pure returns (Float32x32) {
        return a + b;
    }

    function __mul(Float32x32 a, Float32x32 b) internal pure returns (Float32x32) {
        return a.mul(b);
    }

    function __exp(Float32x32 a) internal pure returns (Float32x32) {
        return a.exp();
    }

    function __subMax0(uint a, uint b) public pure returns (uint) {
        return a > b ? a - b : 0;
    }

    function max(Float32x32 a, Float32x32 b) internal pure returns (Float32x32) {
        return Float32x32.unwrap(a) > Float32x32.unwrap(b) ? a : b;
    }

    function getConvSize(
        uint[2] memory dim,
        uint[2] memory size,
        uint[2] memory stride,
        PaddingType padding
    ) internal pure returns (
        uint[2] memory out,
        uint[2] memory pad
    ) {
        for(uint i = 0; i < 2; ++i) {
            if (padding == PaddingType.Same) {
                out[i] = (dim[i] + stride[i] - 1) / stride[i];
                uint total_pad = (dim[i] % stride[i] == 0) ? __subMax0(size[i], stride[i]) : __subMax0(size[i], dim[i] % stride[i]);
                pad[i] = total_pad / 2;
            } else if (padding == PaddingType.Valid) {
                // TODO: What if dim[i] < size[i]
                out[i] = (dim[i] - size[i]) / stride[i] + 1;
            } else {
                revert InvalidPaddingType();
            }
        }
    }
    
    // Assuming x is untouched
    function get1DShape(uint n) internal pure returns (uint64[] memory) {
        uint64[] memory shape = new uint64[](1);
        shape[0] = uint64(n);
        return shape;
    }

    // Assuming x is untouched
    function get2DShape(uint n, uint m) internal pure returns (uint64[] memory) {
        uint64[] memory shape = new uint64[](2);
        shape[0] = uint64(n);
        shape[1] = uint64(m);
        return shape;
    }

    // Assuming x is untouched
    function get3DShape(uint n, uint m, uint p) internal pure returns (uint64[] memory) {
        uint64[] memory shape = new uint64[](3);
        shape[0] = uint64(n);
        shape[1] = uint64(m);
        shape[2] = uint64(p);
        return shape;
    }

    // Assuming x is untouched
    function get4DShape(uint n, uint m, uint p, uint q) internal pure returns (uint64[] memory) {
        uint64[] memory shape = new uint64[](4);
        shape[0] = uint64(n);
        shape[1] = uint64(m);
        shape[2] = uint64(p);
        shape[3] = uint64(q);
        return shape;
    }

    function getElementCount(uint64[] memory shape) internal pure returns (uint) {
        uint len = 1;
        for(uint i = 0; i < shape.length; ++i) {
            len *= shape[i];
        }
        return len;
    }

    function getWeightCount(uint64[] memory shape) internal pure returns (uint) {
        return (getElementCount(shape) + 3) / 4;
    }

    function min(uint a, uint b) internal pure returns (uint) {
        return a < b ? a : b;
    }
}
