# dAgent

dAgent is a decentralized conversational agent that can perform various tasks using a combination of toolsets and language models. The dAgent can be configured to execute multiple missions, each with its own set of toolsets and language models.

## Prerequisites
In this project, we tested with ***Python 3.10.0***, but it should work with any version of Python 3.10.x or higher.

## Installation

Set up the project environment using one of the methods below: **Conda** or **Pip**.

### Using Conda

1. **Create a Conda Environment**

    ```bash
    conda create -n dagent python=3.10.0
    ```

2. **Activate the Conda Environment:**

    ```bash
    conda activate dagent
    ```

3. **Install Dependencies:**

    Navigate to the project directory and run:

    ```bash
    pip install -r requirements.txt
    ```

### Using Pip
If you prefer to use Pip without Conda, simply navigate to the project directory and run:

```bash
pip install -r requirements.txt
```

## Setup and Configuration

To get started, follow the steps below to configure the project:
### **1. Configuration File**

A sample configuration file, [**eternal.json**](configs/eternal.json), is provided in the configs directory. You can customize this file to fit your specific requirements.

### **Configuration File Structure**

```json
{
    "characteristic": {
        "agent_personal_info": {
            "twitter_username": "realDonaldTrump",
            "agent_name": "Trump"
        },
        "bio": [
            "SAVED America from the China Virus (while they let cities burn)",
            "secured the Southern Border COMPLETELY (until they DESTROYED it)",
            ...
        ],
        "lore": [
            "Democrats using Secret Service assignments as election interference",
            "they let Minneapolis burn in 2020 (then begged for help)",
            ...
        ],
        "knowledge": [
            "knows EXACT cost to families under Kamala ($29,000)",
            "understands REAL border numbers (worse than reported)",
            ...
        ],
        "example_posts": [
            "NO TAX ON TIPS! NO TAX ON OVERTIME! NO TAX ON SOCIAL SECURITY FOR OUR GREAT SENIORS!",
            "Lyin' Kamala has allowed Illegal Migrants to FLOOD THE ARIZONA BORDER LIKE NEVER BEFORE. I WILL STOP IT ON DAY ONE! DJT",
            ...
        ],
        "topics": [
            "border security crisis",
            "Kamala's tax hikes",
            ...
        ]
    },
    "missions": [
        {
            "task": "Do something",
            "system_reminder": "",
            "toolset_cfg": [
                {
                    "name": "TwitterToolset",
                    "init_params": {
                        "exclude": ["tweet"]
                    }
                },
                {
                    "name": "TradingToolset",
                    "init_params": {
                        "exclude": ["sell"]
                    }
                }
            ],
            "llm_cfg": {
                "name": "EternalAIChatCompletion",
                "init_params": {
                    "model_name": "NousResearch/Hermes-3-Llama-3.1-70B-FP8",
                    "max_tokens": 1024,
                    "model_kwargs": {},
                    "temperature": 0.3,
                    "max_retries": 2
                }
            },
            "scheduling": {
                "interval_minutes": 3 
            }
        }
    ]
}
```

### **Configuration Fields**

### **Characteristic**

- **`agent_personal_info`**: Contains information about the dAgent, including:
    - **`twitter_username`**: The Twitter username of the dAgent.
    - **`agent_name`**: The name of the dAgent.
- **`bio`**: An array of strings representing the dAgent's bio.
- **`lore`**: An array of strings representing the dAgent's lore or background information.
- **`knowledge`**: An array of strings representing the dAgent's knowledge or facts.
- **`example_posts`**: An array of strings representing example posts the dAgent might make.
- **`topics`**: An array of strings representing topics the dAgent might discuss.

### **Missions**

- **`task`**: A string representing the task or mission of the dAgent.
- **`system_reminder`**: A string representing a reminder or note for the system.
- **`toolset_cfg`**: An array of objects representing the toolset configurations, including:
    - **`name`**: The name of the toolset.
    - **`init_params`**: An object containing initialization parameters for the toolset, including:
        - **`exclude`**: An array of strings representing features or functions to exclude from the toolset.
