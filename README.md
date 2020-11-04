# LogApi

Logapi is a tool to run alongside your containers that will allow you to interpret logs and serve them as a rest api

## Usage

All the command are available in the `Makefile` but the easiest way is to `docker-compose up web` and `curl 127.0.0.1:8080/files` you can also try `tilt up` to test the application directly on kubernetes
All commands will use the example log file ´examples/log.txt'

## Test

'go test pkg/...'

## Architecture

```
📦 cmd
│ 📜 *.go # cli definition
📦 pkg
└───📂 conf # Configuration management and stop signal handling
└───📂 file # The file pagckage that handles log parsing and http translation
└───📂 server # Handles the http server bolierplate
└───📂 store # Defines the storage
📦 kustomize # Kubernetes manifests
📜 config.yaml # The default config for the app
📜 logapi.go # common interfaces
📜 Dockerfile # The application image definition
📜 Tiltfile # tilt is a tool that allows you to run your code with live reload on your local k8s cluster
📜 docker-compose.yaml # To make docker commands easier
📜 Makefile # Colletion of usefull commands

```

## TODO (Things that are kinda of outside scope)

- [ ] Handle file update (Maybe watch the file)
- [ ] Read only parts of the file that were not read (Maybe index the reads)
- [ ] handle multiple files (Maybe use the file name as index)
- [ ] Maybe Create A view layer for store
- [ ] Handle more test cases
