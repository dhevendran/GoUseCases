package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"

	"../../db/dbpb"

	"google.golang.org/grpc"

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

var db *sql.DB = nil

func (*server) MyPost(ctx context.Context, req *dbpb.PostMsgRequest) (*dbpb.PostMsgResponse, error) {
	fmt.Printf("MyPost fun invoked %v\n", req)
	firstname := req.GetMsg().GetFirstName()
	lastname := req.GetMsg().GetLastName()
	id := req.GetMsg().GetId()

	// Write query to store from DB for a given id
	sqlStatement := `
	INSERT INTO emp_table (role_id,first_name, last_name)
	VALUES ($1,$2, $3)
	RETURNING first_name`
	queriedFirstName := "ToBeQueried"

	fmt.Println("db.QueryRow :", ", First Name: "+firstname+", Last Name: "+lastname+", ID : "+id)
	err := db.QueryRow(sqlStatement, id, firstname, lastname).Scan(&queriedFirstName)
	if err != nil {
		fmt.Println("The error message is :" + err.Error())
		return nil, err
	}
	response := "First Name: " + queriedFirstName + ", Last Name: " + lastname + ", ID : " + id
	res := &dbpb.PostMsgResponse{
		Response: response,
	}
	fmt.Println("Success MyPost Response : " + response)
	return res, nil
}

func (*server) MyGet(ctx context.Context, req *dbpb.GetMsgRequest) (*dbpb.GetMsgResponse, error) {
	fmt.Printf("MyGet fun invoked %v\n", req)
	id := req.GetId()
	roleID, firstName, lastName := "ToBeQueried", "ToBeQueried", "ToBeQueried"
	sqlStatement := `SELECT * FROM emp_table WHERE role_id=$1`
	fmt.Println("Before db.QueryRow()")
	err := db.QueryRow(sqlStatement, id).Scan(&roleID, &firstName, &lastName)
	fmt.Println("Before db.QueryRow()")

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	res := &dbpb.GetMsgResponse{
		Msg: &dbpb.Msg{
			FirstName: firstName,
			LastName:  lastName,
			Id:        roleID,
		},
	}

	response := "First Name: " + firstName + " Last Name: " + lastName + " ID : " + roleID
	fmt.Println("Success MyGet Response : " + response)
	return res, nil
}

func (*server) MyDelete(ctx context.Context, req *dbpb.GetMsgRequest) (*dbpb.GetMsgResponse, error) {
	fmt.Printf("MyDelete fun invoked %v\n", req)
	id := req.GetId()

	sqlStatement := `DELETE FROM emp_table WHERE role_id=$1`
	err := db.QueryRow(sqlStatement, id).Scan()
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}

	res := &dbpb.GetMsgResponse{
		Msg: &dbpb.Msg{
			FirstName: "",
			LastName:  "",
			Id:        id,
		},
	}

	fmt.Println("Success MyDelete Response")
	return res, nil
}

func main() {
	fmt.Println("*** Server Started ***")

	//Just to log the errors
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Server failed to serve %v\n", err)
	}

	openDbConnection()
	defer db.Close()

	s := grpc.NewServer()
	dbpb.RegisterGetPostServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server failed to serve %v\n", err)
	}
	fmt.Println("*** Server End ***")

}

func openDbConnection() {
	if db != nil {
		return
	}
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error = nil
	// Opening a connection to the database
	// The sql.Open() function takes two arguments - a driver name, and a string that tells that driver how to connect to our database - and then returns a pointer to a sql.DB and an error.
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	// db.Ping() forces our code to actually open up a connection to the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully database connected!")
}
