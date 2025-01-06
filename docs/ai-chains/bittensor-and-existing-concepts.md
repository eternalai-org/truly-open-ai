# Bittensor and existing concepts

Bittensor is a radical development in building AI-specific blockchains. On the Bittensor network, AI chains are called subnets, each focusing on a specific AI use case. However, there are a few critical centralized components.

<figure><img src="../.gitbook/assets/image (47).png" alt="" width="563"><figcaption><p>The Bittensor architecture</p></figcaption></figure>

Eternal AI intends to create an alternative protocol for building AI chains, providing a different set of tradeoffs with particular emphasis on decentralization, scalability, and programmability.&#x20;

Architecturally, Eternal AI is similar to Bittensor but with a few key differences:

* Bittensor uses a Proof-of-Authority, while Eternal AI inherits Bitcoin’s Proof-of-Work security.
* Bittensor implements subnets with 64 validators per subnet, while Eternal AI implements AI chains with ZK rollups on Bitcoin. Each AI chain is secured by math with the power of ZK proofs instead of validators.
* Bittensor relies on validators to run ad hoc validation programs to verify miners' work, while Eternal AI uses trustless verification with the power of smart contracts.
* Bittensor implements apps, while Eternal AI implements decentralized apps.
