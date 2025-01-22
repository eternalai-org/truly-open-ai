from enum import Enum
from typing import List

LayerType = Enum('LayerType', [
    'InputLayer',
    'Dense',
    'Flatten',
    'Rescaling',
    'MaxPooling2D',
    'AveragePooling2D',
    'Conv2D',
    'BatchNormalization',
    'Embedding',
    'SimpleRNN',
    'LSTM',
    'Softmax',
    'Sigmoid',
    'ReLU',
    'Linear',
    'Add',
    'Dropout',
    'GlobalAveragePooling2D',
    'ZeroPadding2D',
    'Concatenate'
], start=0)

InputType = Enum('InputType', [
    'Scalar',
    'Tensor1D',
    'Tensor2D',
    'Tensor3D',
], start=0)

Activation = Enum('Activation', [
    'leakyrelu',
    'linear',
    'relu',
    'sigmoid',
    'tanh',
    'softmax',
], start=0)

Padding = Enum('Padding', [
    'valid',
    'same',
], start=0)

ZeroPaddingFormat = Enum('ZeroPaddingFormat', [
    'channels_first',
    'channels_last',
], start=0)


class LayerData:
    def __init__(
            self,
            layerType: LayerType,
            layerName: str,
            configData: str,
            inputIndices: List[int]):
        self.layerType = layerType
        self.layerName = layerName
        self.configData = configData
        self.inputIndices = inputIndices


class LayerConfig:
    def __init__(
            self,
            layerTypeIndex: int,
            layerAddress: str,
            inputIndices: List[int]):
        self.layerTypeIndex = layerTypeIndex
        self.layerAddress = layerAddress
        self.inputIndices = inputIndices

    def toContractParams(self):
        return (self.layerTypeIndex, self.layerAddress, self.inputIndices)


KERAS_ACTIVATIONS = {
    "relu": "ReLU",
    "sigmoid": "Sigmoid",
    "tanh": "Tanh",
    "softmax": "Softmax",
    "softplus": "Softplus",
    "softsign": "Softsign",
    "elu": "ELU",
    "selu": "SELU",
    "swish": "Swish",
    "gelu": "GELU",
    "exponential": "Exponential",
    "hard_sigmoid": "HardSigmoid",
    "linear": "Linear",
    "leaky_relu": "LeakyReLU",
    "silu": "SiLU",
    "hard_silu": "HardSiLU",
}