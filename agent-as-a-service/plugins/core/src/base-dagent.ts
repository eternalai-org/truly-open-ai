import { Wallet } from "ethers";
import { IAgentCharacter, IENV, IGetAgentsParams, InitAgent, TokenSetupMode } from "./types";
import { AgentClient } from "./client";
import { dagentLogger } from "./index";

class BaseDagent {
    protected env: IENV;
    protected dagentCharacter: IAgentCharacter;
    protected coreAPI: AgentClient;

    protected signer: Wallet;

    protected accessToken: string;

    constructor({ dagentCharacter, environment }: InitAgent) {
        dagentLogger.log("Environment variables loading...");
        this.env = environment;
        this.dagentCharacter = dagentCharacter;
        this.signer = new Wallet(this.env.PRIVATE_KEY);

        this.coreAPI = new AgentClient({
            endpoint: this.env.ETERNAL_AI_URL,
        });

        this.accessToken = "";
        dagentLogger.log("Environment variables loaded.");
    }

    configAccessToken = async () => {
        try {
            dagentLogger.log("Access token loading...");
            const signerAddress = this.signer.address;

            const message = `Sign this message to get access token: ${signerAddress}`;
            const signature = await this.signer.signMessage(message);

            const accessToken = await this.coreAPI.getAccessToken({
                address: signerAddress,
                signature: signature,
                message: message,
            });

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

    yourAgents = async (params: IGetAgentsParams) => {
        try {
            const agents = await this.coreAPI.list({
                ...params,
                creator: this.signer.address,
            });
            return agents;
        } catch (error) {
            dagentLogger.error("Agent list error:", error);
            throw error;
        }
    };
}

export default BaseDagent;