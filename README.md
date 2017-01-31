# Watcher

Watcher is a lightweight, web-based network and process manager. Watcher provides easy-to-use displays for showing network and process management.

## Frontman

Frontman is a React based frontend for displaying computer metrics.

## Lookout

Lookout handles networking and process inspection. It formats data that can be stored into SQL or NoSQL databases. A lookout client can be installed on devices for remote monitoring.

## Sleuth

Slueth aggrogates and annotates metrics stored within databases. Slueth serves out meaningful data through REST calls.

# Running the project

In order to run the project follow these steps:
- Create an environment variable called GOPATH: `export GOPATH=/path/to/directory`
- Install npm
- In the top level folder, run npm install
- Run `grunt build`
- Run `grunt watch`
- Run the command `go get` to install Go dependencies. You can ignore the errors.
- Run the command `go build main.go && ./main`
