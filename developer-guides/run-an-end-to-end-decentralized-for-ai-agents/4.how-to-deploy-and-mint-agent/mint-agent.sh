#!/bin/bash

if [ -z "$1" ]; then
  echo "Usage: $0 <system_prompt_path>"
  exit 1
fi

SYSTEM_PROMPT_PATH=$1

export AGENT_SYSTEM_PROMPT_PATH=$SYSTEM_PROMPT_PATH

npx hardhat run ./scripts/mintAgent.ts --network localhost