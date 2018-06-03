#!/usr/bin/env bash

# build backend container (runs unit tests as well)
docker build -t backend-image ./backend

# deploy backend
docker stop backend || true
docker rm backend || true
docker run -d -p 8080:9000 --name backend backend-image

# run blackbox tests
cd backend
BASEURL=localhost:8080 go test -v -tags blackbox
