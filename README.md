# Fullstack Web Demo

Check the build status here:
[![Build Status](https://travis-ci.org/carlomorelli/fullstackdemo.svg?branch=master)](https://travis-ci.org/carlomorelli/fullstackdemo)


## Requirements
* Golang 1.9 or later
* NodeJS 8.9 or later
* Docker 18.x or later
* docker-compose 1.17.x or later
* Chrome version 59 or later

## Quick setup (Linux only)
After successfully installing the requirements on the system, run
```
./deploy.sh
```
This will trigger:
* Unit tests and integration tests for backend
* Unit tests for frontend
* Build of docker images for backend and frontend
* Instantiation of the full solution (frontend reachable at http://localhost:4200)

## Component and e2e testing
Once the system is deployed is possible to run blackbox component tests for the backend using `go test`, and e2e tests for the whole solution using Protractor.

To run test the blackbox tests:
```
cd backend
BASEURL=localhost:9000 go test -v -tags blackbox
```

To run Protractor:
```
cd frontend
./node_modules/.bin/webdriver-manager update
./node_modules/.bin/webdriver-manager start
./node_modules/.bin/protractor e2e/protractor.conf.js
```
Notice that Webdriver and Protractor should be launched on separate consoles.

## CI setup
For CI launch
```
./deploy.sh ci
```
This will run a subset of test avoiding the frontend unit test suite (the backend unit test suite is run during the Docker container generation), as Angular Jasmine requires a live browser (which is problematic on some CI servers).
As alternative, simply manage the containers with
```
docker-compose up -d  #create deployment
docker-compose down   #shutdown deployment
```

