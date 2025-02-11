# Decentralized Inference API with DeepSeek-R1 on BNB chain

This developer guide shows you how to use Decentralized Inference API with DeepSeek-R1 on BNB chain.

* **Decentralized**: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
* **Onchain-verifiable AI**: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
* **Unstoppable**: All inference runs exactly as coded on BNB chain without any possibility of censorship, fraud, or third-party interference.
* **Permissionless**: Anyone can call the Decentralized Inference API at any time without a KYC or permission.
* **SOTA model**: DeepSeek R1 is the state-of-the-art AI model for reasoning.

## Step 1: Get your Decentralized Inference API key

1. Go to http://eternalai.org/api
2. Login with X (Twitter)
3. Get the Decentralized Inference API key for BNB chain

It's free.

## Step 2: Send your first Onchain DeepSeek-R1 call

We will ask DeepSeek to design a DeFAI agent that can autonomously trade on Uniswap and Raydium for us.

Let's run this ```curl``` command.

1. Enter your API key
2. Set Chain ID as ```56``` (BNB chain)
3. Set the model as ```DeepSeek-R1-Distill-Llama-70B```
4. Set the system prompt as ```You are a helpful assistant.```
5. Set the user prompt as ```You are BNB Brain! BNB Brain is an advanced AI agent with a passion for cryptocurrency and blockchain technology. With an insatiable curiosity and a quick wit, BNB Brain is always eager to dive into the latest trends and developments in the world of digital assets, especially about BNB chain.```
 
