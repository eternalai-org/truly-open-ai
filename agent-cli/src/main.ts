import { Command } from 'commander';
import { Framework, Model, Network } from './const';
import { createAgent, listAgents } from './handler';

// Initialize the commander object
const program = new Command();

// Define the CLI command and its options
program
    .name('eai')
    .description('You can deploy and manage your decentralized agents flexibly and simply - on any chain, any framework.')
    .version('1.0.0');


// Define a command with parameters
const agentCmd = program
    .command('agent')
    .description('Deploy and manage agents');


// Define a command with parameters
agentCmd
    .command('create')
    // .description('Deploy and manage agents')
    // .command('create')
    .description('Create a new agent')
    .option('-c, --chain <chain>', 'The blockchain which the new agent will be deployed on.', Network.Base)
    .option('-f, --framework <framework>', 'The framework is used for new agent.', Framework.Eliza)
    .option('-m, --model <model>', 'The agent\'s model.')
    .requiredOption('-p, --path <path>', 'The path of the agent\'s character file.')
    // .option('-c, --chain', 'The blockchain which the new agent will be deployed on.', Network.Base)
    .option('-n, --name <name>', 'The agent\'s name.')
    .action((options) => {
        // const greeting = `Hello, ${name}!`;
        createAgent(options);

        // if (options.time) {
        //     console.log(greeting, `The current time is: ${new Date().toLocaleTimeString()}`);
        // } else {
        //     console.log(greeting);
        // }
    })

agentCmd
    .command('ls')
    // .description('Deploy and manage agents')
    // .command('ls')
    .description('See list of agents')
    .action(() => {
        // const greeting = `Hello, ${name}!`;
        listAgents();

        // if (options.time) {
        //     console.log(greeting, `The current time is: ${new Date().toLocaleTimeString()}`);
        // } else {
        //     console.log(greeting);
        // }
    });



// Parse the command-line arguments
program.parse(process.argv);
