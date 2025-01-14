import { spawn } from 'child_process';
import path from "path";
import inquirer from "inquirer";

const __dirname = path.dirname(new URL(import.meta.url).pathname);
const dagentDir = path.join(__dirname, "../dagent");

// Define menu options dynamically
function createMenu(options, dagentDir) {
    // Use Inquirer to prompt the user to select an option
    inquirer
        .prompt([
            {
                type: 'list',
                name: 'selectedOption',
                message: 'Please select an option:',
                choices: options.map(option => option.name)
            }
        ])
        .then(answers => {
            // Get the selected option
            const selectedOption = options.find(option => option.name === answers.selectedOption);
            console.log(`Executing: ${selectedOption.name}`);

            // Run the script for the selected option
            const [command, ...args] = selectedOption.script.split(' ');

            const process = spawn(command, args, {
                cwd: path.resolve(__dirname, dagentDir),  // Set working directory
                stdio: 'pipe',  // Pipe the stdio so we can log the output
            });

            // Log stdout data step by step
            process.stdout.on('data', (data) => {
                console.log(data.toString());
            });

            // Log stderr data if there's an error
            process.stderr.on('data', (data) => {
                console.error(data.toString());
            });

            // Handle process completion
            process.on('close', (code) => {
                if (code === 0) {
                    console.log('Script executed successfully.');
                } else {
                    console.error(`Script failed with exit code ${code}`);
                }
            });
        })
        .catch(error => {
            console.error('Error during selection:', error);
        });
}

export {
    createMenu,
    dagentDir
}