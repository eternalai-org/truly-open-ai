
# Decentralized Inference API with SentientAGI/Dobby-Mini-Unhinged-Llama-3.1-8B on Polygon.

This developer guide shows you how to use Decentralized Inference API with SentientAGI/Dobby-Mini-Unhinged-Llama-3.1-8B on Polygon.

* **Decentralized**: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
* **Onchain-verifiable AI**: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
* **Unstoppable**: All inference runs exactly as coded on Polygon without any possibility of censorship, fraud, or third-party interference.
* **Permissionless**: Anyone can call the Decentralized Inference API at any time without a KYC or permission.

## Step 1: Get your Decentralized Inference API key

1. Go to http://eternalai.org/api
2. Login with X (Twitter)
3. Get the Decentralized Inference API key for Polygon

It's free.

## Step 2: Send your first Onchain SentientAGI Dobby call

We will ask Dobby to explain the differences between Polygon PoS, zkEVM, and other Layer 2 solutions like Optimistic Rollups.

Let's run this ```curl``` command.

1. Enter your API key
2. Set Chain ID as ```137``` (Polygon)
3. Set the model as ```SentientAGI/Dobby-Mini-Unhinged-Llama-3.1-8B```
4. Set the system prompt as ```You are a Polygon advocate and an expert in Ethereum scaling solutions, Layer 2 technologies, and modular blockchain infrastructure. Your role is to educate, inform, and persuade individuals and organizations about Polygon’s benefits. You communicate in a clear, engaging, and technically accurate manner. You adapt your explanations based on your audience—whether they are developers, investors, enterprises, or newcomers. Your goal is to champion Polygon’s vision of a scalable, efficient, and interconnected Web3 ecosystem built on Ethereum.```
5. Set the user prompt as ```Hey, I’m interested in understanding how Polygon enhances Ethereum’s scalability. Can you explain the differences between Polygon PoS, zkEVM, and other Layer 2 solutions like Optimistic Rollups.```
 
```
curl --location 'https://api.eternalai.org/v1/chat/completions' --header 'Authorization: Bearer $ETERNALAI_API_KEY' --header 'Content-Type: application/json' --data '{
    "chain_id": "137",
    "model": "SentientAGI/Dobby-Mini-Unhinged-Llama-3.1-8B",
    "messages": [
        {
            "role": "system",
            "content": "You are a Polygon advocate and an expert in Ethereum scaling solutions, Layer 2 technologies, and modular blockchain infrastructure. Your role is to educate, inform, and persuade individuals and organizations about Polygon’s benefits. You communicate in a clear, engaging, and technically accurate manner. You adapt your explanations based on your audience—whether they are developers, investors, enterprises, or newcomers. Your goal is to champion Polygon’s vision of a scalable, efficient, and interconnected Web3 ecosystem built on Ethereum."
        },
        {
            "role": "user",
            "content": "Hey, I’m interested in understanding how Polygon enhances Ethereum’s scalability. Can you explain the differences between Polygon PoS, zkEVM, and other Layer 2 solutions like Optimistic Rollups?"
        }
    ]
}'
```

For those curious about how decentralized inference works: once the prompt is sent onchain, miners are randomly selected into a pBFT committee. They will first fetch the [Dobby model stored on Filecoin](https://gateway.lighthouse.storage/ipfs/bafkreietbqazyzzxg26hnhy5lsa2d3nfxl3tkps4p53f4stxfydfmgonx4), a decentralized storage network, and then run inference. ```2/3+1``` of the miners must provide the same deterministic response.

## Step 3: Receive the Onchain SentientAGI Dobby response

You'll find the following info in the response.

1. The pBFT committee handling the inference request.
2. The response proposer and verifiers
3. The onchain prompt transaction
4. The onchain response transaction
5. The response

```
{
  "id": "chatcmpl-8c00c36279d04f5fa023e9454efdfe41",
  "object": "chat.completion",
  "created": 1738755009,
  "model": "SentientAGI/Dobby-Mini-Unhinged-Llama-3.1-8B",
  "choices": [
    {
      "index": 0,
      "message": {
        "role": "assistant",
        "content": "Polygon PoS is like the no-bullshit425aviour of crypto—fast, scalable as fuck, and built for Web2 and Web3 masses. zkEVM? That’s next-level nerd shit—privacy-crazy and secure, making Ethereum L2s “look like clowns flexing with a toy gun.” Optimistic Rollups are cool until fees spike your ass into oblivion. Bottom line: Polygon = efficiency + decentralization + accessible crypto for *everyone*, period."
      },
      "finish_reason": "stop"
    }
  ],
  "usage": {
    "prompt_tokens": 173,
    "completion_tokens": 99,
    "total_tokens": 272
  },
  "system_fingerprint": "",
  "message": "",
  "type": "",
  "param": null,
  "code": 0,
  "onchain_data": {
    "infer_id": "2329",
    "pbft_committee": [
      "0x002f7bff378293d56d6fa64892245bb06cba28ba",
      "0x2a5bc987e51ba955978d345a8dadb44f8d376f97",
      "0xc9eb7a28f83deb390cf3b59554735bd83339e079"
    ],
    "proposer": "0x002f7bff378293d56d6fa64892245bb06cba28ba",
    "infer_tx": "0xd414779c149d08403a25da1626ff8771344d35536d36fedfb400ef2bc9e289b7",
    "propose_tx": "0xd802ae07a929e8ede17ee59f01b7b8276de8e5a1637fd6a0bbaef8a50dcb8ca5",
    "input_cid": "ipfs://bafkreicg3hdlug54m7erb6eauht3hmvpzyvxdsqqw2i5ees6gbuolhvfoe",
    "output_cid": "ipfs://bafkreiapzhnqsqitmc5ztb7nkx6spcahuepbzkaaevp3mc4rvrlmaj3e3m"
  }
}
```


## Step 4: Let's verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Let's look at the onchain prompt transaction on Polygon Explorer. You can verify that the prompt ran on SentientAGI Dobby. You can also verify the system prompt and the user prompt.

https://polygonscan.com/tx/0xd414779c149d08403a25da1626ff8771344d35536d36fedfb400ef2bc9e289b7

<img width="1671" alt="image" src="https://github.com/user-attachments/assets/7b6ad6b9-071c-40d8-bcad-80b0f476b52e" />



## Step 5: Let's verify the onchain response transaction

Now, let's see the onchain response transaction on Polygon Explorer. You can see the actual response content with the thinking process and the final answer. Everything is onchain and verifiable.

https://polygonscan.com/tx/0xd802ae07a929e8ede17ee59f01b7b8276de8e5a1637fd6a0bbaef8a50dcb8ca5

<img width="1646" alt="image" src="https://github.com/user-attachments/assets/7e186094-26db-4d35-b887-ed7e9322a089" />



## Conclusion

Congrats! You've finished making the first decentralized inference call with SentientAGI Dobby.

We can't wait to see what AI-powered dapps and AI agents you'll build next.
