import { BaseDagent, dagentLogger, IENV, IGetAccessTokenParams } from "@eternalai-dagent/core";
import { getEnvironment } from "../utils/environment";
import { dagentCharacter } from "../dagentCharacter";
import { printTableAgents } from "../utils/helpers";

class BasicAgent extends BaseDagent {
    protected environment: IENV;
    protected authenParams?: IGetAccessTokenParams;

    constructor(params?: IGetAccessTokenParams) {
        const environment = getEnvironment();
        super({
            dagentCharacter: dagentCharacter,
            environment: environment,
        });
        this.authenParams = params || undefined;
        this.environment = environment;

        dagentLogger.info("Code run started...");
    }

    /**
     * Initializes the agent by configuring the access token.
     */
    init = async () => {
        await this.configAccessToken(this.authenParams);
    }

    /**
     * Creates a new agent using the character information and deploys the token if available.
     */
    create = async () => {
        const character = this.dagentCharacter.character;
        const agent = await this.coreAPI.create(character);
        if (this.dagentCharacter.deployToken) {
            await this.deployToken(agent.id);
        }

        await this.ownedAgents();
    }

    /**
     * Retrieves and prints the list of owned agents.
     * Filters agents to include only those with an Ethereum address.
     */
    ownedAgents = async () => {
        const agents = (await this.getAgents({
            limit: 30,
            page: 1,
            creator: this.signerAddress || "",
        }))?.filter(agent => !!agent?.agent_info?.eth_address);
        printTableAgents(agents);
        return agents;
    }

    /**
     * Retrieves and prints the details of an agent by its ID.
     * @param agentId - The ID of the agent to retrieve.
     */
    getAgentById = async (agentId: string) => {
        const agent = await this.getAgent(agentId);
        printTableAgents([agent]);
    }

    /**
     * \[DEPRECATED\] Sets up missions for a Twitter agent.
     * @param agentId - The ID of the agent to set up missions for.
     */
    setupMissionForTwitter = async (agentId: string) => {
        await this.coreAPI.setupMissions({
            agentId: agentId,
            missions: this.dagentCharacter.twitterMissions || [],
        });
    }

    /**
     * Sets up missions for app agent.
     * @param agentId - The ID of the agent to set up missions for.
     */
    setupStoreMission = async (agentId: string) => {
        await this.coreAPI.setupMissionsVer2({
            agentId: agentId,
            missions: this.dagentCharacter?.agentMissions || [],
        })
    }
}

const basicAgent = new BasicAgent();
await basicAgent.init();

// Create a new agent and deploy the token if available in the character configuration.
await basicAgent.create();

// Get your agent information.
await basicAgent.getAgentById("6763d7524ee1600e1122b6f6");