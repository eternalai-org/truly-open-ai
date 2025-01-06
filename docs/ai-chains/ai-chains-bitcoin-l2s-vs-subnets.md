# AI chains: Bitcoin L2s vs Subnets

Bittensor implements AI chains as subnets. There are 64 subnets, each with 64 validators, and they are implemented as proof-of-stake blockchains.

Eternal AI implements AI chains as AI-specific Bitcoin L2s. The number of Bitcoin L2s is infinite. Creating one is permissionless; anyone can spin up one with a few clicks. These AI chains are ZK rollups on Bitcoin and secured by cryptography and math, not validators.

<figure><img src="../.gitbook/assets/image (10).png" alt=""><figcaption><p>Subnets vs Bitcoin L2s</p></figcaption></figure>

## AI-specific, ZK-rollup Bitcoin L2s

AI chains are AI-specific Bitcoin L2s. They are implemented as ZK rollups on Bitcoin.

<figure><img src="../.gitbook/assets/image (46).png" alt=""><figcaption></figcaption></figure>

The steps of a ZK rollup from an AI chain to Bitcoin are as follows:

1. Users create transactions and send them to the AI Chain.
2. The AI Chain processes the block of transactions and forwards it to the AI Chain Prover.
3. The AI Chain Prover generates cryptographic proof of the block’s execution and a commitment and then submits these to BitAI VM via a transaction.
4. The AI Chain Sequencer also submits the block’s transactions to an alternative DA by sending another transaction to a supported DA layer.
5. BitAI's Local Mempool pre-executes these transactions to filter out invalid ones.
6. BitAI's Inscriber inscribes valid transactions on Bitcoin.
7. BitAI's Indexer pulls newly created Bitcoin blocks, parses related BitAI transactions and passes them to the State Machine for execution.
8. BitAI's State Machine executes all transactions passed from the TX reader. In the ZK Rollup flow, a verifier contract deployed on BitAI verifies the state transitions claimed by the AI Chain Sequencer.

## Miners

Neurons are miners of the Eternal AI network. A Neuron typically has some sort of GPU to handle high-performance on-chain computation.

Neurons run a novel proof called [Proof-of-Compute](../eternals/proof-of-compute.md) that provides trustless verification.

