#!/bin/bash

# Navigate to the project directory
cd /Users/macbook_autonomous/projects/new-bitcoin-city/eternal/sol || exit

# Run npm install and then the Hardhat script
npm install && npx hardhat run /Users/macbook_autonomous/projects/new-bitcoin-city/eternal/sol/scripts/auto_deploy.ts --network localhost