# Watcher

Watcher is a lightweight, web-based network and process manager. Watcher provides easy-to-use displays for showing network and process management.

## Frontman

Frontman is a React based frontend for displaying computer metrics.

## Lookout

Lookout handles networking and process inspection. It formats data that can be stored into SQL or NoSQL databases. A lookout client can be installed on devices for remote monitoring.

## Sleuth

Slueth aggrogates and annotates metrics stored within databases. Slueth serves out meaningful data through REST calls.



# Running the project

## Installation
- Install Go. Instructions can be found at [golang.org](https://golang.org/doc/install)
- Create an environment variable called GOPATH: `export GOPATH=~/go`
  - Note: You may want to add this command to your `.bashrc` and/or `.bash_profile`
- Create a Go src directory (`mkdir -p ~/go/src`)
- Install Node.js. Installers and binaries can be found [here](https://nodejs.org/en/download/)
- Install Redis. Instructions can be found [here](https://redis.io/download#installation)
  - Note: To easy connect to redis, start the redis daemon

## Development 
- Within the Go src directory, clone [Watcher](https://github.com/MrGirlyMan/watcher)
- Cd into the watcher directory
- Install Go dependencies with `go get`
- Install Node.js dependencies with `npm install`
- Run `grunt build`
- Run `grunt watch`
