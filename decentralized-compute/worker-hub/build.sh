#!/bin/bash
docker build -f Dockerfile.build --output build .
chmod +x ./build/cli
cp ./build/cli ./build/cli