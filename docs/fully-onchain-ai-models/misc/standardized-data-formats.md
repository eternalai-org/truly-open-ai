# Standardized data formats

To ensure seamless communication and accurate processing of user input with AI models deployed on the blockchain, a clear set of instructions is essential. This is where the metadata JSON file plays a crucial role. This file acts as a translator, bridging the gap between your on-chain AI model and the interface client, guiding the client on how to prepare input data and interpret model outputs.

## Why is Metadata Important in a Blockchain Context?

* **On-Chain Efficiency:** Smart contracts running on the blockchain have resource constraints. By specifying the required input format and transformations in the metadata, the interface client can preprocess the data before sending it to the smart contract, optimizing gas usage and minimizing costs.
* **Standardized Interface:** The metadata establishes a consistent structure for interacting with different AI models on the blockchain, simplifying the development of interface clients and promoting interoperability.
* **Data Integrity:** By defining the expected input format and transformations, the metadata helps ensure that only valid data is sent to the smart contract, safeguarding the integrity of the model's predictions.

## Structure of the Metadata JSON File

The metadata file adheres to a structured JSON format with the following key sections:

```json
{
  "input": {
    "type": "...",
    "transforms": [...]
  },
  "output": {
    "type": "...",
    "parameters": {...}
  }
}
```

Let's delve into each section:

### **Input Section**

* **type:** Specifies the type of input data your blockchain-based AI model expects. Common types include:
  * `image`
  * `text (coming soon)`
  * `audio (coming soon)`
* **transforms:** An array detailing the transformations that the interface client must apply to the input data before sending it to the smart contract. Each transformation is an object with the following properties:
  * `type`: The type of transformation (e.g., "resize," "normalize," "grayscale").
  * `parameters`: An object containing the parameters specific to the transformation.

### **Output Section**

* **type:** Specifies the type of output your AI model produces. Common types include:
  * `classification`
  * `...`
* **parameters:** An object containing additional information relevant to the output format. For example:
  * For `classification` tasks, you might include an array of `labels` corresponding to the possible classes.

### Example Metadata File

The following is a sample metadata JSON file for an image classification model that runs on the blockchain:

```json
{
    "input": {
      "type": "image",
      "transforms":[{
        "type": "grayscale",
        "parameters": {
        "enabled": true
        }},
        {
          "type": "resize",
          "parameters": {
            "width": 28,
            "height": 28
          }
        },
        {
          "type": "normalize",
          "parameters": {
            "method": "minmax",
            "range": {
              "min": 0,
              "max": 1
            }
          }
        }
      ]
    },
    "output": {
      "type": "classification",
      "parameters": {
        "labels": [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
      }
    }
  }
```

### Explaination

* **Grayscale Transformation**

```json
{
  "type": "grayscale",
  "parameters": {
    "enabled": true
  }
}
```

This transformation converts the input image to grayscale if enabled is set to true.

* **Resize Transformation:**

```json
{
  "type": "resize",
  "parameters": {
    "width": 28,
    "height": 28
  }
}
```

This transformation resizes the input image to a width and height of 28 pixels.

* **Normalize Transformation:**

```json
{
  "type": "normalize",
  "parameters": {
    "method": "minmax",
    "range": {
      "min": 0,
      "max": 1
    }
  }
}
```

This transformation normalizes the pixel values of the image using the min-max method, scaling the values to the range between 0 and 1.

