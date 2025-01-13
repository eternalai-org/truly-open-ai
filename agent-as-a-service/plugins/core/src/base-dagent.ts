import { Wallet } from "ethers";
import {
    IAgentCharacter,
    IENV,
    IGetAccessTokenParams,
    IGetAgentsParams,
    InitAgent,
    TokenSetupMode
} from "./types";
import { AgentClient } from "./client";
import { dagentLogger } from "./index";

class BaseDagent {
    protected env: IENV;
    protected dagentCharacter: IAgentCharacter;
    protected coreAPI: AgentClient;

    protected signer?: Wallet;
    protected signerAddress: string;
    protected accessToken: string;

    constructor({ dagentCharacter, environment }: InitAgent) {
        dagentLogger.log("Environment variables loading...");
        this.env = environment;
        this.dagentCharacter = dagentCharacter;
        if (this.env.PRIVATE_KEY) {
            this.signer = new Wallet(this.env.PRIVATE_KEY);
            this.signerAddress = this.signer ? this.signer.address : "";
        } else {
            this.signerAddress = "";
        }

        this.coreAPI = new AgentClient({
            endpoint: this.env.ETERNAL_AI_URL,
        });


        this.accessToken = "";
        dagentLogger.log("Environment variables loaded.");
    }

    configAccessToken = async (params?: IGetAccessTokenParams) => {
        try {
            dagentLogger.log("Access token loading...");
            let _params: IGetAccessTokenParams | undefined = params || undefined;

            if (!!params && !params?.address || !params?.signature || !params?.message) {
                throw new Error("Please provide address, signature and message.");
            }

            if (!params && !this.signer) {
                throw new Error("Please set signer or provide address, signature and message.");
            }

            if (!params && this.signer) {
                const signerAddress = this.signer.address;

                const message = `Sign this message to get access token: ${signerAddress}`;
                const signature = await this.signer.signMessage(message);

                _params = {
                    address: signerAddress,
                    signature: signature,
                    message: message,
                };
            }

            const accessToken = await this.coreAPI.getAccessToken({
                address: _params?.address || "",
                signature: _params?.signature || "",
                message: _params?.message || "",
            });

            this.signerAddress = _params?.address || "";

            this.coreAPI.setAuthToken(accessToken);
            this.accessToken = accessToken;
            dagentLogger.log("Access token loaded.");
            return accessToken;
        } catch (error) {
            dagentLogger.error("Access token error:", error);
            throw error;
        }
    };


    createAgent = async () => {
        try {
            dagentLogger.log("Creating agent...");
            const agent = await this.coreAPI.create(this.dagentCharacter.character);
            dagentLogger.log(`Agent created: ${agent.id}`);

            if (this.dagentCharacter.deployToken) {
                await this.deployToken(agent.id);
            }

            return agent;
        } catch (error) {
            dagentLogger.error("Create agent error:", error);
        }
    };

    deployToken = async (agentId: string) => {
        try {
            dagentLogger.log("Deploying token...");
            if (!this.dagentCharacter.deployToken) {
                throw new Error("Please set deploy token in dagentCharacter.");
            } else if (this.dagentCharacter.deployToken.create_token_mode === TokenSetupMode.AUTO_CREATE_RUNE && !this.dagentCharacter.deployToken.ticker) {
                throw new Error("Please set ticker in deploy token.");
            }
            const token = await this.coreAPI.deployToken({
                ...this.dagentCharacter.deployToken!,
                agent_id: agentId,
            });
            dagentLogger.log(`Token deployed: ${token.token_address}`, {
                token_address: token.token_address,
                agent_id: agentId,
                agent_name: token.agent_name,
            });
        } catch (error) {
            dagentLogger.error("Deploy token error:", error);
            throw error;
        }
    };

    getAgents = async (params: IGetAgentsParams) => {
        try {
            const agents = await this.coreAPI.list(params);
            return agents;
        } catch (error) {
            dagentLogger.error("Agent list error:", error);
            throw error;
        }
    };

    getSignerAddress = () => {
        return this.signerAddress || this.signer?.address || "";
    }

    yourAgents = async (params: IGetAgentsParams) => {
        try {
            const agents = await this.coreAPI.list({
                ...params,
                creator: this.getSignerAddress(),
            });
            return agents;
        } catch (error) {
            dagentLogger.error("Agent list error:", error);
            throw error;
        }
    };

    getAgent = async (agentId: string) => {
        try {
            const agent = await this.coreAPI.detail(agentId);
            return agent;
        } catch (error) {
            dagentLogger.error("Agent get error:", error);
            throw error;
        }
    }
}

export default BaseDagent;