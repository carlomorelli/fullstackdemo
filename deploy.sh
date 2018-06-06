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

# do frontend tests and e2e test only when running locally
if [ "$1" != "ci" ]; then 
    
    cd frontend
    ./node_modules/.bin/ng test --watch=false

    ./node_modules/.bin/webdriver-manager update
    ./node_modules/.bin/webdriver-manager start --detach
    ./node_modules/.bin/protractor e2e/protractor.conf.js
    cd ..

fi
