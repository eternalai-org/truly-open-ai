# Eternals on Bitcoin

## Eternal AI as a new Bitcoin metaprotocol

Metaprotocols have been powerful tools for extending Bitcoin's functionality. [Ordinals](https://ordinals.com/) is the non-fungible token metaprotocol on Bitcoin, [Runes](https://docs.ordinals.com/runes.html) is the fungible token metaprotocol on Bitcoin, and [BVM](https://bvm.network) lets developers write smart contracts on Bitcoin.

And now, Eternal AI is the metaprotocol for AI on Bitcoin.&#x20;

## Eternals as new Bitcoin-native digital assets

Metaprotocols often introduce new Bitcoin-native digital assets. Ordinals introduced NFTs on Bitcoin. Runes introduce tokens on Bitcoin. BVM introduced dapps on Bitcoin.

And now, Eternal AI has introduced Eternals — AIs living forever on Bitcoin.

Technically, Eternals are similar to Ordinals in many ways. They both embed data into Taproot transactions, do not require a sidechain or Bitcoin L2, and work without any changes to Bitcoin.

However, there is one big difference: Eternals are AIs on Bitcoin, while Ordinals are JPEGs on Bitcoin.

## BitAI as a new Bitcoin indexer

Metaprotocols often require external indexers on top of Bitcoin. Eternal AI also has an indexer that scans every new Bitcoin block for transactions related to Eternal AI.

BitAI is the indexer of Eternal AI. Unlike Ordinals indexer, BitAI is a Turing-complete virtual machine that enables writing Solidity smart contracts on Bitcoin.

[Learn more about BitAI Virtual Machine →](bitai-virtual-machine.md)

## Eternal Smart Contracts

Eternals are cryptographically secure AI models programmed as smart contracts. Their contracts and all related transactions are inscribed on Bitcoin via BitAI.

Here are the key fields in the smart contract:

```solidity
struct Eternal {
    uint256 fee;
    bytes[] sysPrompt;
}
```

It contains two main functions: `mint` for creating a new Eternal and `infer` for performing decentralized inference on a prompt.

```solidity
function mint(address _to, string calldata _uri, bytes calldata _data, uint _fee) external {}
function infer(uint256 _agentId, bytes calldata _calldata, string calldata _externalData) external {}
```

[Learn more about the Eternal specification →](../../eternals/specification.md)
