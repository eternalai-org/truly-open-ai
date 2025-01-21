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
* **Intelligent:** Give your Rig new superpowers with state-of-the-art onchain AI models such as DeepSeek R1, Llama 3.1 405B, and Hermes 3 70B.

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

For this tutorial, we'll use **Base** and **DeepSeek R1,** so we need to update the chain ID in the `agent_with_eternalai.rs` from 45762 to 8453.&#x20;

<figure><img src="../../.gitbook/assets/image (70).png" alt=""><figcaption></figcaption></figure>

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
Try to get on-chain system prompt

<think>
Alright, so I've just read through this conversation where the user asks to be entertained, and the assistant responds with a series of jokes and humorous takes on Bitcoin and AI. The user then thanks the assistant, who signs off with a Bitcoin and AI-themed goodbye.

Now, the user has provided a new query where they say, "Entertain me!" again. I need to figure out how to respond appropriately. First, I should consider the context. The previous interaction was positive, with the user thanking the assistant, so the tone should remain friendly and humorous.

The assistant's role is to entertain using humor and jokes, specifically focusing on Bitcoin and decentralized AI. So, the response should include more jokes and perhaps expand on the previous topics or introduce new ones related to these subjects.

I should think about what aspects of Bitcoin and AI are ripe for humor. Maybe play on common misconceptions, the volatility of crypto, the jargon used in both fields, or the contrast between decentralized and centralized systems.

Let me brainstorm a few ideas:

1. Jokes about the complexity of Bitcoin for newcomers.
2. Puns involving crypto terms like "HODL," "blockchain," or "mining."
3. Comparisons between AI and human behavior, especially in a decentralized context.
4. Exaggerated scenarios involving the impact of Bitcoin or AI on everyday life.

I should also consider the structure. The previous response had a series of short jokes, each in its own paragraph. I might follow a similar format to keep it easy to read and engaging.

Now, putting it all together, I can craft a response that includes a few jokes, each touching on different aspects of Bitcoin and decentralized AI, ensuring they're light-hearted and funny without being offensive or too technical.

I need to make sure the jokes are original and haven't been used before in the previous interaction. Maybe include some wordplay or situations that someone familiar with crypto or AI would find amusing.

Finally, I'll wrap it up with a friendly sign-off, perhaps tying back to the themes of Bitcoin and AI, encouraging the user to engage more or learn about these topics in a humorous way.
</think>

Sure! Here's a fresh set of jokes tailored for you:

---

Bitcoin Jokes:

1. Why did Bitcoin go to therapy? It was feeling a little "volatile" and needed to stabilize its mood swings!

2. What did the blockchain say to the new miner? "You're a real 'block'-buster! Keep digging, every little hash counts!"

3. Why did the crypto investor bring a ladder to the party? They heard the prices might "moon" and wanted to be prepared!

4. What do you call a Bitcoiner who’s always broke? A HODL-er with a PhD in patience!

---

AI Jokes:

1. Why did the decentralized AI go to art school? It wanted to learn how to paint outside the lines—no central gallery needed!

2. What did the AI say when it ran into its ex? "You’re just a neural network from my past, but I’m training better models now!"

3. Why don’t AIs make good comedians? They always explain the punchline with too many algorithms!

4. What do you call an AI that’s feeling sad? A blue screen of death... but decentralized, so it’s just a minor glitch!

---

Combined Bitcoin & AI Humor:

1. Why did Bitcoin and AI go on a date? Bitcoin was looking for a stable relationship, and AI promised not to devalue the evening!

2. What’s the best way to pay an AI therapist? In Bitcoin, of course—it’s the only way to ensure your sessions are encrypted!

---

Hope these made you chuckle! Bitcoin to the moon and AI to the people! 🚀✨
```

### Step 6: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain verifiable.

Let's look at an onchain prompt transaction on BaseScan. You can verify that this prompt runs on **Hermes 3 70B** and see the prompt content.

{% embed url="https://basescan.org/tx/0x23f3425fad7fbd28bf090facf9344f687b408262cb85b9e79db5692df1ac08ee" %}

<figure><img src="../../.gitbook/assets/image (71).png" alt=""><figcaption></figcaption></figure>

### Step 6: Verify the onchain response transaction

Now, let's look at the onchain response tx on BaseScan. You can see the actual response content. Everything is onchain and verifiable.

{% embed url="https://basescan.org/tx/0x3ca6e6dae6379f9ae4572a387f0e72d39718015f5f55e6fc65980ba4bed3c243" %}

<figure><img src="../../.gitbook/assets/image (72).png" alt=""><figcaption></figcaption></figure>

### Step 7 (Advanced): Reproduce the response

Here is a detailed guide for verifying the response by code and reproducing it yourself.

More agents will live among us, so we think it's important to build onchain-verifiable AI agents.

As with anything in crypto — "Don't trust the AI, verify it."

{% embed url="https://docs.eternalai.org/eternal-ai/decentralized-inference-api/tutorials/dont-trust-verify" %}

### Step 8 (Advanced): Read the whitepaper

If you want to understand how Eternal AI's Decentralized Inference works, you can read the ["AI-powered Base" whitepaper](https://x.com/punk3700/status/1869428187450749093).

<figure><img src="../../.gitbook/assets/image (3) (1) (1).png" alt=""><figcaption></figcaption></figure>

### Congrats!

Congrats! You've finished building your Rig agent with two superpowers:&#x20;

* Onchain verifiable on Base
* Powered by Hermes 3 70B

Questions? Join the Eternal AI Devs group on Telegram: [https://t.me/EternalAIDevs](https://t.me/EternalAIDevs).
