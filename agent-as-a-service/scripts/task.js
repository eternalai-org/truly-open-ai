import { execSync } from "child_process";
import fs from "fs";
import inquirer from "inquirer";
import path from "path";

const EnumTask = {
  StartDagent: "Start Dagent",
  CleanMonorepo: "Clean",
  BuildAllPackages: "Build Packages",
}

// Function to get all package names from the packages directory
function getPackageNames() {
  const packagesDir = path.join(__dirname, "../packages");
  const packages = fs.readdirSync(packagesDir).filter((folder) => {
    const pkgPath = path.join(packagesDir, folder, "package.json");
    return fs.existsSync(pkgPath);
  });
  return packages;
}

// Function to get all client app names from the client directory
function getClientAppNames() {
  const clientDir = path.join(__dirname, "../client");
  const apps = fs.readdirSync(clientDir).filter((folder) => {
    const appPath = path.join(clientDir, folder, "package.json");
    return fs.existsSync(appPath);
  });
  return apps;
}

async function main() {
  const tasks = {
    [EnumTask.StartDagent]: "node scripts/runDagent.js",
    [EnumTask.BuildAllPackages]: "node scripts/packages-build.js",
    [EnumTask.CleanMonorepo]: "lerna clean && rimraf node_modules packages/**/node_modules packages/**/dist",
  };

  const { selectedTask } = await inquirer.prompt([
    {
      type: "list",
      name: "selectedTask",
      message: "Select the task you want to run:",
      choices: Object.keys(tasks),
    },
  ]);

  switch (selectedTask) {
    default:
      execSync(tasks[selectedTask], { stdio: "inherit" });
      break;
  }
}

main().catch((error) => {
  console.error("❌ Task failed:", error);
  process.exit(1);
});
