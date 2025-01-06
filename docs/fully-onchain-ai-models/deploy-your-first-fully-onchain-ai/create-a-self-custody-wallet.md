# Create a self-custody wallet

To interact with the Eternal AI blockchain, you need a self-custody wallet. More than just a login, it gives you full control over your Eternals. You can use the wallet to hold, send, and receive Eternals permissionlessly.

## Create a wallet

To create a wallet, run:

```bash
eai wallet create
```

or you can restore an existing wallet:

```bash
eai wallet import -p <PRIVATE_KEY> --network "mainnet"
```

_**Note:**_

* The `-p` parameter means private key, and it is **required**. If you don't provide a private key, the system will throw an error.
* The `--network` parameter means the network you want to use, and it is **optional**. If you don't provide a network, the system will use the `testnet` network as default.

To see all available wallet options, you can run:

```bash
eai wallet --help
```

## Add funds to your wallet

After creating a wallet, you will need some [EAI](../../eai/utilities.md), the native cryptocurrency of the Eternal AI network, to cover the network fees for deploying your first Eternal.

{% hint style="info" %}
It is free and fast to get some testnet EAI from the [Eternal AI Faucet](https://eternalai.org/faucet) for development. On the testnet, you can simulate using all of Eternal AI's features without using real money. Once you are ready to deploy on the mainnet, buy some EAI on [Uniswap](https://app.uniswap.org/swap?outputCurrency=0xa84f95eb3DaBdc1bbD613709ef5F2fD42CE5bE8d).
{% endhint %}

Get a new deposit address and send it some EAI.

```bash
eai wallet deposit
```

You can see pending transactions with:

```bash
eai wallet transactions
```

Once the transaction is confirmed, you should be able to see the updated balance with:

```bash
eai wallet balance
```
