# hestia
Registration server

## Configuration

Environment variables: 

| Variable                | Default value | Description                                |
|-------------------------|---------------|--------------------------------------------|
| HESTIA_PORT             | 9000          | The port the application listens on        |
| HESTIA_DB_HOST          |               | The host of the users database             |
| HESTIA_DB_PORT          |               | The port of the users database             |
| HESTIA_PROFILES_HOST    |               | The host of the profiles service           |
| HESTIA_PROFILES_PORT    |               | The port of the profiles service           |

## API (TODO)

### SignUp
Request:

    grpcurl -plaintext -d '{\"username\": \"Marem\", \"email\": \"marem@tortugasoftworks\", \"password\": \"1234\"}' localhost:9000 proto.Registration/SignUp

Response:

    {
    "userTag": "Marem#00000"
    }

## Build

The application is meant to be built using the provided Dockerfile. However, you can also do it manually.

Requirements:
- Go (v1.20)

Command:
    
    go build -o hestia ./cmd

This is assuming the gRPC files have been generated already. If they are not, please reading the following section.

## Development

### Requirements: 
- Go (v1.20)
- Protocol buffer compiler (v3)
- Go plugins:
    - protoc-gen-go (v1.28)
    - protoc-gen-go-grpc (v1.2)

Note: Make sure protoc (Protocol buffer compiler) can find the plugins in the Path environment variable

See https://grpc.io/docs/languages/go/quickstart/

### Generating gRPC source files
To generate the gRPC server and client source files:
    
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/registration.proto
