# Specification

## Eternals

Eternals are cryptographically secure AI models programmed as smart contracts.&#x20;

Here are the key fields in the smart contract:

```solidity
struct Eternal {
    uint256 fee;
    bytes[] systemPrompt;
    bytes[] foundationModel;
}
```

* **`fee`**:  fee collected from users who interact with the Eternal contract
* **`systemPrompt`**: system prompt used to guide the Eternal’s behavior and responses
* **`foundationModel`**: foundational model upon which the Eternal operates

To create a new Eternal, call the `mint` function in the Eternal smart contract.

```solidity
function mint(
    address _to,
    string calldata _uri,
    bytes calldata _data,
    uint _fee
)
```

* **`_to`**: address that will own the newly created Eternal
* **`_uri`**: the URI pointing to the Eternal’s metadata, including details like name and description
* **`_data`**: initial system prompt that guides the Eternal's behavior and responses
* **`_fee`**: fee required from users who interact with this Eternal

{% hint style="info" %}
The Eternal smart contract is designed to be as flexible as possible. For example, the metadata could be an inscription on Bitcoin, a file on Filecoin, or even a file on AWS. We leave it up to the Eternal creators to decide what works for their use cases. For maximum security, we recommend writing the data to Bitcoin. For less critical applications, we recommend writing the data to Filecoin.
{% endhint %}

## Decentralized Inference

To perform decentralized inference, call the `infer` function in the Eternal smart contract.

```solidity
function infer(
    uint256 _eternalId, 
    bytes calldata _calldata
)
```

* **`_eternalId`**: id of the Eternal handling this inference request
* **`_calldata`**: prompt provided by the user, which the Eternal will interpret and respond to based on its foundational model and system prompt

In the `infer` function, a request is created with an Eternal’s system prompt attached. This request is then forwarded to the foundation model’s miners for processing.

The miners responsible for handling the inference requests are managed by a separate contract called `WorkerHub`. This contract allows users to stake to become miners, manage worker assignments, and submit solutions for inference requests.

Key structs in the `WorkerHub`:

```solidity
assignmentsstruct Assignment { 
    uint256 inferenceId; 
    bytes32 commitment; 
    bytes32 digest; 
    uint40 revealNonce; 
    address miner;
    AssignmentRole role; 
    bytes output;
}
```

* **`inferenceId`**: id for the inference request that the WorkerHub assigns to miners
* **`commitment`**: commitment hash that a miner submits to prove their intent to process the inference request securely
* **`digest`**: hash of the response provided by a miner
* **`revealNonce`**: random seed used to verify the miner’s commitment during the reveal phase
* **`miner`**: address of the miner assigned to the inference request
* **`role`**: The role, either as a miner or as a verifier.
* **`output`**: miner’s response to the prompt

```solidity
struct Inference {
    uint256[] assignments;
    bytes input;
    uint256 value;
    uint40 submitTimeout;
    uint40 commitTimeout;
    uint40 revealTimeout;
    InferenceStatus status;
    address creator;
    address processedMiner;
}
```

* **`assignments`**: assignments that specify the miners handling this inference request
* **`input`**: user’s prompt
* **`value`**: total inference fee for processing this request
* **`submitTimeout`**: time limit for the first miner to submit the response
* **`commitTimeout`**: time limit for the second and third miners to submit their commitments
* **`revealTimeout`**: time limit for the second and third miners to reveal their responses
* **`status`**: current processing status of the inference, indicating phases like `Processing`, `Commit`, `Reveal`, `Processed`, or `Failed`
* **`creator`**: address of the user who initiated the inference request
* **`processedMiner`**: address of the first miner to process the inference request

These structs represent the data for each assignment and inference request, which are integral to the workflow of the **`WorkerHub`** smart contract.

The key functions of the `WorkerHub` smart contract are:

```solidity
function seizeMinerRole(uint256 _assignmentId)
```

**`seizeMinerRole`** function is called by an assigned miner to obtain the miner role.

* **`_assignmentId`**: assignment id

```solidity
function submitResponse(uint256 _assignmentId, bytes calldata _data)
```

**`submitResponse`** function is called by a miner who has obtained the miner role for processing the inference request.

* **`_assignmentId`**: assignment id&#x20;
* **`_data`**:  the miner's response to the user's prompt

```solidity
function commit(uint256 _assignmentId, bytes32 _commitment)
```

**`commit`** function is called by the verifiers.

* **`_assignmentId`**: assignment id&#x20;
* **`_commitment`**:  commitment hash, calculated by hashing the response, the miner’s address, and a random nonce. This commitment is later used in the reveal phase to verify the miner's response.

```solidity
function reveal(uint256 _assignmentId, uint40 _nonce, bytes calldata _data)
```

**`reveal`** function is called by the verifiers.

* **`_assignmentId`**: assignment id&#x20;
* **`_nonce`**:  random seed (nonce) initially used to create the commitment hash
* **`_data`**: The response that aligns with the initial commitment hash, providing transparency and proof of the miner’s response

The processing of an inference begins by randomly selecting a few miners from the miner pool to handle the prompt.&#x20;

The mining reward goes to the fastest miner who calls the `seizeMinerRole` function to claim the task and then calls the `submitSolution` function to provide the solution.&#x20;

The other miners become verifiers. They then call the `commit` function and the `reveal` function to submit their verification votes for the fastest miner's solution. The `reveal` function verifies the necessary revealed votes and finalizes the solution.

[Learn more about Proof-of-Compute →](proof-of-compute.md)
