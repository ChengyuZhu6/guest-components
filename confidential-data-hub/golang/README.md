# Confidential Data Hub Golang Client

## Overview
This provides a unified client interface for interacting with Confidential Data Hub (`CDH`) using gRPC or TTRPC protocols. It's designed to be used in Go projects, such as Node Resource Interface (`NRI`) or image verifiers plugins in containerd. And you can also build it as a client binary to interact with `CDH`.

## Getting Started

### Usage as library

Import the package into your Go project:

```go
//grpc package 
import cdhgrpcapi "github.com/confidential-containers/guest-components/confidential-data-hub/golang/pkg/grpc"

//ttrpc package 
import cdhttrpcapi "github.com/confidential-containers/guest-components/confidential-data-hub/golang/pkg/ttrpc"
```

Create a new client instance:

```go
//cdh grpc client
c, err := cdhgrpcapi.CreateCDHGrpcClient("127.0.0.1:8043")

//cdh ttrpc client
c, err := cdhttrpcapi.CreateCDHTtrpcClient("/run/confidential-containers/cdh.sock")
```

Interact with `CDH` using the client, for example :
```go
unsealedValue, err := c.UnsealEnv(ctx, sealedSecret)
```

### Usage as binary

Build and Install the binary, such as:
```bash
$ make RPC=grpc
Building Go binaries...
GOARCH=amd64 go build -o bin/cdh-go-client ./cmd/grpc-client
$ sudo make install
Installing binaries...
install -D -m0755 bin/cdh-go-client /usr/local/bin
```

Interact with CDH using the binary, such as get sealed secret:
```bash
$ cdh-go-client -v sealed.fakeheader.ewogICJ2ZXJzaW9uIjogIjAuMS4wIiwKICAidHlwZSI6ICJ2YXVsdCIsCiAgIm5hbWUiOiAia2JzOi8vL2RlZmF1bHQvdHlwZS90YWciLAogICJwcm92aWRlciI6ICJrYnMiLAogICJwcm92aWRlcl9zZXR0aW5ncyI6IHt9LAogICJhbm5vdGF0aW9ucyI6IHt9Cn0K.fakesignature
unsealed value from env = that's the unsealed secret
```
or get sealed secret from file:
```bash
$ cat <<EOF > sealedsecretfile
sealed.fakeheader.ewogICJ2ZXJzaW9uIjogIjAuMS4wIiwKICAidHlwZSI6ICJ2YXVsdCIsCiAgIm5hbWUiOiAia2JzOi8vL2RlZmF1bHQvdHlwZS90YWciLAogICJwcm92aWRlciI6ICJrYnMiLAogICJwcm92aWRlcl9zZXR0aW5ncyI6IHt9LAogICJhbm5vdGF0aW9ucyI6IHt9Cn0K.fakesignature
EOF
$ cdh-go-client -f sealedsecretfile 
unsealed value from file = that's the unsealed secret
```