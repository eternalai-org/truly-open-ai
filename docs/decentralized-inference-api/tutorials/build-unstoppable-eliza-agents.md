---
cover: ../../.gitbook/assets/eliza_banner.jpg
coverY: 0
layout:
  cover:
    visible: true
    size: hero
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

# Build unstoppable Eliza agents

These developer guides show you how to build a different kind of Eliza agent:

* **Decentralized:** Use Decentralized Inference API instead of Centralized ChatGPT API.
* **Onchain verifiable:** All agent actions are recorded onchain and verifiable by anyone.
* **Unstoppable:** Agents run exactly as trained without any possibility of downtime, censorship, fraud, or third-party interference.
* **Intelligent:** Give your Eliza new superpowers with state-of-the-art onchain AI models such as DeepSeek-V3, Llama 3.1 405B, and Hermes 3 70B.

### Step 1: Get the code

```bash
git clone https://github.com/ai16z/eliza.git
```

Today, the Eliza framework is the easiest way to spin up your own AI agent. Eternal AI has already been built into the Eliza framework. You can pull the Eliza source code and start building with Eternal AI.

### Step 2: Create your .env file

Copy `.env.example` to `.env`

```bash
cp .env.example .env
```

Update the .env file.

```
ETERNALAI_URL=https://api.eternalai.org/v1/
ETERNALAI_MODEL='DeepSeek V3'
ETERNALAI_API_KEY=
```

For the ETERNALAI\_API\_KEY, you can get it [here](https://eternalai.org/api).

For the ETERNALAI\_MODEL, you can find the list of supported chains and models [here](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/onchain-models).&#x20;

For this tutorial, we'll use **Base** and **DeepSeek V3**.

<figure><img src="../../.gitbook/assets/image (4).png" alt=""><figcaption></figcaption></figure>

### Step 3: Build your Eliza agent

Run the following commands to install and build the source code:

```bash
pnpm clear
pnpm i
pnpm build
```

### Step 4: Run your Eliza agent

Start with `eternalai.character.json`  or any character configuration supported by the Eliza framework.&#x20;

{% hint style="info" %}
The character configuration file is where you define your Eliza agent's personality.
{% endhint %}

```bash
pnpm start --characters="characters/eternalai.character.json"
```

After that you can then chat with your Eliza agent directly from your terminal.

```
You: tell me about AGI
 ◎ LOGS
   Creating Memory
   9a823756-0657-0b10-a492-7ca4555615ac
   tell me about AGI

 ["◎ Generating message response.."]

 ["◎ Generating text..."]

 ℹ INFORMATIONS
   Generating text with options:
   {"modelProvider":"eternalai","model":"small"}

 ℹ INFORMATIONS
   Selected model:
   DeepSeek V3

 ◎ LOGS
   Creating Memory

   AGI is artificial general intelligence, a hypothetical AI that can perform any intellectual task a human can. it’s not here yet, but it’s coming.

 ◎ LOGS
   Evaluating
   GET_FACTS

 ◎ LOGS
   Evaluating
   UPDATE_GOAL

 ["✓ Normalized action: none"]

 ["ℹ Executing handler for action: NONE"]

 ["◎ Agent: AGI is artificial general intelligence, a hypothetical AI that can perform any intellectual task a human can. it’s not here yet, but it’s coming."]
```

### Step 5: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain verifiable.

Let's look at an onchain prompt transaction on BaseScan. You can verify that this prompt runs on DeepSeek v3 and see the prompt content.

{% embed url="https://basescan.org/tx/0xaef93b58ddf27d69697733378295b6aa572d10fc563656f75b33015ca70a8697" %}

<figure><img src="../../.gitbook/assets/image (1) (1).png" alt=""><figcaption></figcaption></figure>

### Step 6: Verify the onchain response transaction

Now, let's look at the onchain response tx on BaseScan. You can see the actual response content. Everything is onchain and verifiable.

{% embed url="https://basescan.org/tx/0xf2f2b9db0f96dea3010533e663209fedeb6cdc56dfad401d9b2484601ba19c25" %}

<figure><img src="../../.gitbook/assets/image (2) (1).png" alt=""><figcaption></figcaption></figure>

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
* Powered by DeepSeek V3

Questions? Join the Eternal AI Devs group on Telegram: [https://t.me/EternalAIDevs](https://t.me/EternalAIDevs).
