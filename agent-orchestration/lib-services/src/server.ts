import * as dotenv from "dotenv";
import express from "express";
import bs58 from "bs58";

import { GetAddressPrk, StoreAddress } from "./core";
import { decodeBase58, decodeBase64, encodeBase58, ethers } from "ethers";
import * as solanaWeb3 from "@solana/web3.js";
import axios from "axios";
import * as cheerio from "cheerio";
import { neynarAPIRouterInit } from "./neynar";
import { createUmi } from '@metaplex-foundation/umi-bundle-defaults';
import { GetProgramAccountsFilter } from "@solana/web3.js";
import { AuthorityType, createSetAuthorityInstruction, createTransferInstruction, getOrCreateAssociatedTokenAccount, TOKEN_PROGRAM_ID } from "@solana/spl-token";
import { cloneAndSnakeCaseFields } from "./utils";
import { raydiumAPIRouterInit } from "./raydium";
import { createAndMint, mplTokenMetadata, TokenStandard } from "@metaplex-foundation/mpl-token-metadata";
import { createSignerFromKeypair, generateSigner, percentAmount, signerIdentity } from "@metaplex-foundation/umi";
import { SOLANA_RPC_ENDPOINT } from "./contants";

dotenv.config();

const app = express();
app.use(express.json({ limit: '50mb' }));
app.use(express.urlencoded({ limit: '50mb' }));


// SOLANA
app.post('/solana/address', async (req, res) => {
    try {
        const keyPair = solanaWeb3.Keypair.generate()
        const address = keyPair.publicKey.toBase58()
        const privateKey = encodeBase58(keyPair.secretKey)
        var fromKey = solanaWeb3.Keypair.fromSecretKey(ethers.toBeArray(decodeBase58(privateKey)));
        if (fromKey.publicKey.toBase58() != address) {
            throw 'prk not valid'
        }
        await StoreAddress(
            'mainnet',
            address,
            privateKey,
        )
        const prk = await GetAddressPrk(address)
        if (privateKey != prk) {
            throw 'prk not found'
        }
        var fromKey = solanaWeb3.Keypair.fromSecretKey(ethers.toBeArray(decodeBase58(prk)));
        if (fromKey.publicKey.toBase58() != address) {
            throw 'prk not valid'
        }
        res
            .status(200)
            .send({
                result: {
                    address: address,
                },
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error
                }
            })
    }
});

app.get('/solana/blockheight', async (req, res) => {
    try {
        let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
        const number = await connection.getBlockHeight("finalized")
        res
            .status(200)
            .send({
                result: Number(number),
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error
                }
            })
    }
});

app.get('/solana/balance/:address', async (req, res) => {
    try {
        const address = new solanaWeb3.PublicKey(req.params.address)
        let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
        const balance = await connection.getBalance(address)
        res
            .status(200)
            .send({
                result: Number(balance),
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error
                }
            })
    }
});

app.post('/clean-html', async (req, res) => {
    const { url } = req.body;
    let { html_data } = req.body;

    try {
        if (url != "") {
            const result = await axios.get(url);
            html_data = result.data
        }
        // const result = await axios.get(url);
        console.log(html_data)
        const $ = cheerio.load(html_data);
        const text = $('html').prop('innerText');
        let replacedText = text.replace(/[\n\t]/g, '');
        replacedText = replacedText.replace(/  +/g, ' ');
        console.log(replacedText);
        res.status(200).send({ result: replacedText })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error
                }
            })
    }
});

app.get('/solana/token-info/:mint', async (req, res) => {
    try {
        const mint = new solanaWeb3.PublicKey(req.params.mint)
        let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT, "confirmed");
        let mintInfo = await connection.getParsedAccountInfo(
            new solanaWeb3.PublicKey(mint)
        )
        res
            .status(200)
            .send({
                result: mintInfo.value,
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error
                }
            })
    }
});

