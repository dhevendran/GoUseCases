package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"../../greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("greet fun invoked %v", req)
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()
	response := "hello" + firstname + "" + lastname
	res := &greetpb.GreetResponse{
		Response: response,
	}
	return res, nil
}

func (*server) Sum(ctx context.Context, req *greetpb.SumRequest) (*greetpb.SumResponse, error) {

	fmt.Printf("sum fun invoked %v", req)

	firstnum := req.GetSuming().GetFirstNum()
	lastnum := req.GetSuming().GetLastNum()
	response := firstnum + lastnum
	res := &greetpb.SumResponse{
		Response: response,
	}
	return res, nil
}

func (*server) PrimeNumberDecomposition(req *greetpb.PrimedecoRequest, stream greetpb.GreetService_PrimeNumberDecompositionServer) error {
	fmt.Printf("primedeco fun invoked %v", req)

	num := req.GetNum()
	divisor := int64(2)
	for num > 1 {
		if num%divisor == 0 {
			stream.Send(&greetpb.PrimedecoResponse{
				PrimeRes: divisor,
			})
			num = num / divisor
		} else {

			divisor++
		}
	}

	return nil

}

// this is for the client stream for computing average
func (*server) ComputeAverage(stream greetpb.GreetService_ComputeAverageServer) error {

	fmt.Println("client streaming func invoked")

	sum := int64(0)
	count := 0
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			average := float64(sum) / float64(count)
			stream.SendAndClose(&greetpb.CompAverageResponse{
				AverageRes: average,
			})

		}
		if err != nil {
			log.Fatalf("erreo wile passing numbers %v", err)
		}
		sum += req.GetNumber()
		count++
	}
	return nil
}
func main() {
	fmt.Println("hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatal("failed", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}

}
