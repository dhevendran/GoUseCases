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

	// "grpc-go/blog/blogpb"
	"github.com/vamsikrishnasiddu/gRPC-go-Projects/GRPC-GO-COURSE/blog/blogpb"

	"io"
	"log"

	"github.com/spf13/cobra"
)

// ListBlogCmd represents the ListBlog command
var ListBlogCmd = &cobra.Command{
	Use:   "ListBlog",
	Short: "List in MongoDB",
	Long:  `This function is to list the data from MongoDB`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ListBlog called")
		doListBlog(BlogClient)
	},
}

func init() {
	rootCmd.AddCommand(ListBlogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ListBlogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ListBlogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doListBlog(c blogpb.BlogServiceClient) {
	//ListBlog
	stream, listerr := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if listerr != nil {
		log.Fatalf("Error while listing blog : %v", listerr)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error while recieving:%v", err)
		}
		fmt.Println(res.GetBlog())
	}

}
