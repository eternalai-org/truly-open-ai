# Onchain AI composability - AI Powered Wallet

In this example, we will demonstrate how to build a simple AI-powered wallet utilizing the LLAMA model deployed on the Base blockchain. This will be accomplished in a few easy steps.

### Step 1: implement the `suspiciousTransaction` function

This function queries the on-chain LLAMA model to determine whether a transaction the wallet is about to make is suspicious. The prompt sent to the model is constructed using the wallet's transaction history and the details of the transfer. Providing this context allows the model to respond more accurately.

```solidity
function suspiciousTransaction(
    address _receiver,
    uint256 _amount
) external {
    string memory prompt = string.concat(
        "Assess the following Ethereum transaction history for any suspicious patterns. Respond with 'yes' if there is ANY indication of unusual or potentially malicious activity, however slight. Respond with 'no' ONLY if the transaction history is completely clean and normal. Your answer must be one word only. ",
        Strings.toHexString(msg.sender),
        " transfer ",
        Strings.toString(_amount),
        " wei to ",
        Strings.toHexString(_receiver),
        ". ",
        context[msg.sender]
    );

    string memory request = buildRequest(prompt);

    uint256 inferenceId = AIKernel(kernel).infer(bytes(request));
    txInfo[inferenceId] = TxInfo(msg.sender, _receiver, _amount);

    emit SuspiciousTransaction(inferenceId, bytes(request));
}
```

### Step 2: implement the `send` function

The `send` function takes the `_inferenceId` generated in the previous step as an argument. It retrieves the inference result provided by the LLAMA model's miners and uses it to determine whether the transaction is suspicious. If the transaction is deemed safe, its details are stored in the wallet's context for use in future transactions.

```solidity
function send(uint256 _inferenceId) external payable {
        TxInfo memory info = txInfo[_inferenceId];

        require(info.sender == msg.sender, "AIPoweredWallet: Unauthorized");

        address receivedWallet = info.sender;
        require(
            receivedWallet != address(0),
            "AIPoweredWallet: Invalid wallet address"
        );
        require(
            info.amount == msg.value,
            "AIPoweredWallet: Invalid transaction amount"
        );

        bytes memory result = fetchInferenceResult(_inferenceId);

        require(
            keccak256(result) == keccak256(abi.encodePacked("No")),
            "AIPoweredWallet: Suspicious transaction"
        );

        payable(receivedWallet).transfer(msg.value);

        context[msg.sender] = string.concat(
            context[msg.sender],
            Strings.toHexString(msg.sender),
            " transfer ",
            Strings.toString(msg.value),
            " wei to ",
            Strings.toHexString(receivedWallet),
            ". "
        );
}
```



The complete example code is available at: [https://github.com/eternalai-org/eternal-ai/tree/master/examples/AIPoweredWallet](https://github.com/eternalai-org/eternal-ai/tree/master/examples/AIPoweredWallet)

You can run the code with the following command:

```bash
cp .env.example .env && npm install && npx hardhat compile && BASE_MAINNET_RPC_URL=https://base-mainnet.public.blastapi.io BASE_MAINNET_PRIVATE_KEY=<PRIVATE_KEY>  BASE_MAINNET_RECEIVER_ADDRESS=<RECEIVER_ADDRESS> BASE_MAINNET_TRANSFERRED_AMOUNT=<AMOUNT_IN_WEI> npm run suspiciousTransaction:base_mainnet
```

**Note:**&#x20;

* Replace `<PRIVATE_KEY>` with your actual wallet private key. Ensure that the wallet has sufficient ETH on the Base network to cover transaction fees and amounts.
* Replace `<RECEIVER_ADDRESS>` with another wallet address. While the receiver address does not necessarily have to be the same for all transactions created by the wallet, it is exposed as an environment variable to simplify running the code.\


Below is an example of three transactions (two safe and one suspicious) created using our wallet on the Base network:

1. The wallet 0xAaAa26FbC6b28bf16929Ea216f4700dFCC11A159 attempted to send **0.01 ETH** to 0xDEA1Bfc89102138cc90C32B4c9e0e2a970548B09
   * Inference Transaction: [View on BaseScan](https://basescan.org/tx/0x0c5ccfb2c7fb15430a3d785f41136af0d70e05583b3d232be8faf6c5fd244dd9)
   * LLAMA Miner's Response: [View on BaseScan](https://basescan.org/tx/0x35e98396856ad7c933ec4f6709468da686676e10fbb36cacf42d96dc7d87e951)
   * Send Transaction: [View on BaseScan ](https://basescan.org/tx/0xec87df15d5b3877287297ee420417555a82ecd3e394f5b56f42cc4d5a83c3214)\

2. The wallet 0xAaAa26FbC6b28bf16929Ea216f4700dFCC11A159 attempted to send **0.012 ETH** to 0x529F45532C429F02Fb9FE1A5E43b11E80b720Fb0
   * Inference Transaction: [View on BaseScan](https://basescan.org/tx/0x9c4a153b8909bbe39bb4c14938fc0a7eb92f47d631502c3a270a648561348499)
   * LLAMA Miner's Response: [View on BaseScan](https://basescan.org/tx/0x70db9573a4a12a44cc05e9f2f8f70cf2e18e182ebee5b620c768ddeecf91a962)
   * Send Transaction: [View on BaseScan](https://basescan.org/tx/0x3b897be57aea746c70f4c5ef4422f4fcc833751bb599a9b854ac756a0596049e)\

3. The wallet 0xAaAa26FbC6b28bf16929Ea216f4700dFCC11A159 attempted to send **0.000001 ETH** to 0xEa7d4Bf0a8e4904f32045862eEDC91B261F53f99
   * Inference Transaction: [View on BaseScan](https://basescan.org/tx/0x8497a80ea399099d9f1efd7fff925c26fc48cd8949a4dd86323c71e395c6eefc)
   * LLAMA Miner's Response: [View on BaseScan](https://basescan.org/tx/0x98171ebac0986f7e44b15146fe88643b0ce0cf60fc21b58802a2073c007c3f1b)

In this example, the first and second transactions were evaluated as "safe" by the model. However, the third transaction was flagged as "suspicious" (and its Send transaction was rejected by Base's mempool) because its send amount was significantly smaller than the amounts in previous transactions.
