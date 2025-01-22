#!/bin/bash
# Step 1: Install Dependencies
echo "Installing dependencies..."
npm install &&
npx hardhat run ./scripts/autoDeploy.ts --network localhost && cd ..