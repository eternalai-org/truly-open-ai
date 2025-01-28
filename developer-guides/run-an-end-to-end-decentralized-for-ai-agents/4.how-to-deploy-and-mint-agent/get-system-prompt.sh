#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <agent_id>"
  exit 1
fi

AGENT_ID=$1

export AGENT_ID=$AGENT_ID

npx hardhat run ./scripts/getSystemPrompt.ts --network localhost