package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"../../greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("hi, i am client")
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnot connect %v", err)

	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)

	//doUnary(c)
	//fmt.Printf("created client: %f", c)

	//doUnarySum(c)

	doServerStream(c)

	//doClientStream(c)

}
func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("unary rc starting...\n")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "hema",
			LastName:  "deepika",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("errer while calling rpc %v", err)
	}
	log.Printf("the response from greet %v", res.Response)
}
func doUnarySum(c greetpb.GreetServiceClient) {
	fmt.Println("unary rc starting...for sum \n")
	req := &greetpb.SumRequest{
		Suming: &greetpb.Suming{
			FirstNum: 10,
			LastNum:  3,
		},
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("error while calling rpc for sum %v", err)
	}
	log.Printf("the response from sum is %v", res.Response)

}

func doServerStream(c greetpb.GreetServiceClient) {

	fmt.Println("server stream fun invokes \n")

	req := &greetpb.PrimedecoRequest{
		Num: 31313,
	}
	stream, err := c.PrimeNumberDecomposition(context.Background(), req)
	if err != nil {
		log.Fatalf("erreor wil calling serevr streaming rpc %v", err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("something happend", err)
		}
		fmt.Println(res.GetPrimeRes())
	}

}

// func doClientStream(c greetpb.GreetServiceClient){
// 	fmt.Println("server stream fun invokes \n")
// 	stream,err:=c.ComputeAverage(context.Background())

// 	if err!=nil{
// 		log.Fatalf("error wile streaming%v", err)
// 	}
// 	numbers:=[]int64{4,5,2,44}

// 	for _,number:=range numbers{
// 		stream.Send(&greetpb.CompAverageRequest{
// 			Number:number,
// 		})

// 	}
// 	res,err:=stream.CloseAndRecv()
// 	if err!=nil{
// 		fmt.Println("error in closeing stream")
// 	}
// 	fmt.Printf("the average result is \n %v", res.GetAverageRes())

// }
