#!/bin/sh
# Based on Deno installer. Copyright 2019 the Deno authors. All rights reserved. MIT license.
set -e

if [ "$OS" = "Windows_NT" ]; then
  target="windows"
  echo "Windows is not supported at the moment, sorry" >&2
  exit 1
else
  case $(uname -sm) in
    "Darwin x86_64")
      target="macos"
      echo "Macos is not supported at the moment, sorry" >&2
      exit 1
      ;;
    "Darwin arm64")
      target="macos"
      echo "Macos is not supported at the moment, sorry" >&2
      exit 1
      ;;
    *)
      target="linux-amd64"
      ;;
  esac
fi

if [ $# -eq 0 ]; then
  SOURCE_URI="https://github.com/eternalai-org/eternal/releases/latest/download/eternal-${target}"
else
  SOURCE_URI="https://github.com/eternalai-org/eternal/releases/download/${1}/eternal-${target}"
fi

echo
echo "> Removing old files"
echo
rm -f eternal
rm -f run-eternal.sh


echo "> Downloading eternal"
echo
curl -LJ -o eternal $SOURCE_URI 

echo "> Setting up permissions"
echo
chmod +x eternal

echo "> Setup complete!"
echo
echo "> To start the miner run:"
echo "  ./eternal -account <private-key> -lighthouse <lighthouse-api-key> -chain <option default:43338>"
echo
# echo "> To start the validator run:"
# echo "  ./eternal -validator -account <private-key> -lighthouse <lighthouse-api-key>"
