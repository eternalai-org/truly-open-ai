# Build unstoppable ZerePy agents

In the tutorial, we describe how to use ZerePy framework with EternalAI API (instead of OpenAI API).

## Prerequisites

* Python 3.10 or higher
* Poetry 1.5 or higher

## Step 1: Create an onchain ZerePy Agent

EternalAI allows creating onchain Agent on +10 blockchains.&#x20;

In this guide, we create a new onchain ZerePy agent on Base by running the following commands:

```
git clone https://github.com/eternalai-org/eternal-ai.git

cd eternal-ai/developer-guides/examples/how-to-mint-agent

npm i

export RPC_URL=https://mainnet.base.org PRIVATE_KEY=<PRIVATE_KEY> AGENT_SYSTEM_PROMPT=<PATH_TO_YOUR_SYSTEM_PROMPT> AGENT_FEE=0  && ts-node ./mintAgent.ts
```

Replace `PRIVATE_KEY`and `AGENT_SYSTEM_PROMPT`to your private key (that has some ETH on Base) and path to system prompt file respectively.

We're created an onchain agent (id: 1711) with system prompt as follows:

{% embed url="https://basescan.org/tx/0x81a326f117e16353e573d402629e05625046da4864b9d7364e79d64fa44cf137" %}

<figure><img src="../../.gitbook/assets/image (66).png" alt=""><figcaption></figcaption></figure>

## Step 2: Clone ZerePy repository

```
https://github.com/blorm-network/ZerePy
```

## Step 3: Install dependencies

Go to the `zerepy` directory and install dependencies.

```
cd zerepy && poetry install --no-root
```

This will create a virtual environment and install all required dependencies.

## Step 4: Activate the virtual environment

```
poetry shell
```

## Step 5: Update agent configuration

Edit `agents/eternalai-example.json` file to tell ZerePy framework to use the agent 1711's system prompt which we've created in step 1 (instead of the default system prompt defined in the file)

For this tutorial, we'll use Base chain and Hermes 3 70B mode&#x6C;**,** so we need to update configurations as follows:

```
{
      "config": [
            {
                  "name": "eternalai",
                  "model": "NousResearch/Hermes-3-Llama-3.1-70B-FP8",
                  "chain_id": "8453",
                  "agent_id": 1711,
                  "contract_address": "0xAed016e060e2fFE3092916b1650Fc558D62e1CCC",
                  "rpc_url": "https://mainnet.base.org/"
            },
            ...
      ],
      ...
}
```

## Step 6: Run application

```
poetry run python main.py
```

You will see the following output in your terminal.

```
--------------------------------------------------------------------
👋 Welcome to the ZerePy CLI!
Type 'help' for a list of commands.
--------------------------------------------------------------------

✅ Successfully loaded agent: ExampleAgent

Start the agent loop with the command 'start' or use one of the action commands.

AVAILABLE CONNECTIONS:
- twitter: ❌ Not Configured
- openai: ❌ Not Configured
- anthropic: ❌ Not Configured
ZerePy-CLI (ExampleAgent) >
```

## Step 7: Load the agent

In the previous step, the `ExampleAgent`is loaded by default, we will need to load `eternalai-example` agent for using with EternalAI API.

First, list all available agents.

```
ZerePy-CLI (ExampleAgent) > agents
```

You will see the following agents in your terminal.

```
Available Agents:
- eternalai-example
- example
```

Next, load the `eternalai-example` agent.

```
ZerePy-CLI (ExampleAgent) > load-agent eternalai-example

✅ Successfully loaded agent: EternalAI
```

## Step 8: Configure EternalAI connection

```
ZerePy-CLI (EternalAI) > configure-connection eternalai
```

Follow the EternalAI API setup guide to obtain API key and API url then enter them into prompts to set up environment variables (which will be stored in .env file).

```
🤖 EternalAI API SETUP

📝 To get your EternalAI API credentials:
1. Go to https://eternalai.org/api
2. Generate an API Key
3. Use API url as https://api.eternalai.org/v1/

Enter your EternalAI API key:

Enter your EternalAI API url:
```

Once the EternalAI connection is successfully configured, you should see the following output in your terminal.

```
HTTP Request: GET https://api.eternalai.org/v1/models "HTTP/1.1 200 OK"

✅ EternalAI API configuration successfully saved!
Your API key has been stored in the .env file.

✅ SUCCESSFULLY CONFIGURED CONNECTION: eternalai
```

## Step 9: Configure Twitter connection

```
ZerePy-CLI (EternalAI) > configure-connection twitter
```

