# Build unstoppable Rig agents with DeepSeek-R1 on Base

This developer guide shows you how to build a different kind of Rig agent:

- Decentralized: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
- Onchain-verifiable: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
- Unstoppable: Agents run exactly as coded without any possibility of downtime, censorship, fraud, or third-party interference.
- Intelligent: Give your Eliza new superpowers using DeepSeek-R1, the state-of-the-art onchain AI model for reasoning.

## Step 1: Get the code
```
git clone https://github.com/0xPlaygrounds/rig.git
```

Rig is a Rust library for building scalable, modular, and ergonomic LLM-powered applications. Eternal AI has already been built into the Rig library. You can pull the Rig source code and start building with Eternal AI.

## Step 2: Export ETERNALAI_API_KEY environment variable
```
export ETERNALAI_API_KEY="your-eternalai-api-key"
```
You can get the API key [here](https://eternalai.org/api).

## Step 3: Update model and chain ID (if needed)

You can find the list of supported chains and models [here](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/onchain-models).

For this tutorial, we'll use Base and DeepSeek-R1, so we need to update the chain ID in the `agent_with_eternalai.rs` from 45762 to 8453. 

![image](https://github.com/user-attachments/assets/da26890b-3dd0-479d-938d-0373b393993b)

## Step 4: Build your Rig agent

Run the following commands to install dependencies.
```
cargo update
```
## Step 5: Run your Rig agent
```
run --package rig-core --example agent_with_eternalai
```

We will ask your agent a math-related question:
```
A bag contains 3 red balls, 5 blue balls, and 2 green balls. If you randomly pick two balls without replacement, what is the probability that both are blue?
```

## Step 6: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Let’s look at an "onchain prompt" transaction on Base Explorer. You can verify that this prompt runs on DeepSeek-R1 and see the prompt content.

https://basescan.org/tx/0x784b94ed150d85d9a04c3e4a52fa52bbd2af48a7a6070c7d6782459f5ced666a

![Screenshot 2025-01-22 at 16 43 01](https://github.com/user-attachments/assets/28c248f1-c635-4539-8606-afa6ec805105)

## Step 7: Let's see how the DeepSeek-R1 thinks

When DeepSeek-R1 receives a prompt, it first thinks about the reasoning process in its mind.
```
<think>
First, I need to determine the total number of balls in the bag by adding up the red, blue, and green balls: 3 + 5 + 2 = 10 balls in total.

Next, I'll calculate the number of ways to pick any two balls out of the 10. This is done using the combination formula, which is 10 choose 2, resulting in 45 possible pairs.

Then, I'll find the number of favorable outcomes, which are the ways to pick two blue balls. Since there are 5 blue balls, the number of ways to choose 2 is 5 choose 2, giving 10 favorable pairs.

Finally, the probability of picking two blue balls is the ratio of favorable outcomes to the total number of possible outcomes, which is 10/45. Simplifying this fraction gives the final probability.
</think>
```
## Step 8: Let's see how DeepSeek-R1 answers

Based on its prior reasoning in Step 6, DeepSeek-R1 provides the final answer.
```
Sure, let's solve the problem step by step.

**Problem:**
A bag contains:
- 3 red balls
- 5 blue balls
- 2 green balls

If you randomly pick two balls without replacement, what is the probability that both are blue?

**Solution:**

1. **Total number of balls in the bag:**
   
   \[
   \text{Total balls} = \text{Red} + \text{Blue} + \text{Green} = 3 + 5 + 2 = 10 \text{ balls}
   \]

2. **Number of ways to choose 2 balls out of 10:**
   
   \[
   \text{Total ways} = \binom{10}{2} = \frac{10!}{2!(10-2)!} = \frac{10 \times 9}{2} = 45
   \]

3. **Number of ways to choose 2 blue balls out of 5:**
   
   \[
   \text{Favorable ways} = \binom{5}{2} = \frac{5!}{2!(5-2)!} = \frac{5 \times 4}{2} = 10
   \]

4. **Probability of picking 2 blue balls:**
   
   \[
   \text{Probability} = \frac{\text{Favorable ways}}{\text{Total ways}} = \frac{10}{45} = \frac{2}{9}
   \]

**Final Answer:**

\[
\boxed{\dfrac{2}{9}}
\]
```

## Step 9: Verify the onchain response transaction

Now, let's look at the onchain response tx on BaseScan. You can see the actual response content. Everything is onchain and verifiable.

https://basescan.org/tx/0xcb7ca61b77a245121c97918d3387d7b5c80b1fd25aa946532e7269eb08b6f3e5

![Screenshot 2025-01-22 at 17 02 38](https://github.com/user-attachments/assets/4d55cb96-39a7-40b0-b3cb-a44e65df00d8)

# Conclusion
Congrats!

You’ve finished building your Rig agent with two superpowers:

- Onchain-verifiable on Base
- Powered by Deepseek R1

Enjoy building!
