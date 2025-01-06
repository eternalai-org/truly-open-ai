# What are Neurons?

<figure><img src="../.gitbook/assets/image (13).png" alt=""><figcaption><p>A Neuron in every home.</p></figcaption></figure>

## Overview

Neurons are nodes of the Eternal AI network. They are the atomic compute unit powering decentralized inference across all blockchains including Bitcoin, Ethereum, and Solana.

A Neuron typically has some sort of GPU to handle the high-performance on-chain computation. Neurons run a novel proof called [Proof-of-Compute](../eternals/proof-of-compute.md).

It's permissionless to join the Eternal AI network. Anyone can run a Neuron and mine EAI.

## Neuron types

There are two types of Neurons: Compute Neurons and Verify Neurons.

### Compute Neurons

The job of a Compute Neuron is to take a prompt from a user, run some computations, and return the output. In return, it earns fees and mines block rewards.

The key difference from Proof-of-Work is that in Proof-of-Compute, compute nodes generate outputs that users can utilize, making the computation process more useful and less wasteful.

### Verify Neurons

The problem, of course, is that the user can’t verify if a Compute Neuron computed on the desired Eternal or did any work at all. We introduce a special kind of Neuron: Verify Neuron.

Verify Neurons can dispute the output to earn fraud-detection rewards if there is any discrepancy. The Verify Neuron that submitted the incorrect result will be slashed, and its staked tokens will be redistributed as a fraud-detection reward to the challenger.

While the Compute Neurons handle the computational workload by performing AI inference tasks, Verify Neurons oversee the integrity and correctness of these computations. This division of responsibilities ensures a balance between Compute Neurons for operational efficiency and Verify Neurons for correctness.
