# Proof-of-Compute

## Introduction

The primary challenge with building a decentralized AI infrastructure is the limited capability to run complex AI models due to the cost of on-chain storage and the computational limits of blockchains.

We propose a decentralized AI protocol that runs on a novel proof, Proof-of-Compute, where AI prompts are computed by a network of Eternal AI Nodes. The Eternal AI protocol provides compute service via a network of independent GPU nodes that does not rely on a single entity, where (1) creators publish their Eternals to earn money, (2) users pay to run a prompt on a specific Eternal, (3) Eternal AI nodes earn tokens by offering compute service, and (4) Eternal AI verifiers earn tokens by verifying the outputs.

## Prompts

A prompt is an input or set of instructions provided to an Eternal AI node to elicit a specific output. The nature and complexity of prompts can vary significantly depending on the type of Eternal and its intended application (Figure 1).

<figure><img src="../.gitbook/assets/image (21).png" alt=""><figcaption><p>Fig. 1. Diagram illustrating how a prompt, ’Create an image of a futuristic cityscape featuring flying cars,’ is processed by an Eternal AI node to generate the corresponding image. This demonstrates how prompts serve as instructions for Eternals to produce specific outputs.</p></figcaption></figure>

Here are some examples:

* Text Generation AI: Prompts in this context may consist of questions, sentences, or topics that guide the AI in generating textual responses. Examples: "Compose a narrative involving a dragon and a knight" or "Discuss the significance of the Renaissance period in European history."
* Image Generation AI: These prompts typically describe the visual elements the AI should depict in the generated image. Examples: "Create an image of a futuristic cityscape featuring flying cars" or "Produce a painting of a sunset over a mountainous landscape." Music Generation AI: Prompts in this domain might specify the genre, mood, or particular instruments to be used in the musical composition. Examples: "Compose a serene, ambient piece featuring piano and string instruments" or "Generate an energetic electronic dance track."
* Data Analysis AI: These prompts often involve specific queries or tasks that require the AI to analyze and interpret data. Examples: "Examine the sales data from the last quarter and identify emerging trends" or "Forecast the potential market growth for electric vehicles over the next five years."

## Eternal AI Nodes

> Anyone can run an Eternal AI node by staking 25,000 EAI.

Eternal AI nodes are the atomic compute unit of the Eternal AI network. They are nodes in the network. Anyone can run one or more Eternal AI nodes. An Eternal AI node typically has some sort of GPU to handle the computation. The job of the Eternal AI node is to take a prompt from a user, run some computations, and return the output (illustrated in Figure 2). Eternal AI nodes run a novel proof called Proof-of-Compute.

Conceptually, Proof-of-Compute is akin to Proof-of-Work. In this system, an Eternal AI node, the atomic compute unit of the Eternal AI network, must perform some computation on a prompt requiring heavy computation. In return, it earns fees and mines block rewards. The key difference from Proof-of-Work is that in Proof-of-Compute, Eternal AI nodes generate outputs that users can utilize, making the computation process more useful and less wasteful.

<figure><img src="../.gitbook/assets/image (22).png" alt=""><figcaption><p>Fig. 2. Diagram showing three Eternal AI nodes powered by three different GPU networks.</p></figcaption></figure>

## Eternal AI Verifiers

The problem, of course, is that the user can’t verify if an Eternal AI node computed on the desired Eternal or did any work at all. We introduce a special kind of Eternal AI node: verifiers.

Eternal AI verifiers first light-check the outputs, comparing the winning output against the later outputs. Eternal AI verifiers can dispute the output to earn fraud-detection rewards if there is any discrepancy. In the event of a dispute, the Eternal AI verifiers in the current committee will rerun the prompt, verify the result, and vote. If a supermajority (strictly more than two-thirds) of Eternal AI verifiers agrees the dispute is valid, the Eternal AI node that submitted the incorrect result will be slashed, and its staked tokens will be redistributed as a fraud-detection reward to the challenger.

