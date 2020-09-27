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

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("get called")
		getUser(args)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getUser(args []string) {
	fmt.Println("*** getUser Start ***")
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("couldnot connect %v\n", err)

	}
	defer conn.Close()
	c := dbpb.NewGetPostServiceClient(conn)

	doGet(c, args)

	fmt.Println("*** getUser End ***")

}
func doGet(c dbpb.GetPostServiceClient, args []string) {

	if len(args) != 1 {
		log.Fatalf("Enter 'ID'\n")
	}

	req := &dbpb.GetMsgRequest{
		Id: args[0],
	}
	res, err := c.MyGet(context.Background(), req)
	if err != nil {
		log.Fatalf("Errer while calling rpc %v\n", err)
	}
	log.Printf("Success MyGet response : %s %s %s\n", res.GetMsg().GetFirstName(), res.GetMsg().GetLastName(), res.GetMsg().GetId())

}
