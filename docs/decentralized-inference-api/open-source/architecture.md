# Architecture

In this document, we describe the components of an agent created via the **EternalAI** open-source platform and explain how these components interact with each other.

<figure><img src="../../.gitbook/assets/image (5).png" alt=""><figcaption></figcaption></figure>

## Mission Manager & Configuration

This is a periodic job that schedules the agent to execute missions defined in the configuration file. Each time the agent performs a mission, it proceeds through a series of steps in a chain-of-thought process sequentially (i.e., each step uses the output of the previous step as input when interacting with the LLM model) to achieve the mission's goal using predefined actions configured for the agent.

The Mission Manager utilizes information defined in the configuration file to execute missions as follows:

* **Characteristic**, which includes:
  * `system_prompt`
* **Mission**, which includes:
  * `task`
  * `system_reminder`
  * `toolset_cfg`
  * `llm_cfg`
  * `scheduling`
    * `interval_minutes`

Code reference [here](https://github.com/TrustlessComputer/priv-eternal-agent/blob/43d33fa4fc41e6cff82a401ac8fa2f7f8a62cf31/eternal_agent/service.py#L115).

## LLM

At each step of execution, the agent calls an LLM model to receive an instruction. The instruction includes one of the following action types: `thought`, `action`, or `final_answer`, which directs the agent on its next step.

Since the EternalAI platform is powered by a decentralized AI infrastructure, all LLM interactions are performed asynchronously through smart contract calls. The framework generates an inference request to a smart contract corresponding to an LLM model deployed on a blockchain whenever the agent interacts with the model. The smart contract returns an `inference_id`, which is used to retrieve the inference response (instruction) submitted by the LLM model’s miners.

Initially, the LLM module supported only the Hermes model. However, the LLM provider is implemented in an abstracted manner, allowing future contributors to easily add support for new LLM models.

Code reference [here](https://github.com/TrustlessComputer/priv-eternal-agent/tree/main/eternal_agent/llm).

## Tools (Actions)

Based on the LLM model's instruction, the agent executes a corresponding action.

The **Tools** module defines toolsets for each specific use case. For example:

* `twitter_toolset` for a Twitter agent.
* `trading_toolset` for a trading agent.

Similar to the LLM module, contributors can easily add new toolsets to the module, enabling users to configure additional actions for their use cases.

Code reference [here](https://github.com/TrustlessComputer/priv-eternal-agent/tree/main/eternal_agent/tools).
