#!/bin/bash

docker build -f build-arm.Dockerfile --output build .
chmod +x ./build/workersv_arm
cp ./build/workersv_arm ./build/eternal-arm