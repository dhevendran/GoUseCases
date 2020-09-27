package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"../../greet/greetpb"

	"google.golang.org/grpc"

	"database/sql"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "emp"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("greet fun invoked %v", req)
	firstname := req.GetGreeting().GetFirstName()
	lastname := req.GetGreeting().GetLastName()
	response := "hello" + firstname + "" + lastname

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Opening a connection to the database
	// The sql.Open() function takes two arguments - a driver name, and a string that tells that driver how to connect to our database - and then returns a pointer to a sql.DB and an error.
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// db.Ping() forces our code to actually open up a connection to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", db)
	fmt.Println("Successfully connected!")

	sqlStatement := `
	SELECT * FROM emp_table WHERE role_id=$1
	RETURNING role_id`
	role_id := ""
	err = db.QueryRow(sqlStatement, "3").Scan(&role_id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", role_id)

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
	fmt.Println("Server started")

	//Just to log the errors
	fmt.Println("Main is working")
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatal("failed", err)
	}
	fmt.Println("After net.Listen")
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})
	fmt.Println("After RegisterGreetServiceServer")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
	fmt.Println("At End")

}
