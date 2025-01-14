// Define the agent options with commands (can be expanded/reused)
import { createMenu } from "./comma.js";

const agentOptions = [
    {
        name: 'Run Twitter',
        script: `yarn run:twitter`
    },
    {
        name: 'Run Farcaster',
        script: `yarn run:farcaster`
    },
];

// Create and show the menu
createMenu(agentOptions, "../dagent");