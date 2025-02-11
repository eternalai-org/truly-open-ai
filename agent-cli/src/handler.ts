import { exec } from 'child_process';
import { ChainIDMap, Framework, Network, NetworkConfig, ETERNALAI_URL } from "./const";

// for dev
import dotenv from 'dotenv';
import { getSupportedModels } from "./utils";
import { mintAgent } from "./mint";
import { Agent, getAgents, insertAgent } from "./manager";
import { logTable } from "./log";
dotenv.config();

const validateParams = async ({
    chain,
    framework,
    model,
    path,
    name,

}: {
    chain: Network,
    framework: string,
    model: string,
    path: string,
    name: string,
}): Promise<{
    model: string
}> => {
    const chainID = ChainIDMap[chain];
    const supportedModels = await getSupportedModels(chainID);
    if (model) {
        if (!supportedModels[model]) {
            throw new Error(`model ${model} is not supported in chain ${chain}`);
        }
    } else {
        const models = Object.keys(supportedModels);
        if (models.length == 0) {
            throw new Error(`no models supported in chain ${chain}`);
        }
        model = models[0];
    }

    return { model };
}

const createAgent = async ({
    chain,
    framework,
    model,
    path,
    name,

}: {
    chain: Network,
    framework: string,
    model: string,
    path: string,
    name: string,
}) => {
    // console.log("options: ", options);

    const newParam = await validateParams({
        chain,
        framework,
        model,
        path,
        name,
    });
    model = newParam.model;

    console.log("newParam: ", newParam);

    // get info from chain network
    // const chainRPC = "";
    const chainID = ChainIDMap[chain];
    const networkInfo = NetworkConfig[chainID];
    if (!networkInfo || !networkInfo.agentContractAddress || !networkInfo.url) {
        throw new Error("invalid chain argument")
    }

    const ETERNALAI_RPC_URL = networkInfo.url;
    const ETERNALAI_AGENT_CONTRACT_ADDRESS = networkInfo.agentContractAddress;


    // call contract to mint agent id
    const agentID = await mintAgent({
        rpcURL: ETERNALAI_RPC_URL,
        privKey: process.env.PRIVATE_KEY || "",
        agentSystemPrompPath: path,
        agentContractAddress: ETERNALAI_AGENT_CONTRACT_ADDRESS,
        promptSchedulerAddress: networkInfo.promptSchedulerAddress,
        gpuManagerAddress: networkInfo.gpuManagerAddress
    });

    if (!agentID) {
        console.error("Mint agent failed! Please retry.");
        return;
    }

    console.log("Mint agent successfully: Agent ID: ", agentID);

    const ETERNALAI_API_KEY = process.env.ETERNALAI_API_KEY;
    const TWITTER_USERNAME = process.env.TWITTER_USERNAME;
    const TWITTER_PASSWORD = process.env.TWITTER_PASSWORD;
    const TWITTER_EMAIL = process.env.TWITTER_EMAIL;
    const TWITTER_TARGET_USERS = process.env.TWITTER_TARGET_USERS;


    // const randomId = randomID();
    const AGENT_UID = `${chain}_${agentID}`;
    const agentName = name || AGENT_UID;

    switch (framework) {
        case Framework.Eliza: {
            // Path to your Bash script
            const scriptPath = `sh src/eliza/start.sh ${AGENT_UID} ${ETERNALAI_URL} ${ETERNALAI_API_KEY} ${chainID} ${ETERNALAI_RPC_URL} ${ETERNALAI_AGENT_CONTRACT_ADDRESS} ${agentID} ${model} ${TWITTER_USERNAME} ${TWITTER_PASSWORD} ${TWITTER_EMAIL} ${TWITTER_TARGET_USERS} ${agentName}`;

            // Run the Bash script
            exec(scriptPath, (error: any, stdout: any, stderr: any) => {
                if (error) {
                    console.error(`Error executing script: ${error.message}`);
                    // throw error;
                }
                if (stderr) {
                    console.error(`stderr: ${stderr}`);
                }
                console.log(`stdout: ${stdout}`);
            });

        }
    }
    const newAgent: Agent = {
        Name: agentName,
        AgentID: agentID.toString(),
        Network: chain,
        ChainID: chainID,
        Model: model,
    }

    // insert into the file to manage agents
    await insertAgent(newAgent);
}

const listAgents = async () => {
    const agents = await getAgents();
    logTable(agents);

}

export {
    createAgent,
    listAgents
}