- **`llm_cfg`**: An object representing the large language model (LLM) configuration, including:
    - **`name`**: The name of the LLM.
    - **`init_params`**: An object containing initialization parameters for the LLM, including:
        - **`model_name`**: The name of the LLM model.
        - **`max_tokens`**: The maximum number of tokens the LLM can generate.
        - **`model_kwargs`**: An object containing additional keyword arguments for the LLM model.
        - **`temperature`**: The temperature of the LLM, controlling the level of randomness in its responses.
        - **`max_retries`**: The maximum number of retries for the LLM.
- **`scheduling`**: An object representing the scheduling configuration, including:
    - **`interval_minutes`**: The interval at which the dAgent will perform its task, in minutes.


The configuration file is in JSON format and contains the following fields:

### **Characteristic**

- **`agent_personal_info`**: Contains information about the dAgent, including:
    - **`twitter_username`**: The Twitter username of the dAgent.
    - **`agent_name`**: The name of the dAgent.
- **`bio`**: An array of strings representing the dAgent's bio.
- **`lore`**: An array of strings representing the dAgent's lore or background information.
- **`knowledge`**: An array of strings representing the dAgent's knowledge or facts.
- **`example_posts`**: An array of strings representing example posts the dAgent might make.
- **`topics`**: An array of strings representing topics the dAgent might discuss.

### **Missions**

- **`task`**: A string representing the task or mission of the dAgent.
- **`system_reminder`**: A string representing a reminder or note for the system.
- **`toolset_cfg`**: An array of objects representing the toolset configurations, including:
    - **`name`**: The name of the toolset.
    - **`init_params`**: An object containing initialization parameters for the toolset, including:
        - **`exclude`**: An array of strings representing features or functions to exclude from the toolset.
- **`llm_cfg`**: An object representing the large language model (LLM) configuration, including:
    - **`name`**: The name of the LLM.
    - **`init_params`**: An object containing initialization parameters for the LLM, including:
        - **`model_name`**: The name of the LLM model.
        - **`max_tokens`**: The maximum number of tokens the LLM can generate.
        - **`model_kwargs`**: An object containing additional keyword arguments for the LLM model.
        - **`temperature`**: The temperature of the LLM, controlling the level of randomness in its responses.
        - **`max_retries`**: The maximum number of retries for the LLM.
- **`scheduling`**: An object representing the scheduling configuration, including:
    - **`interval_minutes`**: The interval at which the dAgent will perform its task, in minutes.
    

2. **Environment Variables**

To get started, you'll need to configure your environment variables. Here's a step-by-step guide:

### Step 1: Create a new `.env` file

A sample environment file, [.env-example](.env-example), is located in the root directory. Create a new `.env` file by copying the contents of `.env-example`:

### Step 2: Modify the `.env` file

Modify the values in the `.env` file to match your environment settings.

**Available Environment Variables**:
- **`IS_SANDBOX`**: Set to **`0`** for production or **`1`** for sandbox mode.

- **`ETERNALAI_URL`**: Set to the URL of your contract-based LLM backend API.

- **`ETERNALAI_API_KEY`**: Set to your API key for the contract-based LLM backend API.

- **`ETERNAL_X_API`**: Set to the URL of your Twitter API.

- **`ETERNAL_X_API_APIKEY`**: Set to your API key for the Twitter API.

Example `.env` file:

```bash
# Environment
IS_SANDBOX=0

# for contract based llm
ETERNALAI_URL=https://api.eternalai.org
ETERNALAI_API_KEY=PUT_YOUR_API_KEY_HERE

# twitter api
ETERNAL_X_API=https://agent.api.eternalai.org/api/developer
ETERNAL_X_API_APIKEY=PUT_YOUR_API_KEY_HERE
```


## Usage

The repository is created to be used for many purposes.

### Start the daemon to schedule missions periodically 

To start the dAgent, run the following command:

```bash
python daemon.py
```

### Integration

<!-- label -->


To integrate a dAgent to your existing `python project`, run the following commands to install the `dagent` package:

```bash
python -m pip install .

# or directly from github
python -m pip install git+https://github.com/eternalai-org/Eternals 
```

Then, in your python project, just import the `dagent`. Documentations for the app integration will be updated soon!

```python3
import dagent
```

### Chat with a dAgent 

To chat with a dAgent, please execute these following steps:

- Install the `dagent` package first: 

```bash
python -m pip install .
```

- And run:

```bash
python toolkits/chat-lite.py
```
