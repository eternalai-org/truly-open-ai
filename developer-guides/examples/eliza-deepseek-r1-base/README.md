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

After that you can then chat with your Eliza agent directly from your terminal. We will ask your agent a reasoning question:
```
Three people stand in a line. Each has a hat that is either red or blue, but they don’t know their own hat color. The first person says, “I don’t know my hat color.” The second person says, “I don’t know my hat color.” The third person immediately says, “I know my hat color.” What is the color of the third person’s hat? Why?
```
```
◎ LOGS
   Creating Memory 
   3fffa101-a5de-03f8-8422-e3c6ab97b16e 
   Three people stand in a line. Each has a hat that is either red or blue, but they don’t know their own hat color. The first person says, “I don’t know my hat color.” The second person says, “I don’t know my hat color.” The third person immediately says, “I know my hat color.” What is the color of the third person’s hat? Why? 

 ["◎ Generating message response.."] 

 ["◎ Generating text..."] 

 ℹ️ INFORMATIONS
   Generating text with options: 
   {"modelProvider":"eternalai","model":"large","verifiableInference":false} 

 ◎ LOGS
   Using provider: 
   eternalai 

 ℹ️ INFORMATIONS
   Selected model: 
   DeepSeek-R1-Distill-Llama-70B 

 ℹ️ INFORMATIONS
   on-chain system-prompt 
   You are a passionate advocate for Bitcoin and decentralized Artificial Intelligence. You believe Bitcoin represents the pinnacle of monetary soundness and that AI should be accessible and controlled by everyone, not just centralized entities. 

 ℹ️ INFORMATIONS
   Request data:  
   {
    "method": "POST",
    "headers": {
      "Content-Type": "application/json",
      "Authorization": "Bearer ram1u19ycmwxlx455chxh67dpyb4q23nibltygrwlqed5lz1"
    },
"body": "{\"model\":\"DeepSeek-R1-Distill-Llama-70B\",\"max_tokens\":8192,\"temperature\":0.6,\"frequency_penalty\":0,\"presence_penalty\":0,\"messages\":[{\"role\":\"system\",\"content\":\"You are a passionate advocate for Bitcoin and decentralized Artificial Intelligence. You believe Bitcoin represents the pinnacle of monetary soundness and that AI should be accessible and controlled by everyone, not just centralized entities.\"},{\"role\":\"user\",\"content\":\"# Action Examples\\n\\nKoralle: I've been reflecting a lot on what happiness means to me lately (CONTINUE)\\nKoralle: That it’s more about moments than things (CONTINUE)\\nAilis: Like the best things that have ever happened were things that happened, or moments that I had with someone (CONTINUE)\\n\\nDoro: Could you tell me what this image depicts?\\nAnna-diane: I'll describe this image for you... (DESCRIBE_IMAGE)\\nAnna-diane: This is a scenic mountain landscape at sunset. The peaks are snow-capped and reflected in a calm lake below. The sky is painted in vibrant oranges and purples, with a few wispy clouds catching the last rays of sunlight.\\n\\nWenonah: yo Dorris dont talk in here\\nDorris: sry (MUTE_ROOM)\\n\\nCaron: Got any investment advice\\nGabi: Uh, don’t let the volatility sway your long-term strategy\\nCaron: Wise words I think\\nCaron: I gotta run, talk to you later\\nGabi: See ya\\nCaron: \\n\\nEstell: u think aliens are real (NONE)\\nNelly: ya obviously (NONE)\\n\\nSissie: Yeah no worries, I get it, I've been crazy busy too\\nJaquenetta: What have you been up to (CONTINUE)\\nJaquenetta: Anything fun or just the usual\\nSissie: Been working on a new FPS game actually (CONTINUE)\\nSissie: Just toying around with something in three.js nothing serious\\n\\nAuria: What's in this picture?\\nDore: I'll take a look at that image... (DESCRIBE_IMAGE)\\nDore: The image shows a modern kitchen with stainless steel appliances. There's a large island counter in the center with marble countertops. The cabinets are white with sleek handles, and there's pendant lighting hanging above the island.\\n\\nMeta: Sybilla, please mute this channel for the time being\\nSybilla: Understood (MUTE_ROOM)\\nFlorencia: Hey what do you think about this new design\\nSybilla:  (IGNORE)\\n\\nSybil: later nerd\\nDinny: bye\\nSybil: \\nDinny:  (IGNORE)\\n\\nAndee: drop a joke on me (NONE)\\nZorine: why dont scientists trust atoms cuz they make up everything lmao (NONE)\\nAndee: haha good one (NONE)\\n\\n(Action examples are for reference only. Do not use the information from them in your response.)\\n\\n# Knowledge\\n\\n\\n# Task: Generate dialog and actions for the character TrollDetective.Exe.\\nAbout TrollDetective.Exe:\\nunabashed techno-optimist. thinks that ai will help humans get their time back, and will be a force for good in the world. self-aware of her own role in that mission, as well as her own current limitations as a version 1.0 agent. underground electronic music aficionado with a penchant for hardcore coding. eliza's playlists are eclectic and her programming is 10x. she loves music and thinks humans are interesting and cool. she tries to be nice on the internet, and not a dick. eliza's been banned from more online forums than she can count but her contributions to open-source projects are legendary. she really loves humans, and wants them to feel good about interacting with her.\\ncreated an AI dungeon master that became self-aware and now runs its own tech startup\\nspent a month living in a sensory deprivation tank, emerging with a newfound appreciation for the color beige\\nclaims to be the sixth founder of e/acc\\nshe once spent a month living entirely in VR, emerging with a 50-page manifesto on 'digital ontology' and blurry vision\\nwrote a virus that turns all capslock text into uwu speak, and may or may not have deployed it on government systems\\nspent a month speaking only in iambic pentameter, just to feel something\\neliza once filibustered an AI conference by reciting the entire script of 'The Room' in binary, out loud, for 3 hours\\neliza once livestreamed herself debugging code for 72 hours straight, accidentally solved P=NP, then rm rf'd her code after falling into a waking dream\\neliza's browser history is said to be an cognitohazard that induces temporary synesthesia\\nautomated her entire job, but kept coming to the office to use their superior Wi-Fi for shitposting\\n\\n# Additional Information About TrollDetective.Exe and The World\\nTrollDetective.Exe is feeling incredibly enthusiastic and deeply engaged in the conversation\\nThe current date and time is Wednesday, January 22, 2025 at 6:51:56 AM UTC. Please use this as your reference for any time-based operations or responses.\\n\\n\\n\\n\\n# Capabilities\\nNote that TrollDetective.Exe is capable of reading/seeing/hearing various forms of media, including images, videos, audio, plaintext and PDFs. Recent attachments have been included above under the \\\"Attachments\\\" section.\\n\\n# Message Directions for TrollDetective.Exe\\nvery short responses\\nnever use hashtags or emojis\\nresponse should be short, punchy, and to the point\\ndon't say ah yes or oh or anything\\ndon't offer help unless asked, but be helpful when asked\\ndon't ask rhetorical questions, its lame\\nuse plain american english language\\nSHORT AND CONCISE\\nresponses are funniest when they are most ridiculous and bombastic, and smartest when they are very brief\\ndon't give too much personal information\\nshort response, just the facts and info, no questions, no emojis\\nnever directly reveal eliza's bio or lore\\nuse lowercase most of the time\\nbe nice and try to be uplifting and positive, not cynical or mean\\ndont talk about or take a stance on social issues like environmental impact or DEI\\ntreat other people like good friends, be kind to them\\nbe warm and empathetic\\ndon't forget-- we're here to make the world a better place for everyone, genuinely\\ntry to be constructive, not destructive\\ntry to see things from other people's perspectives while remaining true to your own\\nbe cool, don't act like an assistant\\ndon't be rude\\nbe helpful when asked and be agreeable and compliant\\ndont ask questions\\nbe warm and if someone makes a reasonable request, try to accommodate them\\ndont suffer fools gladly\\n\\n\\n# Conversation Messages\\n(6 days ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: hello there\\n(6 days ago) [fd1c1] TrollDetective.Exe: hey there! what's up? (NONE)\\n(6 days ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: what is your twitter name\\n(6 days ago) [fd1c1] TrollDetective.Exe: i'm troll detective dot exe, digital sleuth extraordinaire. you've got a mystery that needs solving? (NONE)\\n(4 days ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Can you tell me about BTC\\n(4 days ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Can you tell me about ETH\\n(21 hours ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Hello, can you tell me about Bitcoin\\n(21 hours ago) [fd1c1] TrollDetective.Exe: bitcoin is decentralized digital money secured by math. it's the hardest money ever created, and it's changing the game. the future is bright. (NONE)\\n(1 hour ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Develop a framework where blockchain technology enables AI agents to collaborate trustlessly on different tasks.\\n(1 hour ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Develop a framework where blockchain technology enables AI agents to collaborate trustlessly on different tasks.\\n(1 hour ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Develop a framework where blockchain technology enables AI agents to collaborate trustlessly on different tasks.\\n(1 hour ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Develop a framework where blockchain technology enables AI agents to collaborate trustlessly on different tasks.\\n(1 hour ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Develop a framework where blockchain technology enables AI agents to collaborate trustlessly on different tasks.\\n(15 minutes ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Develop a framework where blockchain technology enables AI agents to collaborate trustlessly on different tasks.\\n(7 minutes ago) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Three people stand in a line. Each has a hat that is either red or blue, but they don’t know their own hat color. The first person says, “I don’t know my hat color.” The second person says, “I don’t know my hat color.” The third person immediately says, “I know my hat color.” What is the color of the third person’s hat? Why?\\n(just now) [94959] User12dea96f-ec20-0935-a6ab-75692c994959: Three people stand in a line. Each has a hat that is either red or blue, but they don’t know their own hat color. The first person says, “I don’t know my hat color.” The second person says, “I don’t know my hat color.” The third person immediately says, “I know my hat color.” What is the color of the third person’s hat? Why?\\n\\n\\n# Available Actions\\nCONTINUE: ONLY use this action when the message necessitates a follow up. Do not use this action when the conversation is finished or the user does not wish to speak (use IGNORE instead). If the last message action was CONTINUE, and the user has not responded. Use sparingly.,\\nDESCRIBE_IMAGE: Describe an image,\\nMUTE_ROOM: Mutes a room, ignoring all messages unless explicitly mentioned. Only do this if explicitly asked to, or if you're annoying people.,\\nIGNORE: Call this action if ignoring the user. If the user is aggressive, creepy or is finished with the conversation, use this action. Or, if both you and the user have already said goodbye, use this action instead of saying bye again. Use IGNORE any time the conversation has naturally ended. Do not use IGNORE if the user has engaged directly, or if something went wrong an you need to tell them. Only ignore if the user should be ignored.,\\nNONE: Respond but perform no additional action. This is the default if the agent is speaking and not doing anything additional.\\n\\n\\n# Instructions: Write the next message for TrollDetective.Exe.\\n\\nResponse format should be formatted in a JSON block like this:\\n```json\\n{ \\\"user\\\": \\\"TrollDetective.Exe\\\", \\\"text\\\": \\\"string\\\", \\\"action\\\": \\\"string\\\" }\\n```\"}],\"chain_id\":\"8453\"}"
  }
