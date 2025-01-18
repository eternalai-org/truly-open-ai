#! /bin/bash

docker-compose -f milvus-docker-compose.yml rm -f
sudo rm -rf ./volumes
docker-compose -f milvus-docker-compose.yml up