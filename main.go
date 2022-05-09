package main

import (
	"bufio"
	"fmt"
	pb "grpc-chat/grpcchat"
	"io"
	"log"
	"net"
	"os"
	"sync"

	"google.golang.org/grpc"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

type ChatServer struct {
	pb.UnimplementedChatServer
}

func (c *ChatServer) ReceiveAndSend(srv pb.Chat_ReceiveAndSendServer) error {
	log.Println("start new server")
	ctx := srv.Context()
	var wg sync.WaitGroup
	//go routine for receive message
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
			}
			req, err := srv.Recv()
			if err == io.EOF {
				wg.Done()
				log.Println("exit")
				break
			}
			fmt.Println("Client => ", req.GetMessage())
		}
	}()

	//go routine for send message
	wg.Add(1)
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			select {
			case <-ctx.Done():
				wg.Done()
				return
			default:
			}
			err := srv.Send(&pb.Message{Message: scanner.Text()})
			if err != nil {
				wg.Done()
				log.Fatalf(err.Error())
			}

		}
	}()
	wg.Wait()
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50005")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServer(s, &ChatServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
