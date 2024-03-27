# gRPC Petshop

Welcome to gRPC Petshop, a hands-on project to learn the basics of gRPC (Google Remote Procedure Call) and Protocol Buffers in a simulated petshop environment.

## Requirements

- [Golang v1.21.4](https://go.dev/doc/install)
- [Protobuf compiler v3.12.4](https://grpc.io/docs/protoc-installation/)

## Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/raissonsouto/grpc-petshop.git
    ```

2. Navigate to the project directory:

    ```bash
    cd petshop-simulator
    ```

3. Install dependencies:

    ```bash
    go mod tidy
    ```

## Running the Server

To start the server, run the following command:

```bash
go run main.go
```

The server should start listening on the specified port (default is `localhost:9191`).