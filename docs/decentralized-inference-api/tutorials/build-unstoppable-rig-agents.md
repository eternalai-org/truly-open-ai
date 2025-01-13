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

We've implemented a simple LLM-powered agent using Rig in `rig-core/examples/agent_with_eternalai.rs`.

By default, it uses **Symbiosys** chain and **Hermes 3 70B**  model.&#x20;

You can find the list of supported chains and models [here](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/onchain-models).&#x20;

For this tutorial, we'll use **Base** and **Hermes 3 70B,** so we need to update the chain ID in the `agent_with_eternalai.rs` from 45762 to 8453.&#x20;

<figure><img src="../../.gitbook/assets/image (3).png" alt=""><figcaption></figcaption></figure>

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
Running basic agent with eternalai

onchain_data: {
  "infer_id": "12202",
  "infer_tx": "0x409089de35d5529a8ef2ee8af93d71a1eef7c774e27bfd939c0bffdffcdef82a",
  "input_cid": "",
  "output_cid": "",
  "pbft_committee": [
    "0x78fabfdca5ac093af921f3bc17ee419e671ae34f",
    "0x7f1b49493d8477a2172984b30820c61fd1b15427",
    "0x02dba3a5dc107665d4f00e0ac2d414e43acc6f16"
  ],
  "propose_tx": "0x55756e7b4368bd39b7e7e16741361f4d4871836f93a67a14d98c88ff2b5dce59",
  "proposer": "0x02dba3a5dc107665d4f00e0ac2d414e43acc6f16"
}

Sure, here are a few jests to tickle your funny bone!

1. I just got a job at a bakery because I kneaded dough.

2. How do mathematicians scold their children? "If I've told you once, I've told you n times..."

3. I'm reading a book about anti-gravity. It's impossible to put down.

4. What do you call a dinosaur that's a noisy sleeper? A Brontosnorus.

5. Did you hear about the mathematician who's afraid of negative numbers? He will stop at nothing to avoid them.

6. Why don't scientists trust atoms? Because they make up everything.

7. I used to work at an orange juice factory... I was sacked because I couldn't concentrate.

8. My friend asked me what the trick was to folding an inflatable raft. I said, "Aren't you just full of hot air?"

I hope these gave you a chuckle or two!
```

### Step 6: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain verifiable.

Let's look at an onchain prompt transaction on BaseScan. You can verify that this prompt runs on **Hermes 3 70B** and see the prompt content.

{% embed url="https://basescan.org/tx/0x409089de35d5529a8ef2ee8af93d71a1eef7c774e27bfd939c0bffdffcdef82a" %}

<figure><img src="../../.gitbook/assets/image (1).png" alt=""><figcaption></figcaption></figure>

### Step 6: Verify the onchain response transaction

Now, let's look at the onchain response tx on BaseScan. You can see the actual response content. Everything is onchain and verifiable.

{% embed url="https://basescan.org/tx/0x55756e7b4368bd39b7e7e16741361f4d4871836f93a67a14d98c88ff2b5dce59" %}

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

<figure><img src="../../.gitbook/assets/image (3) (1).png" alt=""><figcaption></figcaption></figure>

### Congrats!

Congrats! You've finished building your Eliza agent with two superpowers:&#x20;

* Onchain verifiable on Base
* Powered by Hermes 3 70B

Questions? Join the Eternal AI Devs group on Telegram: [https://t.me/EternalAIDevs](https://t.me/EternalAIDevs).
