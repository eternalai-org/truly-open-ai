# Run a BitAI node

An BitAI node is similar to the Ordinals indexer. Anyone can run one, even on an old laptop. This guide will show you how to run one.

## Quick setup

To set up an BitAI node on a VPS or your local machine, run the following commands:

```
curl -s https://raw.githubusercontent.com/TrustlessComputer/trustless-node-easy/BitAI/setup-mainnet.sh -o setup-trustless.sh
chmod +x setup-trustless.sh
sh ./setup-trustless.sh
```

## Manual setup <a href="#quick-setup" id="quick-setup"></a>

### Install Bitcoin Core

[Bitcoin Core](https://bitcoincore.org/) provides both a Bitcoin full node and a wallet. Working with BitAI requires a Bitcoin full node with RPC enabled.

After installing Bitcoin Core, run `bitcoind` with `-server=1` :

```
./bitcoind -server=1
```

It may take some time for your Bitcoin full node to be fully synced.

### Install BitAI

Download BitAI pre-built binaries at this [download page](https://cdn.trustless.computer/releases/bitai/bitai).

After downloading the binary, run **chmod +x \<filename>** to allow executable permission. On MacOS, you may also have to go to Privacy & Security to allow the file to run.

Start your BitAI node:

```
./bitai -c <btc-node-cookie-file-path> -i <your-bitai-address> # specify your BitAI native address
```

Add these parameters if you set up your Bitcoin full node with a username and password.

```
./bitai -u <btc-node-username> -p <btc-node-password> -i <your-bitai-address>
```

Since BitAI reuses EVM, your BitAI native address is similar to an Ethereum address. You can create a new BitAI address with any EVM-compatible wallet. We recommend MetaMask.

### Add BitAI to MetaMask

In MetaMask, click on **Networks** -> **Add Network** -> **Add a network manually**. Use the following settings to point MetaMask to the BitAI running on your machine.

* **Name:** BitAI
* **URL:** http://localhost:10002
* **ChainID:** 22213
* **Symbol:** EAI

EAI is the native cryptocurrency of Eternal AI. Like ETH, you can use EAI to pay transaction fees, deploy smart contracts, and spend it in dapps.

### Setup BitAI Explorer <a href="#setup-bitcoin-virtual-machine-0-explorer" id="setup-bitcoin-virtual-machine-0-explorer"></a>

BitAI reuses [Blockscout](https://www.blockscout.com/) for blockchain exploration data such as blocks, transactions, and addresses.

Launch BitAI Explorer:

```
docker compose -f docker/docker-compose-blockscout.yml up
```

Open this URL on your browser:

```
http://0.0.0.0:4000
```

### Setup your EAI wallets <a href="#setup-your-bvm-wallets" id="setup-your-bvm-wallets"></a>

BitAI uses two wallets: a Bitcoin wallet and a native EAI wallet. You need funds in both wallets to operate your BitAI node.

#### **Bitcoin wallet**

A Bitcoin Core wallet named **`tc`** is automatically created by BitAI. Because BitAI uses Bitcoin Core to manage private keys, sign transactions, and broadcast transactions to the Bitcoin network, your **`tc`** wallet will need some sats.

Get a new address from your **`tc`** wallet and send it some funds:

```
./bitcoin-cli -rpcwallet=tc -getnewaddress
```

You can check the balance with:

```
./bitcoin-cli -rpcwallet=tc getbalances
```

#### **Native wallet**

Since BitAI reuses EVM, your native address is similar to an Ethereum address. You can create a new native address with any EVM-compatible wallet. We recommend MetaMask.

You will need some EAI, the cryptocurrency of Eternal AI. Like ETH, you can use EAI to deploy smart contracts, pay transaction fees, and spend in dapps.

The native address is the same as the one you use to run your BitAI.

```
./bitai -i <your-bitai-address>
```

You can check your EAI balance on MetaMask.
