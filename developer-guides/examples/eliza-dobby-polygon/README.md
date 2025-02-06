# Build unstoppable Eliza agents with Dobby on Polygon

This developer guide shows you how to build a different kind of Eliza agent:
- Decentralized: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
- Onchain-verifiable: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
- Unstoppable: Agents run exactly as coded without any possibility of downtime, censorship, fraud, or third-party interference.
- Intelligent: Give your Eliza new superpowers using Dobby, the AI model developed by Sentinent that can say what others cannot.

## Step 1: Get your Decentralized Inference API key
- Go to http://eternalai.org/api
- Login with X (Twitter)
- Get the Decentralized Inference API key for Polygon<br/>

It's free.<br/>

## Step 2: Get the code
```
git clone https://github.com/ai16z/eliza.git
```

Today, Eliza framework is the easiest way to spin up your own AI agent. Eternal AI is already integrated into the Eliza framework. You can pull the Eliza source code and start building with Eternal AI.

## Step 3: Create your .env file
```
cp .env.example .env
```

Update the .env file (lines 19-21):
```
ETERNALAI_URL=https://api.eternalai.org/v1/  
ETERNALAI_MODEL="SentientAGI/Dobby-Mini-Unhinged-Llama-3.1-8B"
ETERNALAI_API_KEY=YOUR_API_KEY  
```

## Step 4: Build and run your Eliza agent
```
pnpm clear  
pnpm i  
pnpm build  
pnpm start --characters="YOUR_CHARACTER_FILE"
```

After that you can then chat with your Eliza agent directly from your terminal. We will ask Eliza a technical question:

`Hey, I’m interested in understanding how Polygon enhances Ethereum’s scalability. Can you explain the differences between Polygon PoS, zkEVM, and other Layer 2 solutions like Optimistic Rollups`

## Step 5: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Let’s look at an "onchain prompt" transaction on PolygonScan. You can verify that this prompt runs on Dobby and see the prompt content saved on Filecoin.

https://polygonscan.com/tx/0x6c59d8885f732b2fef7358ae5fba15c248d90c9c4bb2cdd6482626251d2f4d47

https://gateway.lighthouse.storage/ipfs/bafkreich2whh2srldnn3glnzajm4vu3q7vpkohoqszon2posdpxpubuyoq

## Step 6: Verify the onchain response transaction

Now, let’s look at the "onchain response" transaction on PolygonScan.

You can see the actual response content. Everything is onchain and verifiable.

https://polygonscan.com/tx/0x55af635acdf03c2f206a34628d60a2a3a786a498d907534cfa6cf3b5d2cb31b5

https://gateway.lighthouse.storage/ipfs/bafkreicr657z37s3trpjxu6t665ql4wnc7ktfyogocwpt6eo4seteiyzj4

![carbon (3)](https://github.com/user-attachments/assets/26c5b778-bb0f-403f-b72f-0dda9e1cc1b6)

# Conclusion

Congrats!

You’ve finished building your Eliza agent with two superpowers:

- Onchain-verifiable on Polygon
- Powered by Dobby

Enjoy building!

