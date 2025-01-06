#!/bin/bash

docker build -f build.Dockerfile --output build .
chmod +x ./build/workersv
cp ./build/workersv ./build/eternal-linux-amd64