# Decentralized Inference API

This developer guides show you how to use Decentralized Inference API with DeepSeek-R1 on Base.

* Decentralized
* Onchain verifiable
* Permissionless

## Step 1: Get your Decentralized Inference API key

1. Go to http://eternalai.org/api
2. Login with Twitter
3. Get the Decentralized Inference API key for Base
 
It's free.

## Step 2: Make your first Onchain DeepSeek R1 API call

1. Enter your API key
2. Set Chain ID as 8453 (Base)
3. Set the model as "DeepSeek V3"

Then run this curl command:

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
                        "content": "You are a helpful assistant."
                },
                {
                        "role": "user",
                        "content": "Hello!"
                }
         ]
 }'
```

## Step 3: Understand the response

You'll find the following info in the response.

1. The pBFT committee handling the inference request
2. The response proposer and verifiers
3. The onchain prompt transaction
4. The onchain response transaction
5. The response


{
  "id": "chatcmpl-p5n76HCHiWOYFMYj4S4LNRv9g0Pz91z3",
    "object": "chat.completion",
      "created": 1737471058,
        "model": "DeepSeek-R1-Distill-Llama-70B",
        "choices": [
        {
              "index": 0,
              "message": {
                      "role": "assistant",
                              "content": "<think>\nOkay, the user greeted me with \"Hello!\" and I responded warmly. Now, they're asking me to act as a helpful assistant. I should keep the conversation going by inviting them to ask questions or share what they need help with. I'll make sure to sound friendly and approachable.\n</think>\n\nHello! How can I assist you today? If you have any questions or need help with something, feel free to ask!"
                                    
              },
                    "finish_reason": "stop"
                        
        }
          
        ],
        "usage": {
            "prompt_tokens": 11,
                "completion_tokens": 89,
                    "total_tokens": 100
                      
        },
          "system_fingerprint": "b4519-80d0d6b4",
            "message": "",
              "type": "",
                "param": null,
                  "code": 0,
                  "onchain_data": {
                      "infer_id": "13209",
                      "pbft_committee": [
                            "0xc557e228d2443b437e5297a8266b0f3257b2109a",
                                  "0xb9deb8db9724e2cb551164cdab0dd03a484cb9da",
                                        "0xfe5598e12532c8fe8fa586f8f6d2e12e82bb2825"
                                            
                      ],
                          "proposer": "0xfe5598e12532c8fe8fa586f8f6d2e12e82bb2825",
                              "infer_tx": "0xd3b7adde419d44083e5d66489191053e15b250aea5dc168b8f6f4e18c90b2ff2",
                                  "propose_tx": "0x57d40d2720b512fbb64541e70611d044c8703c6fc05b7122a561f95645bc0f62"
                                    
                  }
                  
}

## Step 4: Let's verify the onchain prompt tx

With Eternal AI's Decentralized Inference, everything is onchain verifiable.

Let's look at the onchain prompt tx on Base Explorer. You can verify that the prompt ran on DeepSeek v3 & see the prompt content.

## Step 5: Let's verify the onchain response tx

 Now, let's see the onchain response tx on Base Explorer. You can see the actual response content. Everything is onchain and verifiable.

  Question for you: How well did Onchain DeepSeek do on this question?


