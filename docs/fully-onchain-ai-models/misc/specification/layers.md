# Layers

## The base layer

This is the class from which all layers inherit.

```solidity
interface ILayer {
    error TensorTypeNotSupported();
    error IncorrectTensorType();
    error IncorrectTensorDim();

    function appendWeights(Float32x32[] calldata weights, uint idx) external returns (uint, bool);
    function getParamsCount() external view returns (uint);
    function predict(Tensors.TensorData[] calldata input) external returns (Tensors.TensorData memory);    
}
```

## Layer activations

:white\_check\_mark: relu\
:white\_check\_mark: sigmoid\
:white\_check\_mark: softmax\
:white\_check\_mark: leakyrelu\
:white\_check\_mark: tanh\
:white\_check\_mark: linear

_Coming soon:_

* softplus
* softsign

## Core layers

:white\_check\_mark: Input layer\
:white\_check\_mark: Dense layer\
:white\_check\_mark: Embedding layer

_Coming soon:_

* _Masking layer_
* _Lambda layer_
* _Identity layer_

## Convolutional Layers

:white\_check\_mark: Conv2D layer

:white\_check\_mark: _Coming soon:_

* _Conv3D layer_

## Pooling layers

:white\_check\_mark: MaxPooling2D layer\
:white\_check\_mark: AveragePooling2D layer

## Recurrent layers

:white\_check\_mark: SimpleRNN layer\
:white\_check\_mark: LSTM layer

_Coming soon:_

* _GRU layer_
* _Bidirectional layer_

## Normalization layers

_Coming soon:_

* _BatchNormalization layer_
* _LayerNormalization layer_
* _UnitNormalization layer_
* _GroupNormalization layer_

## Reshaping layers

:white\_check\_mark: Flatten layer\
:white\_check\_mark: Rescale layer

_Coming soon:_

* _ZeroPadding2D layer_
* _Reshape layer_
* _Permute layer_

## Merging layers

:white\_check\_mark: Add layer\
:white\_check\_mark: Subtract layer\
:white\_check\_mark: Multiply layer

_Coming soon:_

* _Concatenate layer_
* _Average layer_
* _Maximum layer_
* _Minimum layer_
* _Dot layer_

## Activation Layers

:white\_check\_mark: ReLU layer\
:white\_check\_mark: Sigmoid layer

_Coming soon:_

* _Softmax layer_
* _LeakyReLU layer_
* _PReLU layer_
* _ELU layer_

## Special Purpose Layers

:white\_check\_mark: OnesLike layer\
:white\_check\_mark: ZerosLike layer

