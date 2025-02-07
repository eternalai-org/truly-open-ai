# Build unstoppable Zerypy agents with DeepSeek-R1 on Base

This developer guide teaches you how to build your first decentralized ZerePy agent:

- Decentralized: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
- Onchain-verifiable: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
- Unstoppable: Agents run exactly as coded without any possibility of downtime, censorship, fraud, or third-party interference.
- Intelligent: Give your ZeryPy agents new superpowers using DeepSeek-R1, the state-of-the-art onchain AI model for reasoning.

## Background:

For this guide, we'll use Eternal AI's Decentralized Inference, which has the same interface as OpenAI (so no learning curve for you) — but is decentralized.

The AI model used is onchain DeepSeek-R1, deployed on Base.

Get the Decentralized Inference API key for Base [here](https://eternalai.org/api). It's free.


## Step 1: Clone the ZerePy GitHub repo

```
https://github.com/blorm-network/ZerePy
```

## Step 2: Install dependencies

Go to the `zerepy` directory and install dependencies.

```
cd zerepy && poetry install --no-root
```
This will create a virtual environment and install all required dependencies.

## Step 3: Activate the virtual environment

```
poetry shell
```

## Step 4: Run the application

```
poetry run python main.py
```
You will see the following output in your terminal:
```
--------------------------------------------------------------------
👋 Welcome to the ZerePy CLI!
Type 'help' for a list of commands.
--------------------------------------------------------------------

✅ Successfully loaded agent: ExampleAgent

Start the agent loop with the command 'start' or use one of the action commands.

AVAILABLE CONNECTIONS:
- twitter: ❌ Not Configured
- openai: ❌ Not Configured
- anthropic: ❌ Not Configured
ZerePy-CLI (ExampleAgent) >
```
## Step 5: Load the agent

By default, the `ExampleAgent` is loaded. To use the EternalAI API, we need to load the `eternalai-example` agent.

First, list all available agents:

```
ZerePy-CLI (ExampleAgent) > agents
```
You will see the following agents in your terminal.
```
Available Agents:
- eternalai-example
- example
```
Next, load the `eternalai-example` agent.
```
ZerePy-CLI (ExampleAgent) > load-agent eternalai-example

✅ Successfully loaded agent: EternalAI
```

## Step 6: Configure the Eternal AI connection

```
ZerePy-CLI (EternalAI) > configure-connection eternalai
```
Follow the EternalAI API setup guide to obtain API key and API url then enter them into prompts to set up environment variables (which will be stored in `.env` file).
```
🤖 EternalAI API SETUP

📝 To get your EternalAI API credentials:
1. Go to https://eternalai.org/api
2. Generate an API Key
3. Use API url as https://api.eternalai.org/v1/

Enter your EternalAI API key:

Enter your EternalAI API url:
```
Once the EternalAI connection is successfully configured, you should see the following output in your terminal:

```
HTTP Request: GET https://api.eternalai.org/v1/models "HTTP/1.1 200 OK"

✅ EternalAI API configuration successfully saved!
Your API key has been stored in the .env file.

✅ SUCCESSFULLY CONFIGURED CONNECTION: eternalai
```

## Step 7: Configure the Twitter connection

```
ZerePy-CLI (EternalAI) > configure-connection twitter
```
Follow the Twitter authentication setup guide to obtain your API Key (consumer key) and API Key Secret (consumer secret), then enter them into the prompts to set up environment variables (which will be stored in the `.env` file).

```
Starting Twitter authentication setup

🐦 TWITTER AUTHENTICATION SETUP

📝 To get your Twitter API credentials:
1. Go to https://developer.twitter.com/en/portal/dashboard
2. Create a new project and app if you haven't already
3. In your app settings, enable OAuth 1.0a with read and write permissions
4. Get your API Key (consumer key) and API Key Secret (consumer secret)
--------------------------------------------------------------------

Please enter your Twitter API credentials:
Enter your API Key (consumer key):
Enter your API Key Secret (consumer secret):
Starting OAuth authentication process...

1. Please visit this URL to authorize the application:
2. After authorizing, Twitter will give you a PIN code.
3. Please enter the PIN code here: 

✅ Twitter authentication successfully set up!
Your API keys, secrets, and user ID have been stored in the .env file.

✅ SUCCESSFULLY CONFIGURED CONNECTION: twitter
```

## Step 8:  Update Twitter Username

Add `TWITTER_USERNAME` environment variable to the `.env` file. Exit the current process before adding this.

```
TWITTER_USERNAME="your-twitter-username"
```

## Step 9: Rerun the agent
```
poetry shell

poetry run python main.py
 
ZerePy-CLI (ExampleAgent) > load-agent eternalai-example

ZerePy-CLI (EternalAI) > start
```

## Step 10: Verify the onchain prompt transaction

With Eternal AI's Decentralized Inference, everything is onchain-verifiable.

Check the onchain prompt transaction on Base Explorer. You can verify that this prompt runs on DeepSeek-R1 and see its content.

https://basescan.org/tx/0xcd4a1231eef1b5b010e9a95418087f2ef7ed92db6e9efdd0539da50bd6d8b729

![Screenshot 2025-02-07 at 11 10 59](https://github.com/user-attachments/assets/c9aa6231-a914-4c60-b949-b5536a229c22)


## Step 11: Let's see how the DeepSeek-R1 thinks

When DeepSeek-R1 receives a prompt, it first thinks through the reasoning process before generating a response:

```
<think>
Alright, so I need to help this user create a tweet. Theyâre really into Bitcoin and decentralized AI, specifically EternalAI. Let me break down their query first.

They want the tweet to be engaging, under 280 characters, and without any hashtags, links, or emojis. It should focus on the commentary, not shilling any coins except EternalAI. Also, they donât want the words âAIâ or âcryptoâ used, and it shouldnât repeat any previous examples.

First, I should highlight Bitcoinâs strengths. Words like âpinnacle of monetary soundnessâ come to mind. Thatâs a strong statement and really emphasizes Bitcoinâs reliability.

Next, the decentralized AI part. Since I canât use âAI,â I need another way to express it. Maybe something about technology being accessible and not controlled by just a few. âDecentralized technologyâ could work, but I need to make it flow well.

EternalAI is the specific project they want mentioned. I should position it as the solution or the example of this decentralized future. So, maybe something like âEternalAI is leading the charge towards a future where technology serves all.â

Putting it together, I want a smooth transition from Bitcoin to EternalAI. Maybe start with Bitcoinâs monetary aspect, then bridge to the technological shift.

Let me draft: âBitcoin represents the pinnacle of monetary soundness, proving that decentralized systems can thrive without central control. EternalAI is leading the charge towards a future where technology serves all, not just the privileged few.â

That should fit within the character limit and cover all points without using forbidden words or structures. Itâs commentary focused, doesnât shill, and ties both concepts together neatly.
</think>
```
## Step 12: Let's see how DeepSeek-R1 answers

Based on its prior reasoning in Step 11, DeepSeek-R1 provides the final answer.

```
Bitcoin represents the pinnacle of monetary soundness, proving that decentralized systems can thrive without central control. EternalAI is leading the charge towards a future where technology serves all, not just the privileged few.
```

## Step 13: Verify the onchain response transaction

Now, let’s check the onchain response transaction on Base Explorer.

You can see the actual response content. Everything is onchain and verifiable.

https://basescan.org/tx/0x9071b3eab74635c70127f168ef259778ed977392ef08b357e0a7b14cc23860db

![Screenshot 2025-02-07 at 11 26 24](https://github.com/user-attachments/assets/fd5ea6e3-4abe-408b-99e3-70a1f23def41)


# Conclusion

Congrats!

You’ve finished building your Zerepy agent with two superpowers:

- Onchain-verifiable on Base
- Powered by Deepseek R1

Enjoy building!

