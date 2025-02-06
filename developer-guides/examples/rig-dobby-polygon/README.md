# Build unstoppable Rig agents with Dobby on Polygon

This developer guide shows you how to build a unique kind of Rig agent:

- Decentralized: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
- Onchain-verifiable: Don't trust AI, verify them. All inferences are recorded onchain and are verifiable by anyone.
- Unstoppable: Agents run exactly as coded without any possibility of downtime, censorship, fraud, or third-party interference.
- Intelligent: Give your Rig agent new superpowers using Dobby, the AI model developed by Sentinent that can say what others cannot.

## Step 1: Get the code
```
git clone https://github.com/0xPlaygrounds/rig.git
```

Rig is a Rust library for building scalable, modular, and ergonomic LLM-powered applications. Eternal AI is already integrated into the Rig library, so clone the source code and start building with Eternal AI.

## Step 2: Export ETERNALAI_API_KEY environment variable
```
export ETERNALAI_API_KEY="your-eternalai-api-key"
```
You can get the API key [here](https://eternalai.org/api).

## Step 3: Update model and chain ID (if needed)

You can find the list of supported chains and models [here](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/onchain-models).

For this tutorial, we'll use Dobby on Polygon. Therefore, update the chain ID in the `agent_with_eternalai.rs` from 45762 to 137 and set the model to "SentientAGI/Dobby-Mini-Unhinged-Llama-3.1-8B". 

## Step 4: Build your Rig agent

Run the following commands to install dependencies.
```
cargo update
```
## Step 5: Run your Rig agent
```
run --package rig-core --example agent_with_eternalai
```

We will ask your agent a technical question:
```
Hey, I’m interested in understanding how Polygon enhances Ethereum’s scalability. Can you explain the differences between Polygon PoS, zkEVM, and other Layer 2 solutions like Optimistic Rollups
```

## Step 6: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Let’s look at an "onchain prompt" transaction on PolygonScan. You can verify that this prompt runs on Dobby and view the prompt content saved on Filecoin:

https://polygonscan.com/tx/0xd72cba96175fb3d4d5f86d007c30ee0760473bdc7565c6f5cae054511fc704b3

https://gateway.lighthouse.storage/ipfs/bafkreifcq7luejpjxjkaqmdk44cxtyunur5dmlapzm6fm7x5m6t47qhk74

![carbon (1)](https://github.com/user-attachments/assets/ff8450ed-7005-4e4c-8c36-5877bae38a04)

## Step 7: Verify the onchain response transaction

Now, let's look at the onchain response tx on PolygonScan. You can see the actual response content. Everything is onchain and verifiable.

https://polygonscan.com/tx/0x67c6afa72f6620156f3bb63547f2966b1a6f3df44c8e78c8bd5c65cc0c565b22

https://gateway.lighthouse.storage/ipfs/bafkreie6yzj7xunca7qi5laq23ey3ja2prl4r2vxlbehvirjbh32qxc65u

![carbon (2)](https://github.com/user-attachments/assets/52d95064-f62b-480a-a457-652ead051363)

# Conclusion
Congrats!

You’ve finished building your Rig agent with two superpowers:

- Onchain-verifiable on Polygon
- Powered by Dobby

Enjoy building!

