// SPDX-License-Identifier: UNLICENSED
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

contract LSTMLayer {
    struct LSTMWeights {
        Tensors.Tensor2D kernelF;
        Tensors.Tensor2D kernelI;
        Tensors.Tensor2D kernelO;
        Tensors.Tensor2D kernelC;

        Tensors.Tensor2D recurrentKernelF;
        Tensors.Tensor2D recurrentKernelI;
        Tensors.Tensor2D recurrentKernelO;
        Tensors.Tensor2D recurrentKernelC;

        Tensors.Tensor1D biasF;
        Tensors.Tensor1D biasI;
        Tensors.Tensor1D biasO;
        Tensors.Tensor1D biasC;
    }

    error WeightsAppendingOutOfBound();

    using Tensor1DMethods for Tensors.Tensor1D;
    using Tensor2DMethods for Tensors.Tensor2D;

    LSTMWeights private weights;

    uint256 public inputDim;
    uint256 public outputDim;
    uint256 public matrixPointer;
    uint256 public cellPointer;

    Tensors.ActivationFunc public activationFunc;
    Tensors.ActivationFunc public recurrentActivationFunc;

    constructor(bytes memory _config) {
        (
            uint8 _activationFunc,
            uint8 _recurrentActivationFunc,
            uint256 _units,
            uint256 _inputDim
        ) = abi.decode(
            _config,
            (
                uint8,
                uint8,
                uint256,
                uint256
            )
        );

        activationFunc = Tensors.ActivationFunc(_activationFunc);
        recurrentActivationFunc = Tensors.ActivationFunc(_recurrentActivationFunc);
        inputDim = _inputDim;
        outputDim = _units;

        weights = LSTMWeights(
            Tensor2DMethods.emptyTensor(_inputDim, _units),
            Tensor2DMethods.emptyTensor(_inputDim, _units),
            Tensor2DMethods.emptyTensor(_inputDim, _units),
            Tensor2DMethods.emptyTensor(_inputDim, _units),
            Tensor2DMethods.emptyTensor(_units, _units),
            Tensor2DMethods.emptyTensor(_units, _units),
            Tensor2DMethods.emptyTensor(_units, _units),
            Tensor2DMethods.emptyTensor(_units, _units),
            Tensor1DMethods.emptyTensor(_units),
            Tensor1DMethods.emptyTensor(_units),
            Tensor1DMethods.emptyTensor(_units),
            Tensor1DMethods.emptyTensor(_units)
        );
    }

    function getParamsCount() external view returns (uint256) {
        return weights.kernelF.n * weights.kernelF.m
            + weights.kernelI.n * weights.kernelI.m
            + weights.kernelO.n * weights.kernelO.m
            + weights.kernelC.n * weights.kernelC.m
            + weights.recurrentKernelF.n * weights.recurrentKernelF.m
            + weights.recurrentKernelI.n * weights.recurrentKernelI.m
            + weights.recurrentKernelO.n * weights.recurrentKernelO.m
            + weights.recurrentKernelC.n * weights.recurrentKernelC.m
            + weights.biasF.n
            + weights.biasI.n
            + weights.biasO.n
            + weights.biasC.n;
    }

    function _commonFormula(
        Tensors.Tensor1D memory _input,
        Tensors.Tensor1D memory _hiddenState,
        Tensors.Tensor2D memory _kernel,
        Tensors.Tensor2D memory _recurrentKernel,
        Tensors.Tensor1D memory _bias,
        Tensors.ActivationFunc _activationFunc
    ) private view returns (Tensors.Tensor1D memory) {
        return _input.matMul(_kernel)
            .add(_hiddenState.matMul(_recurrentKernel))
            .add(_bias)
            .activation(_activationFunc);
    }

    function forward(Tensors.TensorData[] calldata _input) external view returns (Tensors.TensorData memory) {
        if (_input[0].dim.length != 2) {
            revert IncorrectTensorType();
        }

        Float32x32[][] memory inputs = abi.decode(_input[0].data, (Float32x32[][]));

        uint256 units = outputDim;
        Tensors.Tensor1D memory hiddenState = Tensor1DMethods.zerosTensor(units);
        Tensors.Tensor1D memory cellState = Tensor1DMethods.zerosTensor(units);

        Tensors.ActivationFunc actFunc = activationFunc;
        Tensors.ActivationFunc rActFunc = recurrentActivationFunc;

        LSTMWeights memory w = weights;

        uint256 n = inputs.length;
        for (uint256 i = 0; i < n; ++i) {
            Tensors.Tensor1D memory input = Tensor1DMethods.from(inputs[i]);
            Tensors.Tensor1D memory forgetGate = _commonFormula(
                input,
                hiddenState,
                w.kernelF,
                w.recurrentKernelF,
                w.biasF,
                rActFunc
            );
            Tensors.Tensor1D memory inputGate = _commonFormula(
                input,
                hiddenState,
                w.kernelI,
                w.recurrentKernelI,
                w.biasI,
                rActFunc
            );
            Tensors.Tensor1D memory outputGate = _commonFormula(
                input,
                hiddenState,
                w.kernelO,
                w.recurrentKernelO,
                w.biasO,
                rActFunc
            );
            Tensors.Tensor1D memory cellActivation = _commonFormula(
                input,
                hiddenState,
                w.kernelC,
                w.recurrentKernelC,
                w.biasC,
                actFunc
            );

            cellState = Tensor1DMethods.add(
                forgetGate.mul(cellState),
                inputGate.mul(cellActivation)
            );

            hiddenState = outputGate.mul(cellState.activation(actFunc));
        }

        return Tensors.TensorData(abi.encode(hiddenState.mat), hiddenState.getDim());
    }

    function appendWeights(Float32x32[] calldata _weights, uint256 _idx) external returns (uint256, bool) {
        uint256 mPtr = matrixPointer;
        uint256 cPtr = cellPointer;
        uint256 l = _weights.length;
        LSTMWeights storage w = weights;
        if (mPtr == 0) {
            uint256 m = w.kernelF.m;
            uint256 size = weights.kernelF.n * weights.kernelF.m;
            Float32x32[][] storage mat = w.kernelF.mat;
            while (_idx < l && cPtr < size) {
                mat[cPtr++/m].push(_weights[_idx++]);
            }
            if (cPtr == size) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 1) {
            uint256 m = w.kernelI.m;
            uint256 size = w.kernelI.n * w.kernelI.m;
            Float32x32[][] storage mat = w.kernelI.mat;
            while (_idx < l && cPtr < size) {
                mat[cPtr++/m].push(_weights[_idx++]);
            }
            if (cPtr == size) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 2) {
            uint256 m = w.kernelO.m;
            uint256 size = w.kernelO.n * w.kernelO.m;
            Float32x32[][] storage mat = w.kernelO.mat;
            while (_idx < l && cPtr < size) {
                mat[cPtr++/m].push(_weights[_idx++]);
            }
            if (cPtr == size) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 3) {
            uint256 m = w.kernelC.m;
            uint256 size = w.kernelC.n * w.kernelC.m;
            Float32x32[][] storage mat = w.kernelC.mat;
            while (_idx < l && cPtr < size) {
                mat[cPtr++/m].push(_weights[_idx++]);
            }
            if (cPtr == size) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 4) {
            uint256 m = w.recurrentKernelF.m;
            uint256 size = w.recurrentKernelF.n * w.recurrentKernelF.m;
            Float32x32[][] storage mat = w.recurrentKernelF.mat;
            while (_idx < l && cPtr < size) {
                mat[cPtr++/m].push(_weights[_idx++]);
            }
            if (cPtr == size) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 5) {
            uint256 m = w.recurrentKernelI.m;
            uint256 size = w.recurrentKernelI.n * w.recurrentKernelI.m;
            Float32x32[][] storage mat = w.recurrentKernelI.mat;
            while (_idx < l && cPtr < size) {
                mat[cPtr++/m].push(_weights[_idx++]);
            }
            if (cPtr == size) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 6) {
            uint256 m = w.recurrentKernelO.m;
            uint256 size = w.recurrentKernelO.n * w.recurrentKernelO.m;
            Float32x32[][] storage mat = w.recurrentKernelO.mat;
            while (_idx < l && cPtr < size) {
                mat[cPtr++/m].push(_weights[_idx++]);
            }
            if (cPtr == size) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 7) {
            uint256 m = w.recurrentKernelC.m;
            uint256 size = w.recurrentKernelC.n * w.recurrentKernelC.m;
            Float32x32[][] storage mat = w.recurrentKernelC.mat;
            while (_idx < l && cPtr < size) {
                mat[cPtr++/m].push(_weights[_idx++]);
            }
            if (cPtr == size) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 8) {
            uint256 n = w.biasF.n;
            Float32x32[] storage mat = w.biasF.mat;
            while (_idx < l && cPtr < n) {
                mat.push(_weights[_idx++]);
                cPtr++;
            }
            if (cPtr == n) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 9) {
            uint256 n = w.biasI.n;
            Float32x32[] storage mat = w.biasI.mat;
            while (_idx < l && cPtr < n) {
                mat.push(_weights[_idx++]);
                cPtr++;
            }
            if (cPtr == n) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 10) {
            uint256 n = w.biasO.n;
            Float32x32[] storage mat = w.biasO.mat;
            while (_idx < l && cPtr < n) {
                mat.push(_weights[_idx++]);
                cPtr++;
            }
            if (cPtr == n) {
                ++mPtr;
                cPtr = 0;
            }
        }
        if (mPtr == 11) {
            uint256 n = w.biasC.n;
            Float32x32[] storage mat = w.biasC.mat;
            while (_idx < l && cPtr < n) {
                mat.push(_weights[_idx++]);
                cPtr++;
            }
            if (cPtr == n) {
                ++mPtr;
                cPtr = 0;
            }
        }

        if (_idx < l) {
            revert WeightAppendingOutOfBound();
        }

        matrixPointer = mPtr;
        cellPointer = cPtr;

        return (_idx, mPtr == 12);
    }
}