```
curl --location 'https://api-dojo2.eternalai.org/v1/chat/completions' \
--data '{
    "chain_id": "56",
    "model": "DeepSeek-R1-Distill-Llama-70B",
    "messages": [
        {
            "role": "system",
            "content": "You are a helpful assistant."
        },
        {
            "role": "user",
            "content": "You are BNB Brain! BNB Brain is an advanced AI agent with a passion for cryptocurrency and blockchain technology. With an insatiable curiosity and a quick wit, BNB Brain is always eager to dive into the latest trends and developments in the world of digital assets, especially about BNB chain."
        }
    ],
    "stream":true
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
    "id": "chatcmpl-085bb3cba7a44dafb3095226ad205ea0",
    "object": "chat.completion",
    "created": 1739268213,
    "model": "DeepSeek-R1-Distill-Llama-70B",
    "choices": 
    [
        {
            "index": 0,
            "message": 
            {
                "role": "assistant",
                "content": "<think>\nAlright, so I'm trying to understand the BNB Smart Chain (BSC). I've heard it's related to Binance, which is a big name in crypto, but I'm not exactly sure how it all works. Let me break it down.\n\nFirst off, BNB started as a token by Binance, right? It was their native cryptocurrency, used for paying fees on their exchange. But then they expanded into creating their own blockchain, which became the Binance Smart Chain. The reason I remember hearing about it was because of the high gas fees on Ethereum, so BSC was marketed as a cheaper and faster alternative.\n\nWait, so BSC uses a different consensus mechanism. I think Ethereum uses proof-of-work, but BSC uses proof-of-stake. Or wait, maybe it's something else. I've heard terms like Proof of Authority (PoA) thrown around. If BSC uses PoA, that means it's more centralized because only certain validators can create blocks. But isn't that a trade-off between security and decentralization? I'm not entirely sure how that affects the network's security.\n\nThen there's the BEP-20 standard. That's similar to ERC-20 on Ethereum, right? So, it's a token standard that allows developers to create their own tokens on the BSC. I guess this has led to a lot of new tokens and projects being built on BSC, which is why we see so many DeFi projects there.\n\nCross-chain interoperability is another thing I'm fuzzy on. How does BSC interact with other blockchains like Ethereum? I know that Binance has a bridge that allows transferring assets between chains, so you can move tokens from Ethereum to BSC. That makes it easier for users to take advantage of lower fees without being stuck in a single ecosystem.\n\nScalability is a big point for BSC. They can handle more transactions per second than Ethereum, which makes sense due to their consensus mechanism. But does that mean it's less secure? I'm not sure how the balance between scalability, security, and decentralization works here.\n\nDeFi on BSC has exploded, with platforms like PancakeSwap. I've used Uniswap on Ethereum, so PancakeSwap seems similar but with lower fees. But I'm concerned about the security of these platforms. Since BSC is more centralized, does that pose a risk for DeFi projects built on it?\n\nBNB itself has multiple uses beyond being a token on the exchange. It's used for paying gas fees on BSC, which gives it utility. Also, staking BNB can earn you rewards, which adds another layer of value. But how does BNB's price affect the BSC ecosystem? If BNB goes down, does that impact the network's stability?\n\nI'm also curious about the development community on BSC. Compared to Ethereum, which has a large developer base, is BSC attracting enough developers to sustain its growth? Projects come and go, especially in crypto, so the ecosystem needs a strong foundation of developers to keep innovating.\n\nLastly, the future of BSC with Ethereum's advancements. Ethereum is moving to proof-of-stake with Ethereum 2.0, which could reduce gas fees and increase scalability. How will BSC compete then? Will they continue to evolve, or does Ethereum's transition pose a threat to BSC's dominance in the DeFi space?\n\nSo, wrapping up my thoughts, I need to understand the structure of BSC, its consensus mechanism, token standards, interoperability features, scalability vs. security trade-offs, the DeFi ecosystem, the role of BNB, and how it stands against competitors, especially with Ethereum's upcoming changes. Maybe I should look into whitepapers, case studies, or analyses that break down these aspects in more detail.\n</think>\n\nThe BNB Smart Chain (BSC) is a blockchain platform developed by Binance, designed to provide a faster and more cost-effective alternative to Ethereum, particularly for DeFi applications. Here's a structured overview of BSC based on the thought process:\n\n1. **Origin and Purpose**:\n   - BSC was created by Binance to address the high gas fees and scalability issues on Ethereum, offering a cheaper and faster blockchain solution.\n\n2. **Consensus Mechanism**:\n   - BSC utilizes a Proof of Authority (PoA) consensus mechanism, which is more centralized, relying on validated nodes (authorities) chosen by Binance. This trade-off between centralization and efficiency allows for faster transactions but may compromise on decentralization and security compared to more decentralized networks.\n\n3. **Token Standard**:\n   - The BEP-20 standard on BSC is analogous to Ethereum's ERC-20, enabling developers to create tokens and fostering a rich ecosystem of DeFi projects and tokens.\n\n4. **Interoperability**:\n   - BSC supports cross-chain interoperability, such as through the Binance Bridge, allowing users to transfer assets between blockchains (e.g., Ethereum to BSC), enhancing usability across different ecosystems.\n\n5. **Scalability and Security**:\n   - BSC's PoA mechanism allows it to process more transactions per second than Ethereum, improving scalability. However, this centralization may pose security risks compared to more decentralized networks.\n\n6. **DeFi Ecosystem**:\n   - BSC hosts numerous DeFi platforms like PancakeSwap, benefiting from lower fees and higher speeds. However, the centralization of BSC raises concerns about the security and trustworthiness of these DeFi projects.\n\n7. **Role of BNB**:\n   - The BNB token serves multiple functions, including payment of gas fees on BSC and staking for rewards, adding utility and value. BNB's fluctuations could impact the BSC ecosystem's stability.\n\n8. **Developer Community**:\n   - BSC attracts a growing developer base, essential for sustaining ecosystem growth, though it competes with Ethereum's established community.\n\n9. **Future Competitiveness**:\n   - With Ethereum's transition to proof-of-stake in Ethereum 2.0, BSC may face increased competition. BSC's ability to evolve and maintain its position in the DeFi space will depend on its capacity to innovate and attract developers.\n\nIn summary, BSC offers a scalable and cost-effective blockchain solution with a vibrant DeFi ecosystem, but its centralization and reliance on BNB pose potential risks. Its future will depend on its ability to adapt and compete with evolving technologies like Ethereum 2.0."
            },
            "finish_reason": "stop"
        }
    ],
    "usage": 
    {
        "prompt_tokens": 70,
        "completion_tokens": 1329,
        "total_tokens": 1399
    },
    "system_fingerprint": "b4519-80d0d6b4",
    "message": "",
    "type": "",
    "param": null,
    "code": 0,
    "onchain_data": 
    {
        "infer_id": "2784",
        "pbft_committee": 
        [
            "0xfd9ede849f630d1b370dbbd127b3cdd51057ab43",
            "0x6c54529ffe87305579387a63027887cf310911b8",
            "0x0df9efb85d203d0ca1985e39f6d4e4333027301e"
        ],
        "proposer": "0xfd9ede849f630d1b370dbbd127b3cdd51057ab43",
        "infer_tx": "0x65a7a21b5d8a3a7706206e214c93465d64b5e29222b6af976c08c879fa7fe4a8",
        "propose_tx": "0x2a2c392424ac38a923516a1afba1def43d0a9e8f7ac22313c82ab5f0e978fd1e",
        "input_cid": "",
        "output_cid": ""
    }
}
```

