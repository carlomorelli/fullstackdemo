# Fullstack Web Demo

## Requirements
* 
* Golang 1.9 or later
* NodeJS 8.9 or later
* Docker and docker-compose 18.x or later
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
* Run of Protractor E2E tests on the frontend (1)

(1): WARNING On some systems the WebDriver server may not detach correctly. The solution is to comment this line in the `deploy.sh` script:
```
./node_modules/.bin/webdriver-manager update
./node_modules/.bin/webdriver-manager start --detach
./node_modules/.bin/protractor e2e/protractor.conf.js
```
and launch Webdriver and Protractor manually on separate consoles.

## CI setup
For CI launch
```
./deploy.sh ci
```
This will run a subset of test avoiding the Protractor and Angular Jasmine ones that require a live browser (which is problematic on some CI servers).
As alternative, simply manage the containers with
```
docker-compose up -d  #create deployment
docker-compose down   #shutdown deployment
```

