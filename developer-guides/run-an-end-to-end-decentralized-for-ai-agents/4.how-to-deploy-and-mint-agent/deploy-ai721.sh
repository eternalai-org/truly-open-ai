#!/bin/bash
echo "Installing dependencies..."
npm install &&
npx hardhat run ./scripts/autoDeploy.ts --network localhost 