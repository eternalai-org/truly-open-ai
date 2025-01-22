# Build unstoppable Eliza agents with DeepSeek-R1 on Base

This developer guide shows you how to build a different kind of Eliza agent:
- Decentralized: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
- Onchain-verifiable: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
- Unstoppable: Agents run exactly as coded without any possibility of downtime, censorship, fraud, or third-party interference.
- Intelligent: Give your Eliza new superpowers using DeepSeek-R1, the state-of-the-art onchain AI model for reasoning.

## Step 1: Get your Decentralized Inference API key
- Go to http://eternalai.org/api
- Login with X (Twitter)
- Get the Decentralized Inference API key for Base<br/>

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
ETERNALAI_MODEL="DeepSeek-R1-Distill-Llama-70B"  
ETERNALAI_API_KEY=YOUR_API_KEY  
```

## Step 4: Build and run your Eliza agent
```
pnpm clear  
pnpm i  
pnpm build  
pnpm start --characters="YOUR_CHARACTER_FILE"
```

We will ask DeepSeek a reasoning question:
```
Three people stand in a line. Each has a hat that is either red or blue, but they don’t know their own hat color. The first person says, “I don’t know my hat color.” The second person says, “I don’t know my hat color.” The third person immediately says, “I know my hat color.” What is the color of the third person’s hat? Why?
```


## Step 5: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Let’s look at an "onchain prompt" transaction on Base Explorer. You can verify that this prompt runs on DeepSeek R1 and see the prompt content.

https://basescan.org/tx/0xf14cb3ec325ec79b22559f2d4eeac7dffde9b077cfd80e559aeb81b0d96a46f8

![Screenshot 2025-01-22 at 14 14 36](https://github.com/user-attachments/assets/fdd29aac-79a1-4475-be7b-39cc7c778ec0)

## Step 6: Let's see how the DeepSeek-R1 thinks

When DeepSeek-R1 receives a prompt, it first thinks about the reasoning process in its mind.
```
<think>
Alright, so the user presented a logic puzzle about three people with red or blue hats. They each say they don't know their hat color, except the third person who figures it out. I need to solve this and explain why the third person knew their hat color.

First, let's break it down. The hats can only be red or blue. The people are in a line, so the first person sees the second and third, the second sees only the third, and the third sees none.

When the first person says they don't know their hat color, that tells the others that they don't see two hats of the same color. If they did, they'd know their own hat must be the other color. So, the first person's statement implies that the second and third don't both have the same color.

Next, the second person also says they don't know their hat color. The second person, hearing the first person's statement, knows that the third person's hat isn't the same as theirs. If the second person saw that the third had a red hat, for example, and knowing the first person didn't see two of the same, they could deduce their own hat must be blue if the third is red, or vice versa. But since the second person still doesn't know, it means the third person must have a hat color that doesn't allow the second to determine theirs.

This leads the third person to realize that their hat must be a specific color. If the third person saw that the second person had a red hat, and knowing the second couldn't determine their own, the third would deduce their hat must be blue. Similarly, if the third saw a blue hat on the second, their own must be red. 

So, putting it all together, the third person's hat color is the opposite of the second person's. Therefore, the third person can confidently state their hat color based on the previous statements and the logic of elimination.
</think>
```
## Step 7: Let's see how DeepSeek-R1 answers

Based on its prior reasoning in Step 6, DeepSeek-R1 provides the final answer.

```
<answer>
Third person's hat is blue. Here's why: If third had red, second would've seen first's uncertainty and known their own hat must be blue. Since second also didn't know, third deduced theirs was blue.", "action": "NONE" }
</answer>
```

## Step 8: Verify the onchain response transaction

Now, let’s look at the "onchain response" transaction on Base Explorer.

You can see the actual response content. Everything is onchain and verifiable.

https://basescan.org/tx/0x31129d0b2d0ecc7affa9a218c683ab4035baa1139256836c1140d661ac53139d

![Screenshot 2025-01-22 at 14 07 26](https://github.com/user-attachments/assets/2dac4948-18f8-4ce5-9d61-3184c711f7ab)

# Conclusion

Congrats!

You’ve finished building your Eliza agent with two superpowers:

- Onchain-verifiable on Base
- Powered by Deepseek R1

Enjoy building!
