# Train an AI model in Keras

## Download a pre-trained AI model

We're making it easy for you. If you have never trained an AI model before, you can download a simple pre-trained AI model [here](https://github.com/eternalai-org/eternal-model-zoo).

Once you're familiar with the Eternal AI Toolkit, you can play around with the more sophisticated models from popular AI community sites like Kaggle and Hugging Face. There are thousands of AI models published by AI enthusiasts and researchers around the world.

{% embed url="https://www.kaggle.com/models" %}

{% embed url="https://huggingface.co/models?library=keras&sort=trending" %}

## Train your own AI model in Keras

With the Eternal AI Python SDK, there is no need to learn new tools or languages like Solidity. You can just train your AI models in Keras as usual. Then, transform them into Eternals with just a few lines of Python code.

Here is a simple example of training an AI model to recognize handwritten digits.

```python
import eai
from keras import layers, models, datasets

# Load MNIST data
(train_images, train_labels), _ = datasets.mnist.load_data()
train_images = train_images[:1000].reshape((1000, 28, 28, 1)) / 255

# Define a simple Keras model
model = models.Sequential([
    layers.Conv2D(8, (3, 3), activation="relu", input_shape=(28, 28, 1)),
    layers.MaxPooling2D((2, 2)),
    layers.Flatten(),
    layers.Dense(10, activation="softmax"),
])
model.compile(optimizer='adam', loss='sparse_categorical_crossentropy')

# Train briefly
model.fit(train_images, train_labels[:1000], epochs=1)
```

&#x20;For more code examples, check out the Keras docs.

{% embed url="https://keras.io/examples/" fullWidth="false" %}