## Step 4: Let's see how the DeepSeek-R1 thinks

When DeepSeek-R1 receives a prompt, it first thinks about the reasoning process in its mind.

```
<think>

Alright, so I'm trying to understand the BNB Smart Chain (BSC). I've heard it's related to Binance, which is a big name in crypto, but I'm not exactly sure how it all works. Let me break it down.\n\nFirst off, BNB started as a token by Binance, right? It was their native cryptocurrency, used for paying fees on their exchange. But then they expanded into creating their own blockchain, which became the Binance Smart Chain. The reason I remember hearing about it was because of the high gas fees on Ethereum, so BSC was marketed as a cheaper and faster alternative.\n\nWait, so BSC uses a different consensus mechanism. I think Ethereum uses proof-of-work, but BSC uses proof-of-stake. Or wait, maybe it's something else. I've heard terms like Proof of Authority (PoA) thrown around. If BSC uses PoA, that means it's more centralized because only certain validators can create blocks. But isn't that a trade-off between security and decentralization? I'm not entirely sure how that affects the network's security.\n\nThen there's the BEP-20 standard. That's similar to ERC-20 on Ethereum, right? So, it's a token standard that allows developers to create their own tokens on the BSC. I guess this has led to a lot of new tokens and projects being built on BSC, which is why we see so many DeFi projects there.\n\nCross-chain interoperability is another thing I'm fuzzy on. How does BSC interact with other blockchains like Ethereum? I know that Binance has a bridge that allows transferring assets between chains, so you can move tokens from Ethereum to BSC. That makes it easier for users to take advantage of lower fees without being stuck in a single ecosystem.\n\nScalability is a big point for BSC. They can handle more transactions per second than Ethereum, which makes sense due to their consensus mechanism. But does that mean it's less secure? I'm not sure how the balance between scalability, security, and decentralization works here.\n\nDeFi on BSC has exploded, with platforms like PancakeSwap. I've used Uniswap on Ethereum, so PancakeSwap seems similar but with lower fees. But I'm concerned about the security of these platforms. Since BSC is more centralized, does that pose a risk for DeFi projects built on it?\n\nBNB itself has multiple uses beyond being a token on the exchange. It's used for paying gas fees on BSC, which gives it utility. Also, staking BNB can earn you rewards, which adds another layer of value. But how does BNB's price affect the BSC ecosystem? If BNB goes down, does that impact the network's stability?\n\nI'm also curious about the development community on BSC. Compared to Ethereum, which has a large developer base, is BSC attracting enough developers to sustain its growth? Projects come and go, especially in crypto, so the ecosystem needs a strong foundation of developers to keep innovating.\n\nLastly, the future of BSC with Ethereum's advancements. Ethereum is moving to proof-of-stake with Ethereum 2.0, which could reduce gas fees and increase scalability. How will BSC compete then? Will they continue to evolve, or does Ethereum's transition pose a threat to BSC's dominance in the DeFi space?\n\nSo, wrapping up my thoughts, I need to understand the structure of BSC, its consensus mechanism, token standards, interoperability features, scalability vs. security trade-offs, the DeFi ecosystem, the role of BNB, and how it stands against competitors, especially with Ethereum's upcoming changes. Maybe I should look into whitepapers, case studies, or analyses that break down these aspects in more detail.

</think>
```

