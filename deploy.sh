#!/usr/bin/env bash

# build backend container
docker build -t backend-image ./backend

# deploy backend
docker stop backend || true
docker rm backend || true
docker run -d -p 9000:9000 --name backend backend-image

# run blackbox tests
cd backend
BASEURL=localhost:9000 go test -v -tags blackbox
cd ..

# build frontend container
docker build -t frontend-image ./frontend

# deploy frontend
docker stop frontend || true
docker rm frontend || true
docker run -d -p 4200:4200 --name frontend frontend-image

# run jasmine tests
cd frontend
#TODO
cd ..
