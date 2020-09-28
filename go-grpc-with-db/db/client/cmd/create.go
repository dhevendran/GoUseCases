/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/dhevendran/GoUseCases/go-grpc-with-db/db/dbpb"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create <First Name> <Last Name> <ID>",
	Short: "Create a DB entry for the given detals.",
	Long:  `Create a DB entry for the given detals. For example, 'clinet create Dhevendran Kulandaivelu 3000'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create called")
		postUser(args)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func postUser(args []string) {
	fmt.Println("*** connectToServer Start ***")
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnot connect %v\n", err)

	}
	defer conn.Close()
	c := dbpb.NewGetPostServiceClient(conn)

	doPost(c, args)

	fmt.Println("*** connectToServer End ***")

}
func doPost(c dbpb.GetPostServiceClient, args []string) {

	if len(args) != 3 {
		log.Fatalf("Enter 'First Name' 'Last Name' and 'ID' with space seperation in the same order\n")
	}

	req := &dbpb.PostMsgRequest{
		Msg: &dbpb.Msg{
			FirstName: args[0],
			LastName:  args[1],
			Id:        args[2],
		},
	}
	res, err := c.MyPost(context.Background(), req)
	if err != nil {
		log.Fatalf("Errer while calling rpc %v\n", err)
	}
	log.Printf("Success MyPost response : %v\n", res.Response)

}
