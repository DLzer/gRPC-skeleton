# gRPC Skeleton
![MessageService Check](https://github.com/DLzer/gRPC-skeleton/workflows/MessageService%20Check/badge.svg)

Testing the waters with gRPC and some small microservices.

Currently built-out:
- MessageService
- MessageService Test Coverage

# # Run it

```shell
# right now docker is not set up so it's cli only
go run main.go
```

Run the service
``` text
❯ grpcurl --plaintext -d '{Body: "Hello, World!"}' 0.0.0.0:9000 protos.messageservice.MessageService/Message
{
    "body": "Hello, World!"
}
```

# # Test it
```shell
cd messageservice/cmd && go test
```