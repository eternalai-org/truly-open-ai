# Transforming an AI Model into an Eternal

This page explains the transformation process required to convert an AI model into an Eternal AI model suitable for deployment on the EternalAI blockchain. The process involves two primary phases: the export phase and the deployment phase.

## Phase 1: Export

The export phase involves capturing the AI model’s architecture and weights and preparing them for deployment. This phase ensures that the model’s structure and data are accurately represented and stored in a format suitable for blockchain deployment.

### Steps in the Export Phase

#### 1. Model Graph Export:

* Convert the model’s architecture to a JSON-compatible format.
* Handle different Keras versions (2.x and 3.x) to export inbound nodes correctly.
* For Functional models, extract layer configurations and inbound nodes.
* For Sequential models, process each layer to capture its configuration and connections.

#### 2. Weights Export:

* Extract and flatten the model’s weights into a list format.
* Iterate through each layer to convert weights to a single-dimensional list.

#### 3. Hashing:

* Generate a hash for both the model graph and the weights using SHA-256.
* Concatenate these hashes to create a combined hash, ensuring data integrity.&#x20;

#### 4. Export Completion:

* Combine the model graph, weights, and the combined hash into a single dictionary.
* This exported model data can then be used for the deployment phase.

### Example Code

The export functionality is encapsulated in the ModelExporter class, which includes methods for exporting the model graph, weights, and generating the combined hash.

```python
class ModelExporter:
    def _export_model_graph(self, model):
        # Convert the model to JSON format and process layers
        pass

    def _export_weights(self, model):
        # Extract and flatten model weights
        pass

    def _export_tf_model(self, model):
        # Export model graph and weights, generate hashes
        pass

    def export_model(self, model):
        # Export the complete model data
        pass
```

## Phase 2: Deployment

The deployment phase involves deploying the exported model data onto the EternalAI blockchain. This phase ensures that each layer and weight of the model is correctly encoded, deployed as smart contracts, and linked together.

### Steps in the Deployment Phase

#### 1. Initialization:

* Establish a connection to the blockchain using a node endpoint.
* Set up the private key for transaction signing.
* Initialize cache for storing deployment data.

#### 2. Deploy Contract from Artifact:

* Deploy the main model contract using a predefined artifact (ABI and bytecode).
* Wait for the transaction receipt and log the deployed contract address.

#### 3. Extract Model Configuration:

* Process the model layers to extract configuration data.
* Determine the type of each layer and encode its configuration data.
* Calculate the total number of weights required for the model.&#x20;

#### 4. Deploy Individual Layers:

* Deploy each layer of the model as a separate smart contract.
* Retrieve and log the contract address for each deployed layer.
* Store deployed layer information in the cache.

#### 5. Construct the Model:

* Link the deployed layer contracts to construct the model.
* Verify the integrity of the model using the generated hash.
* Update the cache to mark the model as constructed.

#### 6. Upload Model Weights:

* Split the weights into chunks and upload them to the blockchain in multiple transactions.
* Monitor the transaction status to ensure successful uploads.

#### 7. Finalization:

* Clear the cache data related to the deployed model.
* Log the successful deployment of the model.
* Return the contract address of the deployed model.

### Why Upload Model Weights by Chunk?

Uploading model weights by chunk is necessary because blockchain transactions have gas limits, which restrict the amount of data that can be included in a single transaction. By splitting the weights into smaller chunks, we can ensure that each transaction remains within the gas limit, thus avoiding transaction failures and ensuring that the entire set of weights is successfully uploaded to the blockchain.

### Example Code

The deployment functionality is encapsulated in the ModelDeployer class, which includes methods for deploying contracts, extracting configurations, and uploading weights.

```python
class ModelDeployer:
    def deploy_from_artifact(self):
        # Deploy the main contract from artifact
        pass

    def get_model_config(self, layers):
        # Extract and encode layer configurations
        pass

    def uploadModelWeights(self, model, weights, start_idx=0):
        # Upload model weights to the blockchain
        pass

    def deploy_layer(self, layer_data):
        # Deploy individual layer contracts
        pass

    def deploy_model(self, model_data):
        # Deploy the complete model
        pass
```

## Conclusion

The transformation process to convert an AI model into an Eternal AI model involves two critical phases: export and deployment. The export phase captures the model’s architecture and weights, ensuring they are accurately represented and hashed for integrity. The deployment phase ensures that the model is correctly deployed onto the EternalAI blockchain, with each layer and weight encoded and linked together. This comprehensive process ensures the integrity, accuracy, and successful deployment of AI models on the blockchain.
