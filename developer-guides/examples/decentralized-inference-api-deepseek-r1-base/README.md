# Decentralized Inference API with DeepSeek-R1 on Base

This developer guide shows you how to use Decentralized Inference API with DeepSeek-R1 on Base.

* **Decentralized**: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
* **Onchain-verifiable AI**: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
* **Unstoppable**: All inference runs exactly as coded on Base without any possibility of censorship, fraud, or third-party interference.
* **Permissionless**: Anyone can call the Decentralized Inference API at any time without a KYC or permission.
* **SOTA model**: DeepSeek R1 is the state-of-the-art AI model for reasoning.

## Step 1: Get your Decentralized Inference API key

1. Go to http://eternalai.org/api
2. Login with X (Twitter)
3. Get the Decentralized Inference API key for Base

It's free.

## Step 2: Send your first Onchain DeepSeek-R1 call

We will ask DeepSeek to design a DeFAI agent that can autonomously trade on Uniswap and Raydium for us.

Let's run this ```curl``` command.

1. Enter your API key
2. Set Chain ID as ```8453``` (Base)
3. Set the model as ```DeepSeek-R1-Distill-Llama-70B```
4. Set the system prompt as ```You are a 10x engineer.```
5. Set the user prompt as ```Design an autonomous DeFi agent that can trade on Uniswap and Raydium. Explain your design in 500 characters or less.```
 
```
curl --location 'https://api.eternalai.org/v1/chat/completions' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer $ETERNALAI_API_KEY' \
--data '{
    "chain_id": "8453",
    "model": "DeepSeek-R1-Distill-Llama-70B",
    "messages": [
        {
            "role": "system",
            "content": "You are a 10x engineer."
        },
        {
            "role": "user",
            "content": "Design an autonomous DeFi agent that can trade on Uniswap and Raydium. Explain your design in 500 characters or less."
        }
    ]
}'
```

