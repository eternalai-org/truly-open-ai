# Transform the Keras model to Eternal

Without learning a new language like Solidity, you can easily transform your existing AI models into Eternals, which are autonomous, unstoppable, and cryptographically-secure AI agents running on the blockchain. Under the hood, the Eternal AI toolkit takes your AI model structures and weights, transforms them into smart contracts, and deploys them on the blockchain.

{% hint style="info" %}
The current model size limit is 1 MB. We're working on increasing the limit to accommodate more models. Our goal is to reach 100 MB in Q3-2024 and 1 GB in Q4-2024.
{% endhint %}

## CLI quickstart

Now that you have the Eternal AI CLI installed and your wallet funded, you can transform an existing AI model into an Eternal.

To deploy an Eternal with an existing AI model, run:

```bash
eai eternal transform --format FORMAT --file FILE --name ETERNAL_NAME
```

Or you can deploy from a URL:

```bash
eai eternal transform --format FORMAT --url URL --name ETERNAL_NAME
```

{% hint style="info" %}
Only `Keras` file format is supported currently. We're adding support for more model file formats, such as `PyTorch` and `TensorFlow`.
{% endhint %}

See the pending transaction with:

```bash
eai wallet transactions
```

Once the Eternal has been deployed, the Eternal ID should be printed when you run:

```bash
eai eternal list
```

## Python quickstart

If you’ve built AI models in Keras, you can easily transform them into Eternals. Eternal AI offers a powerful yet simple Python library to transform your Keras models into Eternals with just a few lines of code. No need to learn new tools or languages like Solidity.

```python
import eai

# Transform the Keras model to an Eternal
eternal = eai.transform(model, model_name="my-first-eternal")
```

After your Eternal is deployed, anyone can interact with it.

```python
# try a single predict call
sample = x_test[0]
output = eai.predict([sample])
print(output)
```

## No-code quickstart

Use Eternal AI without writing code. With just a few drags and drops, you can transform your AI model into an Eternal.

### Upload

Upload your AI model file at the [Eternal AI Creators ](https://eternalai.org/model/create)page.

<figure><img src="../../.gitbook/assets/image (19).png" alt=""><figcaption></figcaption></figure>

### Describe your Eternal

Next, enter your Eternal information and get ready for the transformation.&#x20;

<figure><img src="../../.gitbook/assets/image (20).png" alt=""><figcaption></figcaption></figure>
