---
description: Demonstration on how to interactive chat with your dagent.
---

# Chain of thought

### 1. Install Miniconda:

\- Download and install Miniconda by following the instructions here: [Miniconda Installation.](https://docs.anaconda.com/miniconda/install/#quick-command-line-install)

### 2. Clone the repository:

```
git clone https://github.com/eternalai-org/Eternals 
cd Eternals
git checkout add/react-qa
```

### 3. Activate Conda Environment:

\- Create and activate a Conda environment for Python 3.10.0:



```
conda create -n eternalai_dagents_chat python=3.10.0
conda activate eternalai_dagents_chat
```

### 4. Install Dependencies:

\- Install the required Python packages using `pip`.&#x20;



```
pip install -r requirements.txt
```

&#x20;\- Compile and build `eternal_dagents` package

```
pip install --force-reinstall .
```

### 5. Create a `.env` File:

\- In the root of your repository, create a `.env` file to store your environment variables.

```
cp .env.example .env
```

### 6. Obtain Inference API Key:

\- Go to [EternalAI and connect your account](https://docs.eternalai.org/eternal-ai/decentralized-inference-api/api-key) to retrieve your inference API key.

### 7. Configure the `.env` File:

\- Open your `.env` file and insert your Twitter API key and inference API key in the following format:



```
# Environment
IS_SANDBOX=0

# for contract based llm
ETERNAL_BACKEND_API=https://api.eternalai.org
ETERNAL_BACKEND_API_APIKEY=your_inference_api_key

```

### 8. Run the Application:

\- Execute the script to start the application:



```
python3 toolkit/chat-lite.py 
```

&#x20;\- Sample output:

<div data-full-width="false"><figure><img src="../../.gitbook/assets/image (64).png" alt=""><figcaption></figcaption></figure></div>
