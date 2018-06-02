# Test plan

## Frontend component tests
login fields empty -> expect error popup
login fields malformed -> expect error popup
login fields valid -> redirect to homepage
homepage -> shows name, surname
login page layout -> horizontal design on browser, vertical on mobile

## Backend component tests
GET /ping -> returns simple keepalive with 200 OK
POST /ping -> unexpected method

POST /login -> input login and password as body, expect 200 OK and valid access token
POST /login -> invalid input body, expect 400 client error

GET /login -> unexpected method

GET /users -> not available
GET /users/id with no auth header -> forbidden 403
GET /users/id with auth header -> 200 OK with valid body (surname hidden if underage)

## E2E tests


