import { AgentClient, IAgent, IAgentClient } from "@eternal-ai/core";
import { IConnectFarcasterParams } from "./types";

export * from "./types";

export interface IFarcasterAgentClient extends IAgentClient {
  /** Link Twitter */
  linkFarcaster: (params: IConnectFarcasterParams) => Promise<IAgent>;
  /** Unlink Twitter */
  unlinkFarcaster: (agentId: string) => Promise<IAgent>;
}

export class FarcasterAgentClient
  extends AgentClient
  implements IFarcasterAgentClient
{
  linkFarcaster = async (params: IConnectFarcasterParams): Promise<IAgent> => {
    try {
      await this.api.post("/farcaster/sponsor/register", {
        fid: params.fid,
        uuid: params.uuid,
        assistant_id: params.agent_id,
      });

      return await this.api.post(`/agent/update-farcaster/${params.agent_id}`, {
        farcaster_id: params?.farcaster_id,
        farcaster_username: params?.farcaster_username,
      });
    } catch (e) {
      throw e;
    }
  };

  unlinkFarcaster = async (neynarSignerId: string): Promise<IAgent> => {
    try {
      const res: IAgent = await this.api.delete(
        "/farcaster/sponsor/unregister",
        { data: { id: neynarSignerId } }
      );
      return res;
    } catch (e) {
      throw e;
    }
  };
}
