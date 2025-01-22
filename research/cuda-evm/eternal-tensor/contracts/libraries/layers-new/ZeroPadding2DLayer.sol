// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "./ILayer.sol";

contract ZeroPadding2DLayer is ILayer {
    using TensorMethods for Tensors.Tensor;

    uint[4] public padding;
    Tensors.ZeroPaddingFormat public data_format;

    constructor(bytes memory config) {
        (uint[4] memory _padding, uint8 _data_format) = abi.decode(
            config,
            (uint[4], uint8)
        );
        padding = _padding;
        data_format = Tensors.ZeroPaddingFormat(_data_format);
    }

    function getWeightCount() external view returns (uint) {
        return 0;
    }

    function getRemainingWeightCount() external view returns (uint) {
        return 0;
    }

    uint internal constant MASK = (1 << 64) - 1;
    function getValue(Tensors.Tensor calldata x, uint idx) internal pure returns (uint) {
        uint bit = (3 - idx % 4) << 6;
        return (x.data[idx/4] >> bit) & MASK;
    }

    function forward(Tensors.Tensor[] calldata input) external view returns (Tensors.Tensor memory) {
        if (input[0].shapes.length != 3) {
            revert IncorrectTensorType();
        }
        uint t = padding[0];
        uint b = padding[1];
        uint l = padding[2];
        uint r = padding[3];

        uint n = input[0].shapes[0];
        uint m = input[0].shapes[1];
        uint p = input[0].shapes[2];

        Tensors.Tensor memory ol;
        if (data_format == Tensors.ZeroPaddingFormat.ChannelsLast) {
            ol.shapes = Tensors.get3DShape(n+t+b, m+l+r, p);
            uint len = Tensors.getWeightCount(ol.shapes);
            ol.data = new uint[](len);

            for (uint i=0; i < n; i++) {
                for (uint j=0; j < m; j++) {
                    for (uint k = 0; k < p; k++) {
                        uint idx = i * m * p + j * p + k;
                        uint padded_idx = (i + t) * (m + l + r) * p + (j + l) * p + k;
                        uint bit = (3 - padded_idx % 4) << 6;
                        ol.data[padded_idx/4] |= getValue(input[0], idx) << bit;
                    }
                }
            }
        } else if (data_format == Tensors.ZeroPaddingFormat.ChannelsFirst) {
            ol.shapes = Tensors.get3DShape(n, m+t+b, p+l+r);
            uint len = Tensors.getWeightCount(ol.shapes);
            ol.data = new uint[](len);

            for (uint j=0; j < m; j++) {
                for (uint k=0; k < p; k++) {
                    for (uint i = 0; i < n; i++) {
                        uint idx = i * m * p + j * p + k;
                        uint padded_idx = i * (m + t + b) * (p + l + r) + (j + t) * (p + l + r) + (k + l);
                        uint bit = (3 - padded_idx % 4) << 6;
                        ol.data[padded_idx/4] |= getValue(input[0], idx) << bit;
                    }
                }
            }    
        }
        return ol;
    }

    function appendWeights(uint256[] calldata x) external returns (bool) {
        return true;
    }
}