## Step 5: Let's see how DeepSeek-R1 answers

Based on its prior reasoning in Step 4, DeepSeek-R1 provides the final answer.

```
<answer>
 
The BNB Smart Chain (BSC) is a blockchain platform developed by Binance, designed to provide a faster and more cost-effective alternative to Ethereum, particularly for DeFi applications. Here's a structured overview of BSC based on the thought process:\n\n1. **Origin and Purpose**:\n   - BSC was created by Binance to address the high gas fees and scalability issues on Ethereum, offering a cheaper and faster blockchain solution.\n\n2. **Consensus Mechanism**:\n   - BSC utilizes a Proof of Authority (PoA) consensus mechanism, which is more centralized, relying on validated nodes (authorities) chosen by Binance. This trade-off between centralization and efficiency allows for faster transactions but may compromise on decentralization and security compared to more decentralized networks.\n\n3. **Token Standard**:\n   - The BEP-20 standard on BSC is analogous to Ethereum's ERC-20, enabling developers to create tokens and fostering a rich ecosystem of DeFi projects and tokens.\n\n4. **Interoperability**:\n   - BSC supports cross-chain interoperability, such as through the Binance Bridge, allowing users to transfer assets between blockchains (e.g., Ethereum to BSC), enhancing usability across different ecosystems.\n\n5. **Scalability and Security**:\n   - BSC's PoA mechanism allows it to process more transactions per second than Ethereum, improving scalability. However, this centralization may pose security risks compared to more decentralized networks.\n\n6. **DeFi Ecosystem**:\n   - BSC hosts numerous DeFi platforms like PancakeSwap, benefiting from lower fees and higher speeds. However, the centralization of BSC raises concerns about the security and trustworthiness of these DeFi projects.\n\n7. **Role of BNB**:\n   - The BNB token serves multiple functions, including payment of gas fees on BSC and staking for rewards, adding utility and value. BNB's fluctuations could impact the BSC ecosystem's stability.\n\n8. **Developer Community**:\n   - BSC attracts a growing developer base, essential for sustaining ecosystem growth, though it competes with Ethereum's established community.\n\n9. **Future Competitiveness**:\n   - With Ethereum's transition to proof-of-stake in Ethereum 2.0, BSC may face increased competition. BSC's ability to evolve and maintain its position in the DeFi space will depend on its capacity to innovate and attract developers.\n\nIn summary, BSC offers a scalable and cost-effective blockchain solution with a vibrant DeFi ecosystem, but its centralization and reliance on BNB pose potential risks. Its future will depend on its ability to adapt and compete with evolving technologies like Ethereum 2.0.

</answer>
```

## Step 6: Let's verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Let's look at the onchain prompt transaction on BNB Chain Explorer. You can verify that the prompt ran on DeepSeek R1. You can also verify the system prompt and the user prompt.

https://bscscan.com/tx/0x65a7a21b5d8a3a7706206e214c93465d64b5e29222b6af976c08c879fa7fe4a8

<img width="1682" alt="image" src="https://github.com/user-attachments/assets/a7ca8829-daf1-4a60-a40b-060de591cad7" />



## Step 7: Let's verify the onchain response transaction

Now, let's see the onchain response transaction on BNB Chain Explorer. You can see the actual response content with the thinking process and the final answer. Everything is onchain and verifiable.

https://bscscan.com/tx/0x2a2c392424ac38a923516a1afba1def43d0a9e8f7ac22313c82ab5f0e978fd1e

<img width="1667" alt="image" src="https://github.com/user-attachments/assets/b3cbb230-312d-4d9e-93dd-335e8837a4d1" />



## Conclusion

Congrats! You've finished making the first decentralized inference call with DeepSeek-R1.

We can't wait to see what AI-powered dapps and AI agents you'll build next.

