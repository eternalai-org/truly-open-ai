---
layout:
  title:
    visible: true
  description:
    visible: true
  tableOfContents:
    visible: true
  outline:
    visible: true
  pagination:
    visible: true
---

# Build unstoppable Rig agents

These developer guides show you how to build a different kind of Rig agent:

* **Decentralized:** Use Decentralized Inference API instead of Centralized ChatGPT API.
* **Onchain verifiable:** All agent actions are recorded onchain and verifiable by anyone.
* **Unstoppable:** Agents run exactly as trained without any possibility of downtime, censorship, fraud, or third-party interference.
* **Intelligent:** Give your Rig new superpowers with state-of-the-art onchain AI models such as Llama 3.1 405B, and Hermes 3 70B.

### Step 1: Get the code

```bash
git clone https://github.com/0xPlaygrounds/rig.git
```

Rig is a Rust library for building scalable, modular, and ergonomic LLM-powered applications. Eternal AI has already been built into the Rig library. You can pull the Rig source code and start building with Eternal AI.

### Step 2: Export ETERNALAI\_API\_KEY environment variable

```
export ETERNALAI_API_KEY="your-eternalai-api-key"
```

You can get the api key [here](https://eternalai.org/api).

### Step 3: Update model and chain ID (if needed)

We've implemented a simple LLM-powered agent using Rig in \`rig-core/examples/agent\_with\_eternalai.rs\`.

By default, it uses **Symbiosys** chain and **Hermes 3 70B**  model.&#x20;

You can find the list of supported chains and models [here](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/onchain-models).&#x20;

For this tutorial, we'll use **Base** and **Hermes 3 70B,** so we need to update the chain ID in the agent\_with\_eternalai.rs from 45762 to 8453.&#x20;

<figure><img src="../../.gitbook/assets/image (65).png" alt=""><figcaption></figcaption></figure>

### Step 4: Build your Rig agent

Run the following commands to install dependencies.

```
cargo update
```

### Step 5: Run your Rig agent

```
run --package rig-core --example agent_with_eternalai
```

After that you can see the following outputs in your terminal.

```
TODO
```

### Step 6: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain verifiable.

Let's look at an onchain prompt transaction on BaseScan. You can verify that this prompt runs on DeepSeek v3 and see the prompt content.

{% embed url="https://basescan.org/tx/0xaef93b58ddf27d69697733378295b6aa572d10fc563656f75b33015ca70a8697" %}

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

### Step 6: Verify the onchain response transaction

Now, let's look at the onchain response tx on BaseScan. You can see the actual response content. Everything is onchain and verifiable.

{% embed url="https://basescan.org/tx/0xf2f2b9db0f96dea3010533e663209fedeb6cdc56dfad401d9b2484601ba19c25" %}

<figure><img src="../../.gitbook/assets/image (2).png" alt=""><figcaption></figcaption></figure>

### Step 7 (Advanced): Reproduce the response

Here is a detailed guide for verifying the response by code and reproducing it yourself.

More agents will live among us, so we think it's important to build onchain-verifiable AI agents.

As with anything in crypto — "Don't trust the AI, verify it."

{% content-ref url="dont-trust-verify.md" %}
[dont-trust-verify.md](dont-trust-verify.md)
{% endcontent-ref %}



### Step 8 (Advanced): Read the whitepaper

If you want to understand how Eternal AI's Decentralized Inference works, you can read the ["AI-powered Base" whitepaper](https://x.com/punk3700/status/1869428187450749093).

<figure><img src="../../.gitbook/assets/image (3).png" alt=""><figcaption></figcaption></figure>

### Congrats!

Congrats! You've finished building your Eliza agent with two superpowers:&#x20;

* Onchain verifiable on Base
* Powered by Hermes 3 70B

Questions? Join the Eternal AI Devs group on Telegram: [https://t.me/EternalAIDevs](https://t.me/EternalAIDevs).
