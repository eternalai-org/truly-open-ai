import { IAgent } from "@eternalai-dagent/core";

const printTableAgents = (agents: IAgent[]) => {
    console.table((agents || []).map(agent => {
        return {
            agent_name: `${agent.agent_name}`,
            id: agent.id,
            topup_evm_address: agent.agent_info.eth_address,
            topup_sol_address: agent.agent_info.sol_address,
        };
    }));
};

export {
    printTableAgents,
}