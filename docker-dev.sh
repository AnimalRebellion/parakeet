#!/bin/bash
name="$(docker ps -a -f ancestor=parakeet-test -q)"
echo $name
docker stop $name
docker rm $name
docker build -t parakeet-test -f Dockerfile .
docker run -p 8080:8080 --network=nats_default --env-file=env.list -t parakeet-test
