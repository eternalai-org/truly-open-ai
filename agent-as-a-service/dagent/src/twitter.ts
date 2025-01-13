import { dagentLogger, IENV } from "@eternal-dagent/core";
import dotenv from "dotenv";
import { dagentCharacter } from "./dagentCharacter";
import { getEnvironment } from "./utils/environment";
import { DagentTwitter } from "@eternal-dagent/client-dagent";

class AgentTwitter {
  protected environment: IENV;
  protected dagent: DagentTwitter;

  constructor() {
    dagentLogger.info("Code run started...");
    this.environment = getEnvironment();
    console.log(this.environment);
    this.dagent = new DagentTwitter({
      dagentCharacter: dagentCharacter,
      environment: this.environment,
    });
  }

  createDagent = async () => {
    return await this.dagent.createAgent();
  };

  getCreatedAgents = async () => {
    const agents = (await this.dagent.yourAgents({
      limit: 30,
      page: 1,
    }))?.filter(agent => !!agent?.agent_info?.eth_address);
    console.table((agents || []).map(agent => {
      return {
        agent_name: `${agent.agent_name}`,
        id: agent.id,
        topup_evm_address: agent.agent_info.eth_address,
        topup_sol_address: agent.agent_info.sol_address,
      };
    }));
    return agents;
  };

  linkDagentToTwitter = async (agentId: string) => {
    // Link twitter
    await this.dagent.linkTwitter({
      agent_id: agentId!,
      callback_url: "https://eternalai.org",
      twitter_client_id: this.environment.TWITTER.CLIENT_ID,
      twitter_oauth_url: "https://imagine-backend.bvm.network/api/webhook/twitter-oauth"
    });
  };

  setupMissions = async (agentId: string) => {
    await this.dagent.setupMissions(agentId);
  };

  createAndRunDagent = async () => {
    await this.dagent.init();
    // await this.dagent.setupMissions("6763d7524ee1600e1122b6f6");
    const agent = await this.createDagent();

    if (agent?.id) {
      await this.linkDagentToTwitter(agent.id);
    }

    await this.getCreatedAgents();
  };

  runDagent = async (agentId: string) => {
    await this.dagent.init();
    const agent = await this.dagent.getAgent(agentId);
    await this.dagent.setupMissions(agent.id);
    console.table(([agent] || []).map(agent => {
      return {
        agent_name: `${agent.agent_name}`,
        id: agent.id,
        topup_evm_address: agent.agent_info.eth_address,
        topup_sol_address: agent.agent_info.sol_address,
      };
    }));
  };

}

dotenv.config();
const agentTwitter = new AgentTwitter();
agentTwitter.createAndRunDagent()
    .then(() => {
      dagentLogger.info("Code run completed...");
    });

// agentTwitter.runDagent("6763d7524ee1600e1122b6f6")
//     .then(() => {
//       dagentLogger.info("Code run completed...");
//     });