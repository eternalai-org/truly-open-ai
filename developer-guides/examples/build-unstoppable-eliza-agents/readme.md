# Build unstoppable Eliza agents with DeepSeek-R1 on Base
### This developer guide shows you how to build a different kind of Eliza agent:
- Decentralized: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
- Onchain-verifiable: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
- Unstoppable: Agents run exactly as coded without any possibility of downtime, censorship, fraud, or third-party interference.
- Intelligent: Give your Eliza new superpowers using DeepSeek R1, the state-of-the-art onchain AI model for reasoning.
## Step 1: Get your Decentralized Inference API key
- Go to http://eternalai.org/api
- Login with X (Twitter)
- Get the Decentralized Inference API key for Base<br/>

It's free.<br/>

Here’s a detailed document for those curious about how Eternal AI's Decentralized Inference works with DeepSeek R1 onchain on Base:
https://github.com/eternalai-org/truly-open-ai/tree/master/developer-guides/examples/decentralized-inference-api-deepseek-r1-base

## STEP 2: Get the Code
> git clone https://github.com/ai16z/eliza.git

Today, Eliza framework is the easiest way to spin up your own AI agent. Eternal AI is already integrated into the Eliza framework. You can pull the Eliza source code and start building with Eternal AI.

## STEP 3: Create Your .env File
> cp .env.example .env

Update the .env file (lines 19-21):

> ETERNALAI_URL=https://api.eternalai.org/v1/  
> ETERNALAI_MODEL="DeepSeek R1”  
> ETERNALAI_API_KEY=YOUR_API_KEY  

## STEP 4: Build and Run Your Eliza Agent
> pnpm clear  
> pnpm i  
> pnpm build  
> pnpm start --characters="YOUR_CHARACTER_FILE"

## STEP 5: Verify the Onchain Prompt Transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Let’s look at an "onchain prompt" transaction on Base Explorer. You can verify that this prompt runs on DeepSeek R1 and see the prompt content.

https://basescan.org/tx/0xf14cb3ec325ec79b22559f2d4eeac7dffde9b077cfd80e559aeb81b0d96a46f8

![Screenshot 2025-01-22 at 14 14 36](https://github.com/user-attachments/assets/fdd29aac-79a1-4475-be7b-39cc7c778ec0)

## STEP 6: Verify the Onchain Response Transaction

Now, let’s look at the "onchain response" transaction on Base Explorer.

You can see the actual response content. Everything is onchain and verifiable.

https://basescan.org/tx/0x31129d0b2d0ecc7affa9a218c683ab4035baa1139256836c1140d661ac53139d

![Screenshot 2025-01-22 at 14 07 26](https://github.com/user-attachments/assets/2dac4948-18f8-4ce5-9d61-3184c711f7ab)


## STEP 7: Don't Trust AI Agents, Verify Them

If you prefer to verify the responses by code and reproduce the responses from the prompts, here’s a detailed guide:

https://x.com/CryptoEternalAI/status/1869611077195161669

# Conclusion

Congrats!

You’ve finished building your Eliza agent with two superpowers:

- Onchain-verifiable on Base
- Powered by Deepseek R1

Enjoy building!

