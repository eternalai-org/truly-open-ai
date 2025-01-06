# BitAI Virtual Machine

Given Bitcoin's lack of smart contract support, we deployed a Turing-complete virtual machine explicitly tailored for AI applications on Bitcoin. This is made possible by implementing the [BVM](https://bvm.network) sharding solution.

At its core, BitAI (Bitcoin AI Virtual Machine) is a state machine that utilizes Bitcoin as a data layer to achieve transaction-level consensus. This approach allows us to create a state machine leveraging Bitcoin's security to execute AI inferencing code.

BitAI does not require a sidechain or a Bitcoin L2 and can run today without any changes to Bitcoin.

<figure><img src="../../.gitbook/assets/image (44).png" alt=""><figcaption><p>The BitAI Virtual Machine architecture</p></figcaption></figure>

BitAI consists of four main components.

### **Local Mempool**

As part of the system, each BitAI node maintains a collection of transactions that the user sends to the node in a local mempool. The mempool helps verify the transaction's validation before it is written to the Bitcoin network.

This mempool is local because it is not shared with other BitAI nodes. The transaction is only exposed publicly when written to a Bitcoin transaction (the process of Inscriber) and broadcast to the Bitcoin network.

### **Inscriber**

Inscriber enables us to embed BitAI transactions into Bitcoin transactions. To achieve this, we use a technique similar to that used by Ordinals to inscribe data into the Bitcoin network.

{% hint style="info" %}
There is a significant difference between Ordinals and BitAI. Ordinals let people inscribe data like images and texts onto Bitcoin. BitAI lets developers build AIs on Bitcoin.
{% endhint %}

The BitAI transaction data is essentially embedded into a Bitcoin transaction via the witness data field. This embedding does not affect the verification process or the transaction logic.

The witness data field is only exposed as an unlock script for the previous output. Pushing data to a Bitcoin transaction requires a two-phase transaction process.

The first phase is the commit transaction, which involves creating an unspent transaction output (UTXO) that indicates the witness/script hash for spending. In the second phase, the revealing transaction exposes the witness data/script containing the BitAI transactions.

Similar to Ordinals, the BitAI transactions are serialized using data pushes within unexecuted conditionals, called "envelopes". Envelopes consist of an `OP_FALSE OP_IF … OP_ENDIF` wrapping any number of data pushes.

A text inscription containing BitAI transactions data (which is split into n chunks) is serialized as follows:

```
OP_FALSE
OP_IF
  OP_PUSH "ord"
  OP_PUSH 1
  OP_PUSH "text/plain;charset=utf-8"
  OP_PUSH 0
  OP_PUSH "BitAI transactions' chunk 1"
  OP_PUSH "BitAI transactions' chunk 2"
  ...
  OP_PUSH "BitAI transactions' chunk n"  
OP_ENDIF
```

### Indexer <a href="#txreader" id="txreader"></a>

This module is responsible for filtering BitAI transactions in every new Bitcoin block and ensuring that the state of BitAI is consistent across all BitAI nodes, even in the event of a reorg.

For each Bitcoin block, the Indexer will first filter the BitAI transactions, sort them by gas fees, and then import them to BitAI as a new BitAI block. As Bitcoin has an immutable, deterministic order, every BitAI node that runs an honest codebase will have the same state.

Bitcoin reorg is a process that occurs when two or more miners simultaneously add different blocks to the blockchain. This can cause a fork in the chain, where different parts of the network have different versions of the blockchain. Indexer is designed to handle Bitcoin reorgs by reverting the BitAI state to the point of the forked branch and then re-importing the valid blocks into the reverted state.

### State Machine

BitAI is an Ethereum-compatible virtual machine. It has been configured to support larger transaction sizes and higher block gas limits, enabling us to support AI applications, which often require more computation and larger transaction sizes.
