# Simple Task Manager Api
API for simple task manager


## Package used
* [gorilla/mux](https://github.com/gorilla/mux): Used to manage routes
* [GoDotEnv](https://github.com/joho/godotenv): Used to read .env

## Endpoints
All calls must have /api at the begining og thhe path

|Route|Method|Description|
|--|--|---|
|/keepalive|GET|Keepalive of the services its return "ok"|
