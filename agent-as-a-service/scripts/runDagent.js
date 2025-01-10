import { execSync } from "child_process";
import path from "path";
import inquirer from "inquirer";

const __dirname = path.dirname(new URL(import.meta.url).pathname);
const dagentDir = path.join(__dirname, "../dagent");

const promptUser = async () => {
    const answers = await inquirer.prompt([
        {
            type: "list",
            name: "script",
            message: "Which script would you like to run?",
            choices: ["run:twitter", "run:farcaster"],
        },
    ]);
    runScript(answers.script);
};

const runScript = (scriptName) => {
    try {
        const result = execSync(`yarn ${scriptName}`, { cwd: dagentDir, encoding: "utf-8" });
        console.log(result);  // Print the output
    } catch (err) {
        console.error(`Error executing the script: ${err.message}`);
        console.error(`stderr: ${err.stderr}`);
    }
};

promptUser();
