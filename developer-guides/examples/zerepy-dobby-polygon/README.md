# Build unstoppable Zerepy agents with Dobby on Polygon

This developer guide teaches you how to build your first decentralized ZerePy agent:

- Decentralized: Use Eternal AI's Decentralized Inference API instead of OpenAI's Centralized API.
- Onchain-verifiable: Don't trust AI, verify them. All inferences are recorded onchain and verifiable by anyone.
- Unstoppable: Agents run exactly as coded without any possibility of downtime, censorship, fraud, or third-party interference.
- Intelligent: Give your ZerePy agents new superpowers using Dobby, the AI model developed by Sentinent that can say what others cannot.

## Background:

For this guide, we'll use Eternal AI's Decentralized Inference, which has the same interface as OpenAI (so no learning curve for you) — but is decentralized.

The AI model used is onchain Dobby, deployed on Polygon.

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

Check the onchain prompt transaction on PolygonScan. You can verify that this prompt runs on Dobby and see its content stored on Filecoin.

https://polygonscan.com/tx/0xb83fcfbf0ec5fe1d0e3a65e387ee0632dd066803f9cfa0543ca760826d307e89

https://gateway.lighthouse.storage/ipfs/bafkreiex4qoqwfrkeoeo74zxcatsvavpv45roceilqscht6qv6l2tov4oq

![carbon (4)](https://github.com/user-attachments/assets/f5019777-3753-470c-858c-784e825d7f2f)

## Step 11: Verify the onchain response transaction

Now, let’s check the onchain response transaction on PolygonScan.

You can see the actual response content. Everything is onchain and verifiable.

https://polygonscan.com/tx/0xf346be037128ca1cb870388986dfb9f5bbffc74ffa43dc786764fb2a3a91c5e1

https://gateway.lighthouse.storage/ipfs/bafkreifzmymuufsqdemfz5ppdt34psizwhownjfljm4fa4a2zfrwwkulru

![carbon (5)](https://github.com/user-attachments/assets/1315e3cd-0584-4d68-8a4a-664b14c2aa87)

# Conclusion

Congrats!

You’ve finished building your Zerepy agent with two superpowers:

- Onchain-verifiable on Polygon
- Powered by Dobby
Enjoy building!