While Eternal AI nodes handle the computational workload by performing AI inference tasks, Eternal AI verifiers oversee the integrity and correctness of these computations. They operate under the assumption that the Eternal AI nodes’ outputs are correct unless discrepancies arise, at which point they conduct thorough checks. This division of responsibilities ensures a balance between Eternal AI nodes for operational efficiency and Eternal AI verifiers for correctness.

## Eternal AI Creators

Anyone can submit an AI model to the network. The AI models are transformed into Eternals, which are cryptographically-secure AI agents running on the blockchain.

Eternals are deployed as standardized smart contracts, so any Eternal AI node can pick them up and run them. At the moment, the basic deep learning models and layers are supported, and more options will be added in the future.

## Network

Figure 3 illustrates a high-level overview of the Eternal AI protocol, which consists of the following components:

1. Model Publication: Creators deploy their Eternals on the network.
2. Prompt Submission: Users submit prompts to the Inference smart contract.
3. Prompt Distribution: Prompts are randomly assigned to three Eternal AI nodes, with the limit dynamically adjusted based on network performance.
4. Computation and Response: Eternal AI nodes process the prompts and return outputs.
5. Reward Allocation: The first Eternal AI node to return an output receives a reward, including the transaction fee and a block reward.
6. Verification and Penalties: Eternal AI verifiers continuously check outputs for accuracy, imposing penalties on dishonest nodes.

<figure><img src="../.gitbook/assets/image (23).png" alt=""><figcaption><p>Fig. 3. High-level overview of the Eternal AI protocol: From Eternal deployment by creators, prompt submission by users, random prompt distribution to nodes, computation and response, to reward allocation and verification by verifiers with penalties for dishonesty.</p></figcaption></figure>

## Incentives

#### For Eternal AI Creators

Eternal AI creators receive a portion of the transaction fee for each prompt that utilizes their Eternals.

#### For Eternal AI Nodes

The first AI node to respond with a valid solution earns the reward, which includes a portion of the transaction fee and a mining reward. Other submissions are accepted without a reward, but anyone can cross-check these solutions and dispute them if discrepancies are detected.

An Eternal AI node who fails to submit a result after requesting a task will be penalized and become inactive. An inactive Eternal AI node must wait one hour to rejoin the network.

#### For Eternal AI Verifiers

Eternal AI verifiers earn a portion of the transaction fees for each task they oversee. This ensures that they are compensated for the time and resources spent maintaining network integrity.

If an Eternal AI verifier initiates a successful dispute proving that an EternalAI node’s computation was incorrect, they receive a fraud-detection reward. This reward includes a substantial part of the Eternal AI node’s slashed stake in addition to a portion of the transaction fees and mining rewards linked to the disputed task.

Eternal AI verifiers who participate in the dispute resolution process by rerunning prompts and voting are eligible for additional rewards. These rewards are designed to encourage active participation and diligence in the consensus process.

Eternal AI verifiers play a critical role in ensuring the accuracy and integrity of computations within the Eternal AI network. To maintain a high standard of operation and to discourage negligence, Eternal AI verifiers are subject to the following penalties:

* Stake slashing for false disputes: If an Eternal AI verifier initiates a dispute that is subsequently proven unfounded (i.e., the accused miner’s computation is validated by consensus), the initiating Eternal AI verifier risks having a portion of their stake slashed. This penalty discourages frivolous or malicious disputes.
* Inactivity penalties: Eternal AI verifiers are expected to remain online and actively participate in network duties. Those who consistently are offline or fail to participate in the validation and consensus processes may lose a portion of their stake over time. This penalty is typically a small percentage but can accumulate if the behavior persists.
* Consensus non-participation penalty: Eternal AI verifiers who fail to participate in the dispute resolution process when selected are subject to penalties. This ensures that all Eternal AI verifiers remain engaged and that consensus decisions are reached quickly and efficiently.
* Improper conduct penalty: Any actions by Eternal AI verifiers that are deemed to undermine the network’s integrity or the fairness of computations—such as attempting to manipulate results, colluding with Eternal AI nodes, or other forms of corruption—will result in severe penalties, including possible expulsion from the network and forfeiture of the entire stake.
