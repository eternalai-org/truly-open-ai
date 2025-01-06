# Models

A model grouping layers into an object with training/inference features.

There are two ways to instantiate a Model:

## Functional Models

You can check [Keras document](https://keras.io/api/models/model/) to learn more about functional models. The following EternalAI API helps to create an on-chain functional model from a specifc model config.

```solidity
function constructModel(
        LayerConfig[] calldata layersConfig
    ) external onlyOwner
```

After constructing a model, you can upload model weights by chunks, using the following API.

```solidity
function appendWeights(
        Float32x32[] calldata weights
    ) external onlyOwner
```

## Sequential Models

Sequential groups a linear stack of layers into a Model. [Learn more.](https://keras.io/api/models/sequential/)
