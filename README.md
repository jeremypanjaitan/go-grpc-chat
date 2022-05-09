# About
gRPC support bidirectinoal streaming RPCs where both sides(server and client) send a seqeuence of messages using a read-write stream. Each stream operates independently, they don't have to wait until the response received or they don't have to wait until the response sent. You can visit to this link to read more about gRPC https://grpc.io/docs/what-is-grpc/core-concepts/

This simple chat application will demonstrate how the bidirectional mechanism works in gRPC. There will be server and client will communicate with each other. Server will send a message and client will receive or the opposite.

# Requirements
- go version go1.17.3 linux/amd64
- proto compiler installed with `libprotoc 3.6.1`

# How to run
- open to terminal, one terminal for server and the other for client
- run server `go run main.go`
- you will see  `server listening at //address`
- run client `go run client/main.go`
- you will see `connected to grpc chat server at //address`
- you can directly type the message into the terminal and then click enter to send the message

# Demo
![img](https://github.com/jeremypanjaitan/go-grpc-chat/blob/master/run-demo.gif)


