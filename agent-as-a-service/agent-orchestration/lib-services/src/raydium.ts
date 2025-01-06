import * as solanaWeb3 from "@solana/web3.js";
import { decodeBase58, ethers, } from "ethers";
import { GetAddressPrk } from "./core";
import { Raydium, TxVersion, CREATE_CPMM_POOL_PROGRAM, CREATE_CPMM_POOL_FEE_ACC, parseTokenAccountResp, API_URLS } from '@raydium-io/raydium-sdk-v2'
import BN from 'bn.js'
import bs58 from "bs58";
import * as core from "express-serve-static-core";
import {
    TOKEN_PROGRAM_ID,
    TOKEN_2022_PROGRAM_ID,
    NATIVE_MINT,
} from '@solana/spl-token';
import axios from "axios";
import { cloneAndSnakeCaseFields } from "./utils";
import { SOLANA_RPC_ENDPOINT } from "./contants";

export const raydiumAPIRouterInit = (app: core.Express) => {

    interface SwapCompute {
        id: string
        success: true
        version: 'V0' | 'V1'
        openTime?: undefined
        msg: undefined
        data: {
            swapType: 'BaseIn' | 'BaseOut'
            inputMint: string
            inputAmount: string
            outputMint: string
            outputAmount: string
            otherAmountThreshold: string
            slippageBps: number
            priceImpactPct: number
            routePlan: {
                poolId: string
                inputMint: string
                outputMint: string
                feeMint: string
                feeRate: number
                feeAmount: string
            }[]
        }
    }

    const fetchTokenAccountData = async (owner: solanaWeb3.Keypair) => {
        let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
        const solAccountResp = await connection.getAccountInfo(owner.publicKey)
        const tokenAccountResp = await connection.getTokenAccountsByOwner(owner.publicKey, { programId: TOKEN_PROGRAM_ID })
        const token2022Req = await connection.getTokenAccountsByOwner(owner.publicKey, { programId: TOKEN_2022_PROGRAM_ID })
        const tokenAccountData = parseTokenAccountResp({
            owner: owner.publicKey,
            solAccountResp,
            tokenAccountResp: {
                context: tokenAccountResp.context,
                value: [...tokenAccountResp.value, ...token2022Req.value],
            },
        })
        return tokenAccountData
    }

    app.post('/solana/compute-raydium', async (req, res) => {
        try {
            const { input_mint, output_mint, slippage, amount } = req.body;
            const txVersion: string = 'V0' // or LEGACY
            const { data: swapResponse } = await axios.get<SwapCompute>(
                `${API_URLS.SWAP_HOST
                }/compute/swap-base-in?inputMint=${input_mint}&outputMint=${output_mint}&amount=${amount}&slippageBps=${slippage * 100}&txVersion=${txVersion}`
            )
            res
                .status(200)
                .send({
                    result: cloneAndSnakeCaseFields(swapResponse['data']),
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

    app.post('/solana/trade-raydium', async (req, res) => {
        try {
            // let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT, "confirmed");
            let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
            const { address, input_mint, output_mint, slippage, amount } = req.body;
            const privateKey = await GetAddressPrk(address)
            const signerKeyPair = solanaWeb3.Keypair.fromSecretKey(bs58.decode(privateKey));

            const { tokenAccounts } = await fetchTokenAccountData(signerKeyPair)
            const [isInputSol, isOutputSol] = [input_mint === NATIVE_MINT.toBase58(), output_mint === NATIVE_MINT.toBase58()]

            const inputTokenAcc = tokenAccounts.find((a) => a.mint.toBase58() === input_mint)?.publicKey
            const outputTokenAcc = tokenAccounts.find((a) => a.mint.toBase58() === output_mint)?.publicKey

            if (!inputTokenAcc && !isInputSol) {
                throw new Error('do not have input token account')
            }

            const { data } = await axios.get<{
                id: string
                success: boolean
                data: { default: { vh: number; h: number; m: number } }
            }>(`${API_URLS.BASE_HOST}${API_URLS.PRIORITY_FEE}`)
            const txVersion: string = 'V0' // or LEGACY
            const isV0Tx = txVersion === 'V0'
            const { data: swapResponse } = await axios.get<SwapCompute>(
                `${API_URLS.SWAP_HOST
                }/compute/swap-base-in?inputMint=${input_mint}&outputMint=${output_mint}&amount=${amount}&slippageBps=${slippage * 100}&txVersion=${txVersion}`
            )

            const { data: swapTransactions } = await axios.post<{
                id: string
                version: string
                success: boolean
                data: { transaction: string }[]
            }>(`${API_URLS.SWAP_HOST}/transaction/swap-base-in`, {
                computeUnitPriceMicroLamports: String(data.data.default.h),
                swapResponse,
                txVersion,
                wallet: signerKeyPair.publicKey.toBase58(),
                wrapSol: isInputSol,
                unwrapSol: isOutputSol, // true means output mint receive sol, false means output mint received wsol
                inputAccount: isInputSol ? undefined : inputTokenAcc?.toBase58(),
                outputAccount: isOutputSol ? undefined : outputTokenAcc?.toBase58(),
            })

            const allTxBuf = swapTransactions.data.map((tx) => Buffer.from(tx.transaction, 'base64'))
            const allTransactions = allTxBuf.map((txBuf) =>
                isV0Tx ? solanaWeb3.VersionedTransaction.deserialize(txBuf) : solanaWeb3.Transaction.from(txBuf)
            )
            const signatures: string[] = []
            if (!isV0Tx) {
                for (const tx of allTransactions) {
                    const transaction = tx as solanaWeb3.Transaction
                    transaction.sign(signerKeyPair)
                    const smRsp = (await connection.simulateTransaction(transaction)).value as solanaWeb3.SimulatedTransactionResponse
                    if (smRsp.err != null) {
                        throw new Error(smRsp.err.toString())
                    }
                    const signature = await connection.sendTransaction(transaction, [signerKeyPair], { skipPreflight: true })
                    signatures.push(signature)
                }
            } else {
                for (const tx of allTransactions) {
                    const transaction = tx as solanaWeb3.VersionedTransaction
                    transaction.sign([signerKeyPair])
                    const smRsp = (await connection.simulateTransaction(transaction)).value as solanaWeb3.SimulatedTransactionResponse
                    if (smRsp.err != null) {
                        throw new Error(smRsp.err.toString())
                    }
                    const signature = await connection.sendTransaction(tx as solanaWeb3.VersionedTransaction, { skipPreflight: true })
                    const { lastValidBlockHeight, blockhash } = await connection.getLatestBlockhash({
                        commitment: 'confirmed',
                    })
                    await connection.confirmTransaction(
                        {
                            blockhash,
                            lastValidBlockHeight,
                            signature: signature,
                        },
                        'confirmed'
                    )
                    signatures.push(signature)
                }
            }
            res
                .status(200)
                .send({
                    result: {
                        'output_amount': swapResponse.data.outputAmount,
                        'signatures': signatures
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

    app.post('/solana/raydium-create-cpmm-pool', async (req, res) => {
        try {
            // let connection = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
            let connection: any = new solanaWeb3.Connection(SOLANA_RPC_ENDPOINT);
            const privateKey = await GetAddressPrk(req.body.address)
            var owner: any = solanaWeb3.Keypair.fromSecretKey(ethers.toBeArray(decodeBase58(privateKey)));
            const raydium = await Raydium.load({
                owner,
                connection,
                cluster: 'mainnet',
                disableFeatureCheck: true,
                disableLoadToken: false,
            })
            const mintA = await raydium.token.getTokenInfo(req.body.mint_a)
            const mintB = await raydium.token.getTokenInfo(req.body.mint_b)
            const txVersion = TxVersion.V0 // or TxVersion.LEGACY
            const feeConfigs = await raydium.api.getCpmmConfigs()
            const { execute, extInfo } = await raydium.cpmm.createPool({
                programId: CREATE_CPMM_POOL_PROGRAM,
                poolFeeAccount: CREATE_CPMM_POOL_FEE_ACC,
                mintA,
                mintB,
                mintAAmount: new BN(req.body.amount_a),
                mintBAmount: new BN(req.body.amount_b),
                startTime: new BN(0),
                feeConfig: feeConfigs[0],
                associatedOnly: false,
                ownerInfo: {
                    useSOLBalance: true,
                },
                txVersion,
            })
            const { txId } = await execute()
            res
                .status(200)
                .send({
                    result: {
                        'signature': txId,
                        'extInfo': extInfo,
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
}