app.get('/solana/balances/:address', async (req, res) => {
    try {
        const address = new solanaWeb3.PublicKey(req.params.address)
        let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
        const filters: GetProgramAccountsFilter[] = [
            {
                dataSize: 165,
            },
            {
                memcmp: {
                    offset: 32,
                    bytes: address.toString(),
                },
            }];
        const accounts = await connection.getParsedProgramAccounts(
            TOKEN_PROGRAM_ID,
            { filters: filters }
        );
        const balances = []
        const balance = await connection.getBalance(address)
        balances.push({
            is_native: true,
            mint: '',
            owner: address,
            state: 'initialized',
            token_amount: {
                amount: balance.toString(),
                decimals: 9,
                ui_amount: Number((balance / 1e9).toFixed(9)),
                ui_amount_string: ((balance / 1e9).toFixed(9)).toString(),
            }
        })
        accounts.forEach((account, i) => {
            const parsedAccountInfo: any = account.account.data;
            if (parsedAccountInfo.parsed.info.tokenAmount.uiAmount > 0) {
                balances.push(cloneAndSnakeCaseFields(parsedAccountInfo.parsed.info))
            }
        });
        res
            .status(200)
            .send({
                result: balances,
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error
                }
            })
    }
});

app.post('/solana/create-pumfun', async (req, res) => {
    try {
        let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT, "confirmed");
        const { address, name, symbol, description, twitter, telegram, website, image_base64, amount } = req.body;
        const privateKey = await GetAddressPrk(address)
        const signerKeyPair = solanaWeb3.Keypair.fromSecretKey(bs58.decode(privateKey));
        // Generate a random keypair for token
        const mintKeypair = solanaWeb3.Keypair.generate();
        // Define token metadata
        const formData = new FormData();
        if (image_base64 && image_base64 != '') {
            const blob = new Blob([decodeBase64(image_base64)], { type: "application/octet-stream" });
            formData.append("file", blob, "icon.png")
        }
        formData.append("name", name)
        formData.append("symbol", symbol)
        formData.append("description", description)
        formData.append("twitter", twitter)
        formData.append("telegram", telegram)
        formData.append("website", website)
        formData.append("showName", "true")
        // Create IPFS metadata storage
        const metadataResponse = await fetch("https://pump.fun/api/ipfs", {
            method: "POST",
            body: formData,
        });
        const metadataResponseJSON = await metadataResponse.json();
        // Get the create transaction
        const response = await fetch(`https://pumpportal.fun/api/trade-local`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                "publicKey": signerKeyPair.publicKey.toBase58(),
                "action": "create",
                "tokenMetadata": {
                    name: metadataResponseJSON.metadata.name,
                    symbol: metadataResponseJSON.metadata.symbol,
                    uri: metadataResponseJSON.metadataUri
                },
                "mint": mintKeypair.publicKey.toBase58(),
                "denominatedInSol": "true",
                "amount": Number(amount), // dev buy of 1 SOL
                "slippage": 5,
                "priorityFee": 0.00005,
                "pool": "pump"
            })
        });
        var signature: solanaWeb3.TransactionSignature
        if (response.status === 200) { // successfully generated transaction
            const data = await response.arrayBuffer();
            const tx = solanaWeb3.VersionedTransaction.deserialize(new Uint8Array(data));
            tx.sign([mintKeypair, signerKeyPair]);
            signature = await connection.sendTransaction(tx)
        } else {
            throw new Error(response.statusText);
        }
        res
            .status(200)
            .send({
                result: {
                    'signature': signature,
                    "mint": mintKeypair.publicKey.toBase58(),
                },
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error['message']
                }
            })
    }
});

app.post('/solana/trade-pumfun', async (req, res) => {
    try {
        let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT, "confirmed");
        const { address, action, mint, amount, pool } = req.body;
        const privateKey = await GetAddressPrk(address)
        const signerKeyPair = solanaWeb3.Keypair.fromSecretKey(bs58.decode(privateKey));
        // Get the create transaction
        const response = await fetch(`https://pumpportal.fun/api/trade-local`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                "publicKey": signerKeyPair.publicKey.toBase58(),
                "action": action,
                "mint": mint,
                "denominatedInSol": action == "buy" ? "true" : "false",
                "amount": Number(amount),
                "slippage": 5,
                "priorityFee": 0.00005,
                "pool": pool
            })
        });
        var signature: solanaWeb3.TransactionSignature
        if (response.status === 200) { // successfully generated transaction
            const data = await response.arrayBuffer();
            const tx = solanaWeb3.VersionedTransaction.deserialize(new Uint8Array(data));
            tx.sign([signerKeyPair]);
            signature = await connection.sendTransaction(tx)
        } else {
            throw new Error(response.statusText);
        }
        res
            .status(200)
            .send({
                result: signature,
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error['message']
                }
            })
    }
});

