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

## Feedback

Overall: 

There is a lot produced, so much that I almost think it misses the mark. The extra additions around it also work poorly in practice. The code separation is rather good but the dependencies between them are rather troubling and a source of several issues. There are enough tests to verify the implementation (although with the big assumption of sidecar containership), although this lacks negatives/creative contributions to the test samples which is where issues are found. The difficulty in tracing these errors was excessive - once digging for one error was started this kept becoming hairier and hairier.

Shortcomings in descending order of importance/value:

In parser.go:85 at setDetails, the string[] is sorted with no purpose and breaking formats - there is no string-defined-order between the tuple of ("api-gateway", "ffd3082fe09d"). One is always one and the other is always the other for any and all values of one and other. This fails for the tuple ("api-gateway", "aad3082fe09d") as the other is sorted before the one and treated as the wrong kind and results in erroneous data inserted into the store. Were a struct for this created this would be a non-issue.
Candidate provided to run this as a sidecar container, however this only reports the one and first error. The kubernetes part is entirely optional (not even part of the exercise per se) however it is non-functioning. The candidate is rated on the submission which includes the non-functioning sidecar. A naive implementation of recounting everything every request would have been acceptable with the listed todos. Exiting and having kubernetes restart the container to re-read the file would also have been acceptable.
The implementation in file.go handles a lot of naive buffering - for questionable gain and difficult to read code. Mixing file.Read and file.ReadBytes and is the source of missing a line:
A log file at exactly the candidates buffer size plus one log entry will fail the last line: when file.ReadBytes encounters EOF it still returns all bytes read up to that point.
In store.go:21 candidate does not defer locking immediately after acquiring the lock in MockStore#Bump, thus it may fail on errors until just before end of function.
