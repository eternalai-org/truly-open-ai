import {dagentLogger, IENV, IGetAccessTokenParams} from "@eternal-dagent/core";
import dotenv from "dotenv";
import { dagentCharacter } from "./dagentCharacter";
import { getEnvironment } from "./utils/environment";
import { DagentTwitter } from "@eternal-dagent/client-dagent";
import {createApiRouter, Direct} from "@eternal-dagent/direct";

class AgentTwitter {
  protected environment: IENV;
  protected dagent: DagentTwitter;
  protected authenParams?: IGetAccessTokenParams;

  constructor(params?: IGetAccessTokenParams) {
    dagentLogger.info("Code run started...");
    this.environment = getEnvironment();
    console.log(this.environment);
    this.dagent = new DagentTwitter({
      dagentCharacter: dagentCharacter,
      environment: this.environment,
    });
    this.authenParams = params || undefined;
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

    // twitter_client_id: "TWlYeVpGTXFRdTQ1WW14aDJkNXk6MTpjaQ",
    // twitter_oauth_url: "https://composed-rarely-feline.ngrok-free.app/api/webhook/twitter-oauth"
  };

  setupMissions = async (agentId: string) => {
    await this.dagent.setupMissions(agentId);
  };

  createAndRunDagent = async () => {
    // const privateKey = this.environment.PRIVATE_KEY;
    //
    // if (!privateKey && !this.authenParams) {
    //   dagentLogger.error("Please provide private key or params token.");
    // }
    // await this.dagent.init(this.authenParams);
    // // await this.dagent.setupMissions("6763d7524ee1600e1122b6f6");
    // const agent = await this.createDagent();
    //
    // if (agent?.id) {
    //   await this.linkDagentToTwitter(agent.id);
    // }

    const direct = new Direct({
      routers: [createApiRouter()]
    })

    direct.start(80);

    await this.getCreatedAgents();
  };

  runDagent = async (agentId: string) => {
    const privateKey = this.environment.PRIVATE_KEY;
    if (!privateKey && !this.authenParams) {
      dagentLogger.error("Please provide private key or params token.");
    }
    await this.dagent.init(this.authenParams);
    const agent = await this.dagent.getAgent(agentId);
    await this.dagent.setupMissions(agent.id);
    console.table(([agent] || []).map(agent => {
      return {
        agent_name: `${agent.agent_name}`,
        id: agent.id,
        topup_evm_address: agent.agent_info.eth_address,
        topup_sol_address: agent.agent_info.sol_address,
        status: agent.status,
      };
    }));
  };

}

dotenv.config();

// Case not set PRIVATE_KEY
// const agentTwitter = new AgentTwitter({
//   address: "0xadee4cba41ebb02ad43b76776dd939ce7f2a0c7d",
//   message: "Eternal Dagent",
//   signature: "0x13cd10b3bd1e6c9deaf0dec085a386c3019fdd21511284dcc084684ddf908edf390b71f7d544b3122971f1e52ce9609de5cf502f73d4e359e57b23ead6ed117c1b"
// });

// Case set PRIVATE_KEY
const agentTwitter = new AgentTwitter();
agentTwitter.createAndRunDagent()
    .then(() => {
      dagentLogger.info("Code run completed...");
    });

// // Case agent created and run
// agentTwitter.runDagent("6763d7524ee1600e1122b6f6")
//     .then(() => {
//       dagentLogger.info("Code run completed...");
//     });