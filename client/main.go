package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	pb "grpc-chat/grpcchat"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:50005"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect : %v", err)
	}
	defer conn.Close()

	c := pb.NewChatClient(conn)
	stream, err := c.ReceiveAndSend(context.Background())
	if err != nil {
		log.Fatalf("open stream error %v", err)
	}
	log.Println("connected to grpc chat server at ", addr)

	var wg sync.WaitGroup

	//go routine for send message
	wg.Add(1)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stream.Send(&pb.Message{Message: scanner.Text()})
			if err == io.EOF {
				wg.Done()
				return
			}
		}
	}()

	//go routine for receive message
	wg.Add(1)
	go func() {
		for {
			resp, err := stream.Recv()

			fmt.Println("Server => ", resp.GetMessage())
			if err == io.EOF {
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
}
