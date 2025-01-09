---
description: Demonstration of how to interactively chat with your dagent.
---

# Decentralized ChatGPT

### 1. Install Miniconda

Download and install Miniconda by following the instructions here: [Miniconda Installation.](https://docs.anaconda.com/miniconda/install/#quick-command-line-install)

### 2. Clone the repository

```
git clone https://github.com/eternalai-org/eternal-ai.git
cd eternal-ai/examples/decentralized-chatgpt
```

### 3. Activate Conda Environment

Create and activate a Conda environment for Python 3.10.0:

```
conda create -n dagents_chat python=3.10.0
conda activate dagents_chat
```

### 4. Install Dependencies

Install the required Python packages using `pip`.&#x20;

```
pip install -r requirements.txt
```

&#x20;Compile and build `eternal_dagents` package

```
pip install --force-reinstall .
```

### 5. Create a `.env` File

In the root of your repository, create a `.env` file to store your environment variables.

```
cp .env.example .env
```

### 6. Obtain Inference API Key:

Go to [the page](https://eternalai.org/api) to retrieve your inference API key.

### 7. Configure the `.env` File:

Open your `.env` file and insert values of environment variables as below:

```
ETERNALAI_URL=https://api.eternalai.org
ETERNALAI_API_KEY=your_inference_api_key
ETERNAL_CHAIN_ID=8453
ETERNAL_MODEL_NAME=PrimeIntellect/INTELLECT-1-Instruct
```

### 8. (Optional) Edit dAgent Personality:

You can customize your dagent's personality by editing the `config/eternal.json` file according to your preferences. [Learn more](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/open-source/adjust-your-agent-personality).

### 9. Run the Simple Chat:

Execute the script to start the application:

```
python3 toolkit/chat-lite.py 
```

&#x20;Sample output:

```
Let's chat with your eternal! (Ctrl-C to break)
> You: tell me about bitcoin
> Eternal: Bitcoin is a decentralized digital currency that operates on a peer-to-peer network without the need for a central authority. It was created in 2009 by an anonymous person or group of people known as Satoshi Nakamoto. Bitcoin is often referred to as a cryptocurrency because it uses cryptography to secure transactions and control the creation of new units.

Bitcoin is unique in that it is the first and most well-known cryptocurrency. It is also the largest and most widely used cryptocurrency by market capitalization. Bitcoin is traded on various cryptocurrency exchanges around the world and can be used to purchase goods and services from merchants who accept it as payment.

One of the key features of Bitcoin is that it is decentralized, meaning that it is not controlled by any central authority or government. This makes it a potentially revolutionary form of currency that could be used to create a more equitable and transparent financial system.

Bitcoin also has the potential to be used as a store of value, similar to gold or other precious metals. This is because Bitcoin has a limited supply, with only 21 million coins ever to be created. This scarcity makes Bitcoin a potentially valuable asset.

However, there are also some challenges associated with Bitcoin. One of the biggest challenges is that it is still a relatively new and untested technology. There are also concerns about the security and stability of the Bitcoin network, as well as the potential for fraud and other issues.

Despite these challenges, Bitcoin has the potential to be a major player in the future of finance. It is already being used by many businesses and individuals around the world, and its popularity is only likely to grow.
------------------------------------------------------------------------------------------------------------------------
Infer-tx: 0x1ebb16eac3a919670a661333073f6ec7b8d6859fc19bb2e12e36d56ec54eb570
Propose-tx: 0x31ce8b84e2588c380ad930e44153b87f77e722ce14b08d3cd71a5d23ece07af5
```

As shown in the sample output above, since we were chatting with a decentralized agent, anyone can access the on-chain data of the conversation.

* Prompt request tx: [https://basescan.org/tx/0x1ebb16eac3a919670a661333073f6ec7b8d6859fc19bb2e12e36d56ec54eb570](https://basescan.org/tx/0x1ebb16eac3a919670a661333073f6ec7b8d6859fc19bb2e12e36d56ec54eb570)
* Response tx: [https://basescan.org/tx/0x31ce8b84e2588c380ad930e44153b87f77e722ce14b08d3cd71a5d23ece07af5](https://basescan.org/tx/0x31ce8b84e2588c380ad930e44153b87f77e722ce14b08d3cd71a5d23ece07af5)
