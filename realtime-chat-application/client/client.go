package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"../../realtime-chat-application/chat"

	"google.golang.org/grpc"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Must have a URL to connect as the first argument")
		return
	}
	conn, err := grpc.Dial(os.Args[1], grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := chat.NewChatClient(conn)
	stream, err := c.Chat(context.Background())
	if err != nil {
		panic(err)
	}
	waitc := make(chan struct{})
	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				panic(err)
			}
			fmt.Println(msg.User + ":" + msg.Message)
		}
	}()
	fmt.Println("Connection established,type \"quit\" or use Ctrl+C to exit the program")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := scanner.Text()
		if msg == "quit" {
			err := stream.CloseSend()
			if err != nil {
				panic(err)
			}
			break
		}
		err := stream.Send(&chat.ChatMessage{
			User:    os.Args[2],
			Message: msg,
		})
		if err != nil {
			panic(err)
		}
	}
	<-waitc
}
