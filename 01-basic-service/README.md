# Basic Service

- Directories
- CLI 
- Logging
- Environment Vars

## Directories

Initially, we'll have only two directories at the root 
level of our service:

- `bin/` - of course, where binaries go
- `cmd/basic-service/` - our service entrypoint

## CLI

Our service entrypoint will be defined inside the `cmd/` 
directory, therefore representing an executable command. 
In this case, the service is implemented as the command 
itself (`cmd/basic-service/main.go`). 

This makes it possible to run a simple 
`go build  -o bin/basic-service cmd/basic-service/*.go` from
the service's root directory, which will create a
`basic-service` binary executable for our service.