// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

library Layers {
    enum LayerType {
        Input,
        Dense,
        Flatten,
        Rescale,
        MaxPooling2D,
        AveragePooling2D,
        Conv2D,
        BatchNormalization,
        Embedding,
        SimpleRNN,
        LSTM,
        Softmax,
        Sigmoid,
        ReLU,
        Linear,
        Add,
        Dropout,
        GlobalAveragePooling2D,
        ZeroPadding2D,
        Concatenate
    }

    enum InputType {
        Scalar,
        Tensor1D,
        Tensor2D,
        Tensor3D
    }
}
