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

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <ID>",
	Short: "Deletes DB entry for the given ID.",
	Long:  `Deletes DB entry for the given ID.. For example, 'clinet delete 3000'`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("delete called")
		deleteUser(args)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deleteUser(args []string) {
	fmt.Println("*** deleteConnectToServer Start ***")
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnot connect %v\n", err)

	}
	defer conn.Close()
	c := dbpb.NewGetPostServiceClient(conn)

	doDelete(c, args)

	fmt.Println("*** deleteConnectToServer End ***")

}

func doDelete(c dbpb.GetPostServiceClient, args []string) {
	fmt.Println("*** doDelete Started ***")
	if len(args) != 1 {
		log.Fatalf("Enter 'ID'\n")
	}
	req := &dbpb.GetMsgRequest{
		Id: args[0],
	}
	res, err := c.MyDelete(context.Background(), req)
	if err != nil {
		log.Fatalf("Errer while calling rpc %v\n", err)
	}
	log.Printf("Success MyDelete response : %s %s %s\n", res.GetMsg().GetFirstName(), res.GetMsg().GetLastName(), res.GetMsg().GetId())
	fmt.Println("*** doDelete End ***")
}
