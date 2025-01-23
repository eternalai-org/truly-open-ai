#!/bin/bash
current_dir=$(pwd)
worker_hub_folder="$current_dir/decentralized-compute/worker-hub"

# Check if at least one argument is provided
if [ $# -lt 1 ]; then
    echo "Usage:"
    echo "- eai miner setup"
    exit 1
fi

# Handle the argument using a case statement
case "$1" in
    #eai miner
    "miner")
      case "$2" in
          "setup")
            #eai miner setup
            cd $worker_hub_folder && make start_cli
          ;;
          *)
              echo "Invalid option: $2"
              exit 1
              ;;
      esac

    ;;
    *)
        echo "Invalid option: $1"
        exit 1
        ;;
esac