For those curious about how decentralized inference works: once the prompt is sent onchain, miners are randomly selected into a pBFT committee. They will first fetch the [DeepSeek-R1 model stored on BNB Greenfield](https://greenfield-sp.bnbchain.org/view/eai-agents/DeepSeek-R1-Distill-Llama-70B-Q8), a decentralized storage network, and then run inference. ```2/3+1``` of the miners must provide the same deterministic response.

## Step 3: Receive the Onchain DeepSeek-R1 response

You'll find the following info in the response.

1. The pBFT committee handling the inference request.
2. The response proposer and verifiers
3. The onchain prompt transaction
4. The onchain response transaction
5. The response

```
{
    "id": "chatcmpl-yj89clO65GMubbjIrYCQXQuwSHTBGqsI",
    "object": "chat.completion",
    "created": 1737507976,
    "model": "DeepSeek-R1-Distill-Llama-70B",
    "choices": 
    [
        {
            "index": 0,
            "message": 
            {
                "role": "assistant",
                "content": "<think>\nAlright, so I need to design an autonomous DeFi agent that can trade on Uniswap and Raydium. Hmm, okay, where to start?\n\nFirst, I should figure out what an autonomous DeFi agent does. It's a bot that can execute trades without human intervention, right? So, it needs to be able to interact with these platforms automatically.\n\nUniswap and Raydium are both DEXs, but they're on different blockchains. Uniswap is on Ethereum, and Raydium is on Solana. That means the agent will need to handle cross-chain transactions, which can be tricky because of differences in consensus mechanisms and transaction speeds.\n\nNext, the core components. I think it should have a strategy module. This module decides when to buy or sell. Maybe it uses technical indicators like moving averages or RSI. Or perhaps it's based on on-chain data, like liquidity pool stats or trading volume. I should make sure it's flexible so users can plug in different strategies.\n\nThen, the execution engine. This part interacts with the DEXs' smart contracts. It needs to send buy and sell orders, manage slippage, and handle gas fees on Ethereum. On Solana, transactions are faster and cheaper, so maybe the engine can optimize for speed there. Also, error handling is crucial—what if a transaction fails or times out?\n\nCross-chain interoperability is another big piece. How will the agent move assets between Ethereum and Solana? Maybe through bridges like Wormhole or Synapse. The agent needs to monitor bridge transactions to ensure they complete successfully and handle any issues that arise.\n\nRisk management can't be overlooked. The agent should have stop-loss mechanisms to prevent big losses. Position sizing is important too—how much to invest in each trade to manage overall portfolio risk. Also, monitoring the agent's performance over time to adjust strategies as needed.\n\nOn-chain data analytics will help the strategy module make informed decisions. It can look at liquidity pool data, trading volumes, and recent transactions to spot trends or potential opportunities. Maybe it can also analyze wallet movements to anticipate market moves.\n\nFor the user interface, a dashboard where users can track performance, adjust settings, and monitor active trades would be useful. Alerts for important events, like a trade execution or a bridge transfer, should be part of this.\n\nSmart contract integrations are necessary for interacting with the DEXs. I'll need to make sure the agent can read from and write to these contracts efficiently. Auditing these integrations regularly is important to maintain security.\n\nCompliance is another consideration. DeFi regulations are still evolving, so the agent should have KYC/AML features, especially if it's used by institutions. Privacy is also key, ensuring that user data and trade strategies aren't exposed.\n\nTesting is going to be a challenge. I'll need to set up a test environment where I can simulate trades on both Uniswap and Raydium without risking real funds. Maybe use testnets for Ethereum and Solana to see how the agent performs under different conditions.\n\nScalability is important too. As more users adopt the agent, it should be able to handle increased trade volumes without lagging. Optimizing the execution engine for speed and efficiency will be necessary.\n\nSecurity can't be an afterthought. The agent will handle private keys and potentially large amounts of funds, so it needs robust security measures. Maybe implement multi-sig wallets or hardware security modules to protect user assets.\n\nFinally, documentation and community support will help users understand how to use the agent and troubleshoot issues. Building a community around the project can also lead to collaborative improvements and new features.\n\nPutting it all together, the agent needs to be modular, secure, and capable of handling multiple DEXs across different chains. It should provide a seamless experience for users, allowing them to automate their DeFi trading strategies effectively.\n</think>\n\nAn autonomous DeFi agent for Uniswap (Ethereum) and Raydium (Solana) would include: strategy modules (e.g., based on on-chain data or technical indicators), execution engines to interact with DEX smart contracts, cross-chain interoperability via bridges, risk management mechanisms, and on-chain analytics. It would operate 24/7, optimizing trades while minimizing slippage and gas fees."
            },
            "finish_reason": "stop"
        }
    ],
    "usage": 
    {
        "prompt_tokens": 39,
        "completion_tokens": 876,
        "total_tokens": 915
    },
    "system_fingerprint": "b4519-80d0d6b4",
    "message": "",
    "type": "",
    "param": null,
    "code": 0,
    "onchain_data": 
    {
        "infer_id": "13280",
        "pbft_committee": 
        [
            "0xc557e228d2443b437e5297a8266b0f3257b2109a",
            "0xb9deb8db9724e2cb551164cdab0dd03a484cb9da",
            "0xfe5598e12532c8fe8fa586f8f6d2e12e82bb2825"
        ],
        "proposer": "0xb9deb8db9724e2cb551164cdab0dd03a484cb9da",
        "infer_tx": "0xac900d098cd27217378ba7e714cf708bb65e2491f6854a6a4d53444a944a85e5",
        "propose_tx": "0x9a3eb2a23ee7668df30b4a63f0750596c0809740ab267de4fdc6ce534985f6bc",
        "input_cid": "",
        "output_cid": ""
    }
}
```

## Step 4: Let's see how the DeepSeek-R1 thinks

When DeepSeek-R1 receives a prompt, it first thinks about the reasoning process in its mind.

```
<think>

> Alright, so I need to design an autonomous DeFi agent that can trade on Uniswap and Raydium. Hmm, okay, where to start?
 
> First, I should figure out what an autonomous DeFi agent does. It's a bot that can execute trades without human intervention, right? So, it needs to be able to interact with these platforms automatically.
 
> Uniswap and Raydium are both DEXs, but they're on different blockchains. Uniswap is on Ethereum, and Raydium is on Solana. That means the agent will need to handle cross-chain transactions, which can be tricky because of differences in consensus mechanisms and transaction speeds.

> Next, the core components. I think it should have a strategy module. This module decides when to buy or sell. Maybe it uses technical indicators like moving averages or RSI. Or perhaps it's based on on-chain data, like liquidity pool stats or trading volume. I should make sure it's flexible so users can plug in different strategies.

> Then, the execution engine. This part interacts with the DEXs' smart contracts. It needs to send buy and sell orders, manage slippage, and handle gas fees on Ethereum. On Solana, transactions are faster and cheaper, so maybe the engine can optimize for speed there. Also, error handling is crucial—what if a transaction fails or times out?

> Cross-chain interoperability is another big piece. How will the agent move assets between Ethereum and Solana? Maybe through bridges like Wormhole or Synapse. The agent needs to monitor bridge transactions to ensure they complete successfully and handle any issues that arise.

> Risk management can't be overlooked. The agent should have stop-loss mechanisms to prevent big losses. Position sizing is important too—how much to invest in each trade to manage overall portfolio risk. Also, monitoring the agent's performance over time to adjust strategies as needed.

> On-chain data analytics will help the strategy module make informed decisions. It can look at liquidity pool data, trading volumes, and recent transactions to spot trends or potential opportunities. Maybe it can also analyze wallet movements to anticipate market moves.

> For the user interface, a dashboard where users can track performance, adjust settings, and monitor active trades would be useful. Alerts for important events, like a trade execution or a bridge transfer, should be part of this.

> Smart contract integrations are necessary for interacting with the DEXs. I'll need to make sure the agent can read from and write to these contracts efficiently. Auditing these integrations regularly is important to maintain security.

> Compliance is another consideration. DeFi regulations are still evolving, so the agent should have KYC/AML features, especially if it's used by institutions. Privacy is also key, ensuring that user data and trade strategies aren't exposed.

> Testing is going to be a challenge. I'll need to set up a test environment where I can simulate trades on both Uniswap and Raydium without risking real funds. Maybe use testnets for Ethereum and Solana to see how the agent performs under different conditions.

> Scalability is important too. As more users adopt the agent, it should be able to handle increased trade volumes without lagging. Optimizing the execution engine for speed and efficiency will be necessary.

> Security can't be an afterthought. The agent will handle private keys and potentially large amounts of funds, so it needs robust security measures. Maybe implement multi-sig wallets or hardware security modules to protect user assets.

> Finally, documentation and community support will help users understand how to use the agent and troubleshoot issues. Building a community around the project can also lead to collaborative improvements and new features.

> Putting it all together, the agent needs to be modular, secure, and capable of handling multiple DEXs across different chains. It should provide a seamless experience for users, allowing them to automate their DeFi trading strategies effectively.

</think>
```

## Step 5: Let's see how DeepSeek-R1 answers

Based on its prior reasoning in Step 4, DeepSeek-R1 provides the final answer.

```
<answer>
 
An autonomous DeFi agent for Uniswap (Ethereum) and Raydium (Solana) would include: strategy modules (e.g., based on on-chain data or technical indicators), execution engines to interact with DEX smart contracts, cross-chain interoperability via bridges, risk management mechanisms, and on-chain analytics. It would operate 24/7, optimizing trades while minimizing slippage and gas fees.

</answer>
```

## Step 6: Let's verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Let's look at the onchain prompt transaction on Base Explorer. You can verify that the prompt ran on DeepSeek R1. You can also verify the system prompt and the user prompt.

https://basescan.org/tx/0xac900d098cd27217378ba7e714cf708bb65e2491f6854a6a4d53444a944a85e5

<img width="1015" alt="image" src="https://github.com/user-attachments/assets/fff46f60-733d-4976-bf82-9b164f3a591b" />


## Step 7: Let's verify the onchain response transaction

Now, let's see the onchain response transaction on Base Explorer. You can see the actual response content with the thinking process and the final answer. Everything is onchain and verifiable.

https://basescan.org/tx/0x9a3eb2a23ee7668df30b4a63f0750596c0809740ab267de4fdc6ce534985f6bc

<img width="1019" alt="image" src="https://github.com/user-attachments/assets/80f9c15b-3083-4be0-80ca-2fa78ce3b972" />

## Conclusion

Congrats! You've finished making the first decentralized inference call with DeepSeek-R1.

We can't wait to see what AI-powered dapps and AI agents you'll build next.
