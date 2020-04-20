#!/bin/bash

docker rm -f $(docker ps -a -q)

docker volume prune

docker network prune

docker rmi $(docker image ls -q)

rm -r ./api-src/wallet
