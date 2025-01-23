#!/bin/bash
current_dir=$(pwd)
worker_hub_folder="$current_dir/decentralized-compute/worker-hub" #step 1: eai miner setup
aaas_folder="$current_dir/agent-as-a-service/agent-orchestration/backend" #step 2: aaas start
aaas_folder="$current_dir/agent-as-a-service/agent-orchestration/backend" #step 3: eai aaas deploy-contract
ai721_folder="$current_dir/developer-guides/run-an-end-to-end-decentralized-for-ai-agents/4.how-to-deploy-and-mint-agent" #step 3: eai aaas deploy-contract

# Check if at least one argument is provided
if [ $# -lt 1 ]; then
    echo "Usage:"
    echo "- eai miner setup"
    echo "- eai aaas start"
    echo "- eai aaas deploy-contract"
    echo "- eai apis start"
    echo "- eai agent chat"
    exit 1
fi

check_command() {
    command -v "$1" >/dev/null 2>&1
}

# Function to check for docker-compose or docker compose
check_docker() {
    if check_command "docker-compose"; then
        echo "docker-compose"
    elif check_command "docker"; then
        # Check if 'docker compose' is available (as a subcommand)
        if docker compose version >/dev/null 2>&1; then
            echo "docker compose"
        else
            #docker is installed, but 'docker compose' is not available.
            echo "-1"
        fi
    else
        #Neither docker-compose nor docker is installed.
        echo "-2"
    fi
}

# Function to handle miner commands
handle_miner_commands() {
    case "$1" in
        "setup")
            echo "Setting up local chain and miners..."
            cd "$worker_hub_folder" && make start_cli
            ;;
        *)
            echo "Invalid option: $1 for miner"
            exit 1
            ;;
    esac
}

handle_aaas_commands() {
    case "$1" in
        "deploy-contract")
           echo "aaas deploy-contract"
           cd $ai721_folder && ./deploy-ai721.sh
        ;;
        "start")
            echo "eai aaas start"
            docker_compose=$(check_docker)
            cd $aaas_folder && $docker_compose build && $docker_compose up
            ;;
        *)
        echo "Invalid option: $1 for miner"
        exit 1
        ;;
    esac
}

handle_agent_commands() {
    case "$1" in
        "chat")
          echo "agent chat"
        ;;
        *)
        echo "Invalid option: $1 for miner"
        exit 1
        ;;
    esac
}

handle_api_commands() {
    case "$1" in
        "start")
          echo "apis start"
        ;;
        *)
        echo "Invalid option: $1 for miner"
        exit 1
        ;;
    esac
}

# Handle the main command using a case statement
case "$1" in
    "miner")
        handle_miner_commands "$2"
        ;;
    "aaas")
        handle_aaas_commands "$2"
        ;;
    "apis")
        handle_api_commands "$2"
        ;;
    "agent")
        handle_agent_commands "$2"
        ;;
    *)
        echo "Invalid option: $1"
        exit 1
        ;;
esac