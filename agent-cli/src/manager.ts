import * as fs from 'fs';
import path from 'path';
import { logError, logInfo } from './log';


interface Agent {
    AgentID: string
    Name: string
    Network: string
    ChainID: string
    Model: string
}

const FilePath = "./agents/list.json";


const insertAgent = async (agent: Agent) => {
    try {
        logInfo(`FilePath: ${FilePath}`);

        if (!fs.existsSync(FilePath)) {
            fs.writeFileSync(FilePath, JSON.stringify([agent], null, 2))

        } else {
            const data = fs.readFileSync(FilePath, 'utf8');
            const currentAgents: Agent[] = JSON.parse(data);
            currentAgents.push(agent);
            fs.writeFileSync(FilePath, JSON.stringify(currentAgents, null, 2))
        }
    } catch (e) {
        logError(`Save agent info ${e}`);
    }
}

const getAgents = async (): Promise<Agent[]> => {
    const data = fs.readFileSync(FilePath, 'utf8');
    const currentAgents: Agent[] = JSON.parse(data);
    return currentAgents;

}

export {
    Agent,
    insertAgent,
    getAgents
}