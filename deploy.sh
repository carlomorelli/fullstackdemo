#!/usr/bin/env bash

# build containers
docker-compose build

# deploy solution
docker-compose down
docker-compose up -d

# run blackbox tests
cd backend
BASEURL=localhost:9000 go test -v -tags blackbox
cd ..

# run unit tests for frontend only when running locally (as they require a browser)
if [ "$1" != "ci" ]; then 
    
    cd frontend
    ./node_modules/.bin/ng test --watch=false

fi
