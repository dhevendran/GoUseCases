package main

import (
	"context"
	"fmt"
	"log"

	"../../db/dbpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("*** Client Started ***")
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnot connect %v\n", err)

	}
	defer conn.Close()
	c := dbpb.NewGetPostServiceClient(conn)

	//doPost(c)

	//doGet(c)

	doDelete(c)

	fmt.Println("*** Client End ***")

}
func doPost(c dbpb.GetPostServiceClient) {

	req := &dbpb.PostMsgRequest{
		Msg: &dbpb.Msg{
			FirstName: "Dhevendran",
			LastName:  "Kulandaivelu",
			Id:        "10",
		},
	}
	res, err := c.MyPost(context.Background(), req)
	if err != nil {
		log.Fatalf("Errer while calling rpc %v\n", err)
	}
	log.Printf("Success MyPost response : %v\n", res.Response)

}

func doGet(c dbpb.GetPostServiceClient) {

	id := "10"
	req := &dbpb.GetMsgRequest{
		Id: id,
	}
	res, err := c.MyGet(context.Background(), req)
	if err != nil {
		log.Fatalf("Errer while calling rpc %v\n", err)
	}
	log.Printf("Success MyGet response : %s %s %s\n", res.GetMsg().GetFirstName(), res.GetMsg().GetLastName(), res.GetMsg().GetId())

}

func doDelete(c dbpb.GetPostServiceClient) {
	fmt.Println("*** doDelete Started ***")
	id := "10"
	req := &dbpb.GetMsgRequest{
		Id: id,
	}
	res, err := c.MyDelete(context.Background(), req)
	if err != nil {
		log.Fatalf("Errer while calling rpc %v\n", err)
	}
	log.Printf("Success MyDelete response : %s %s %s\n", res.GetMsg().GetFirstName(), res.GetMsg().GetLastName(), res.GetMsg().GetId())
	fmt.Println("*** doDelete End ***")
}
