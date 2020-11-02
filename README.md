# LogApi

Given the attached logfile (log.txt), write an app that parses the log file and
outputs the following information to `stdout`:

1. Amount of errors by service name
2. Instance id of the service instance with most errors

Examples:

```
[api-gateway]:  17 errors
[ffd3082fe09d]: 17/17 errors
```

> **NOTE**: these are just for validation purposes, you don't have to match the format

Hints:

Note the app log has the following format: `%DATE% [service-name instance-id]: log-trace`
Errors include the `[error]` string on the trace.

Think about how you would handle concurrent calls to you app and things like
file access and size.

Optional:

  - Dockerize your app.
  - Expose a common API for you app so that it can be reached through HTTP,
    preferably using a RESTful approach

## TODO

- [ ] Handle file update
- [ ] Read only parts of the file that were not read