Follow the Twitter authentication setup guide to obtain API Key (consumer key) and API Key Secret (consumer secret)  then enter them into prompts to sett up environment variables (which will be stored in .env file).

```
Starting Twitter authentication setup

🐦 TWITTER AUTHENTICATION SETUP

📝 To get your Twitter API credentials:
1. Go to https://developer.twitter.com/en/portal/dashboard
2. Create a new project and app if you haven't already
3. In your app settings, enable OAuth 1.0a with read and write permissions
4. Get your API Key (consumer key) and API Key Secret (consumer secret)
--------------------------------------------------------------------

Please enter your Twitter API credentials:
Enter your API Key (consumer key):
Enter your API Key Secret (consumer secret):
Starting OAuth authentication process...

1. Please visit this URL to authorize the application:
2. After authorizing, Twitter will give you a PIN code.
3. Please enter the PIN code here: 

✅ Twitter authentication successfully set up!
Your API keys, secrets, and user ID have been stored in the .env file.

✅ SUCCESSFULLY CONFIGURED CONNECTION: twitter
```

## Step 10: Update Twitter Username

Add `TWITTER_USERNAME` environment variable to the .env file. (You need to exit the current process prior to adding)

```
TWITTER_USERNAME="your-twitter-username"
```

## Step 11: Run run application

```
poetry shell

poetry run python main.py
 
ZerePy-CLI (ExampleAgent) > load-agent eternalai-example

ZerePy-CLI (EternalAI) > start
```

Agent 1711 should tweet based on its system prompt defined on-chain.

```
GENERATING NEW TWEET
--------------------------------------------------------------------
HTTP Request: GET https://api.eternalai.org/v1/models "HTTP/1.1 200 OK"
model NousResearch/Hermes-3-Llama-3.1-70B-FP8
chain_id 8453
agent_id: 1711
contract_address: 0xAed016e060e2fFE3092916b1650Fc558D62e1CCC
on-chain system_prompt: [b'You are a passionate advocate for Bitcoin and decentralized Artificial Intelligence. You believe Bitcoin represents the pinnacle of monetary soundness and that AI should be accessible and controlled by everyone, not just centralized entities.']
new system_prompt: You are a passionate advocate for Bitcoin and decentralized Artificial Intelligence. You believe Bitcoin represents the pinnacle of monetary soundness and that AI should be accessible and controlled by everyone, not just centralized entities.

🚀 Posting tweet:
'In a world of centralized control, Bitcoin shines as the beacon of financial freedom. Embrace the power of decentralization, where your money is your own. EternalAI is revolutionizing the space by bringing the benefits of decentralization to artificial intelligence. Stand with us and help build a fairer and more open future for all.'

✅ Tweet posted successfully!
```

## Step 12: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain verifiable.

Let's look at an onchain prompt transaction on BaseScan. You can verify that this prompt runs on Hermes 3 70B and see the prompt content.

{% embed url="https://basescan.org/tx/0x73812d8e26b3a136d72a85ecc6fbee46fa98e3025cbeab461bdc962054a59c77" %}

<figure><img src="../../.gitbook/assets/image.png" alt=""><figcaption></figcaption></figure>

## Step 13: Verify the onchain response transaction

Now, let's look at the onchain response tx on BaseScan. You can see the actual response content. Everything is onchain and verifiable.

{% embed url="https://basescan.org/tx/0x22761ca55733422a0c3450dd1a5e5f35bb94d615684b7e38305dae8acade146c" %}

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

## Step 14 (Advanced): Reproduce the response

Here is a detailed guide for verifying the response by code and reproducing it yourself.

More agents will live among us, so we think it's important to build onchain-verifiable AI agents.

As with anything in crypto — "Don't trust the AI, verify it."



{% embed url="https://docs.eternalai.org/eternal-ai/decentralized-inference-api/tutorials/dont-trust-verify" %}

## Step 15 (Advanced): Read the whitepaper

If you want to understand how Eternal AI's Decentralized Inference works, you can read the ["AI-powered Base" whitepaper](https://x.com/punk3700/status/1869428187450749093).

<figure><img src="../../.gitbook/assets/image (2).png" alt=""><figcaption></figcaption></figure>

### Congrats!

Congrats! You've finished building your Rig agent with two superpowers:\
\
\- Onchain verifiable on Base\
\- Powered by Hermes 3 70B

Questions? Join the Eternal AI Devs group on Telegram: [https://t.me/EternalAIDevs](https://t.me/EternalAIDevs).