async function getNumberDecimals(mintAddress: string): Promise<number> {
    let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
    const info = await connection.getParsedAccountInfo(new solanaWeb3.PublicKey(mintAddress));
    const result = (info.value?.data as solanaWeb3.ParsedAccountData).parsed.info.decimals as number;
    return result;
}

app.post('/solana/transfer/:address', async (req, res) => {
    try {
        let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
        const { to_address, mint, amount } = req.body;
        const privateKey = await GetAddressPrk(req.params.address)
        const signerKeyPair = solanaWeb3.Keypair.fromSecretKey(bs58.decode(privateKey));
        var transaction: solanaWeb3.Transaction
        if (mint == "") {
            transaction = new solanaWeb3.Transaction().add(
                solanaWeb3.SystemProgram.transfer({
                    fromPubkey: signerKeyPair.publicKey,
                    toPubkey: new solanaWeb3.PublicKey(to_address),
                    lamports: Number((amount * 1e9).toFixed(0)),
                }),
            );
        } else {
            let sourceAccount = await getOrCreateAssociatedTokenAccount(
                connection,
                signerKeyPair,
                new solanaWeb3.PublicKey(mint),
                signerKeyPair.publicKey
            );
            let destinationAccount = await getOrCreateAssociatedTokenAccount(
                connection,
                signerKeyPair,
                new solanaWeb3.PublicKey(mint),
                new solanaWeb3.PublicKey(to_address)
            );
            const numberDecimals = await getNumberDecimals(mint);
            transaction = new solanaWeb3.Transaction().add(
                createTransferInstruction(
                    sourceAccount.address,
                    destinationAccount.address,
                    signerKeyPair.publicKey,
                    Number((amount * Math.pow(10, numberDecimals)).toFixed(0)),
                ),
            );
        }
        const signature = await solanaWeb3.sendAndConfirmTransaction(
            connection,
            transaction,
            [signerKeyPair],
        );
        res
            .status(200)
            .send({
                result: signature,
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error
                }
            })
    }
});

app.post('/solana/create-token', async (req, res) => {
    try {
        const rpcUrl = SOLANA_RPC_ENDPOINT
        const { address, name, symbol, uri, amount } = req.body;
        const privateKey = await GetAddressPrk(address)
        const umi = createUmi(rpcUrl); //Replace with your QuickNode RPC Endpoint
        const userWallet = umi.eddsa.createKeypairFromSecretKey(new Uint8Array(bs58.decode(privateKey)));
        const userWalletSigner = createSignerFromKeypair(umi, userWallet);
        const metadata = {
            name: name,
            symbol: symbol,
            uri: uri,
        };
        const mint = generateSigner(umi);
        umi.use(signerIdentity(userWalletSigner));
        umi.use(mplTokenMetadata())
        var createTokenTx = createAndMint(
            umi,
            {
                mint,
                authority: umi.identity,
                name: metadata.name,
                symbol: metadata.symbol,
                uri: metadata.uri,
                sellerFeeBasisPoints: percentAmount(0),
                decimals: 6,
                amount: amount * 1000000,
                tokenOwner: userWallet.publicKey,
                tokenStandard: TokenStandard.Fungible,
            }
        )
        const signerKeyPair = solanaWeb3.Keypair.fromSecretKey(bs58.decode(privateKey));
        createTokenTx = createTokenTx.add(
            [
                {
                    instruction: createSetAuthorityInstruction(
                        mint.publicKey as any,
                        signerKeyPair.publicKey,
                        AuthorityType.MintTokens,
                        null,
                        [],
                        TOKEN_PROGRAM_ID,
                    ) as any,
                    signers: [userWalletSigner]
                },
                {
                    instruction: createSetAuthorityInstruction(
                        mint.publicKey as any,
                        signerKeyPair.publicKey,
                        AuthorityType.FreezeAccount,
                        null,
                        [],
                        TOKEN_PROGRAM_ID,
                    ),
                    signers: [userWalletSigner]
                }
            ] as any
        )
        const signature = await createTokenTx.sendAndConfirm(umi)
        res
            .status(200)
            .send({
                result: {
                    'mint': mint.publicKey,
                    'signature': signature.signature,
                },
            })
    } catch (error) {
        res
            .status(400)
            .send({
                result: null,
                error: {
                    message: error
                }
            })
    }
});
// RAYDIUM
raydiumAPIRouterInit(app)
// NEYNAR
neynarAPIRouterInit(app)

app.listen(8080, () => {
    console.log(`[server]: Server is running at ${8080}`);
});
