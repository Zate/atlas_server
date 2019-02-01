# Dockerized Atlas Linux Game Server

## Components

* atlas game server container ([atlas](https://github.com/Zate/atlas_server/tree/master/atlas)) - 1 per atlas tile (1x1 grid default)
* gogasm management container ([gogasm](https://github.com/Zate/atlas_server/tree/master/gogasm)) - web rcon, server config, live atlas grid viewer
* redis container ([redis](https://github.com/Zate/atlas_server/tree/master/redis)) - only visible to the game server and management containers
* traefik proxy container ([traefik](https://github.com/Zate/atlas_server/tree/master/traefik)) - allowing external access to manage the grid viaa browser
* atlas server installer/updater ([cmd](https://github.com/Zate/atlas_server/tree/master/cmd)) - allows installing/setup/update of the atlas server grid

### Status

#### atlas

* atlas-s docker image can build and be brought up via docker-compose
* env supplied by .env-atlas-s

#### gogasm

* atlas-g docker image can be built and deployed via docker-compose
* gogasm binary can be invoked to provide live realm status, server qury of a server, rcon SaveWorld/DoExit and rcon command
* `docker-compose exec atlas-g gogasm -live <napve|napvp|eupve|eupvp>`
* `docker-compose exec atlas-g gogasm -s <serverIP> -p <queryPort>`
* `docker-compose exec atlas-g gogasm -s <serverIP> -p <queryPort> -r <rconPort>`
* `docker-compose exec atlas-g gogasm -s <serverIP> -p <queryPort> -r <rconPort> -c <rconCommand>`
* `docker-compose exec atlas-g gogasm -web <webPort>`
* `docker-compose exec atlas-g gogasm -h`
* env supplied by .env-atlas-g
* TODO : vue.js interface

#### redis

* atlas-r docker image can be built and deployed via docker-compose
* redis is only visible to the other containers, does not bind to 0.0.0.0 on the docker host
* uses the default redis.conf supplied in AtlasTools
* password supplied via command line in docker-compose right now
* env supplied by .env-atlas-r

#### traefik

* TODO : build atlas-t container to listen on 80/443 and supply reverse proxy to internal web services
* TODO : setup auth to traefik admin
* TODO : setup auth via cloudflare access to management interface
* TODO : setup .env-atlas-t

#### Server setup binary

* only skeleton from cobra right now.
* TODO : build out checkupdate, help, setup, update commands