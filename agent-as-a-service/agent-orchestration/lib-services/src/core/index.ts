import * as dotenv from "dotenv";
import axios from "axios";

dotenv.config();

const coreUrl = process.env.API_CORE_URL

export const StoreAddress = async function (network: string, addr: string, prk: string) {
    try {
        var data = JSON.stringify({
            'WalletId': `solana_${network}`,
            'Address': addr,
            'PrivateKey': prk,
            'WalletType': 20,
        })
        var config = {
            method: 'post',
            url: `${coreUrl}/wallet/update`,
            headers: {
                'Content-Type': 'application/json'
            },
            data: data
        };
        let resp = await axios(config)
        return {
            Address: resp.data.Result.Address,
            SecretVersion: resp.data.Result.SecretVersion,
        }
    } catch (err) {
        throw err
    }
}

export const GetAddressPrk = async function (addr: string) {
    try {
        var data = JSON.stringify({
            'Address': addr,
        })
        var config = {
            method: 'post',
            url: `${coreUrl}/wallet/get-private-key`,
            headers: {
                'Content-Type': 'application/json'
            },
            data: data
        };
        let resp = await axios(config)
        return resp.data.Result as string
    } catch (err) {
        throw err
    }
}