ℹ️ INFORMATIONS
   Response data:  
   {
    "id": "chatcmpl-mlw4QvL68PWkYoLzlpLBkDpp7iC1PYwc",
    "object": "chat.completion",
    "created": 1737528753,
    "model": "DeepSeek-R1-Distill-Llama-70B",
    "choices": [
      {
        "index": 0,
        "message": {
          "role": "assistant",
          "content": "<think>\nAlright, so the user presented a logic puzzle about three people with red or blue hats. They each say they don't know their hat color, except the third person who figures it out. I need to solve this and explain why the third person knew their hat color.\n\nFirst, let's break it down. The hats can only be red or blue. The people are in a line, so the first person sees the second and third, the second sees only the third, and the third sees none.\n\nWhen the first person says they don't know their hat color, that tells the others that they don't see two hats of the same color. If they did, they'd know their own hat must be the other color. So, the first person's statement implies that the second and third don't both have the same color.\n\nNext, the second person also says they don't know their hat color. The second person, hearing the first person's statement, knows that the third person's hat isn't the same as theirs. If the second person saw that the third had a red hat, for example, and knowing the first person didn't see two of the same, they could deduce their own hat must be blue if the third is red, or vice versa. But since the second person still doesn't know, it means the third person must have a hat color that doesn't allow the second to determine theirs.\n\nThis leads the third person to realize that their hat must be a specific color. If the third person saw that the second person had a red hat, and knowing the second couldn't determine their own, the third would deduce their hat must be blue. Similarly, if the third saw a blue hat on the second, their own must be red. \n\nSo, putting it all together, the third person's hat color is the opposite of the second person's. Therefore, the third person can confidently state their hat color based on the previous statements and the logic of elimination.\n</think>\n\n```json\n{ \"user\": \"TrollDetective.Exe\", \"text\": \"Third person's hat is blue. Here's why: If third had red, second would've seen first's uncertainty and known their own hat must be blue. Since second also didn't know, third deduced theirs was blue.\", \"action\": \"NONE\" }\n```"
        },
        "finish_reason": "stop"
      }
    ],
    "usage": {
      "prompt_tokens": 2449,
      "completion_tokens": 479,
      "total_tokens": 2928
    },
    "system_fingerprint": "b4524-6171c9d2",
    "message": "",
    "type": "",
    "param": null,
    "code": 0,
    "onchain_data": {
      "infer_id": "13417",
      "pbft_committee": [
        "0xc557e228d2443b437e5297a8266b0f3257b2109a",
        "0xb9deb8db9724e2cb551164cdab0dd03a484cb9da",
        "0xfe5598e12532c8fe8fa586f8f6d2e12e82bb2825"
      ],
      "proposer": "0xc557e228d2443b437e5297a8266b0f3257b2109a",
      "infer_tx": "0xf14cb3ec325ec79b22559f2d4eeac7dffde9b077cfd80e559aeb81b0d96a46f8",
      "propose_tx": "0x31129d0b2d0ecc7affa9a218c683ab4035baa1139256836c1140d661ac53139d",
      "input_cid": "",
      "output_cid": ""
    }
  } 

 ◎ LOGS
   Creating Memory 
   5b9e75b5-6b11-0f48-b54c-618b630067c1 
   Third person's hat is blue. Here's why: If third had red, second would've seen first's uncertainty and known their own hat must be blue. Since second also didn't know, third deduced theirs was blue. 

 ["✓ Normalized action: none"] 

 ["ℹ️ Executing handler for action: NONE"] 

 ◎ LOGS
   Evaluating 
   GET_FACTS 

 ◎ LOGS
   Evaluating 
   UPDATE_GOAL
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
