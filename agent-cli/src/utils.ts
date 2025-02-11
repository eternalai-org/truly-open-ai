import { ETERNALAI_URL } from "./const";
import axios from 'axios';



const getSupportedModels = async (chainID: string) => {
    const url = "https://api.eternalai.org/api/chain-config/get";
    const response = await axios.get(url, { params: { chain_id: chainID } });
    console.log('Filtered API Data:', response.data);

    if (response.data.status != 1) {
        throw new Error("get supported models status invalid");
    }

    return response.data.data?.support_model_names;
}

export {
    getSupportedModels
}
