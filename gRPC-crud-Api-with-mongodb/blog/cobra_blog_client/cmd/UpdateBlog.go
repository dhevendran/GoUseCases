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

	"log"

	"github.com/spf13/cobra"
)

// UpdateBlogCmd represents the UpdateBlog command
var UpdateBlogCmd = &cobra.Command{
	Use:   "UpdateBlog",
	Short: "Update in MongoDB",
	Long:  `This function is to update the data to MongoDB`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UpdateBlog called")
		doUpdateBlog(BlogClient, args)
	},
}

func init() {
	rootCmd.AddCommand(UpdateBlogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UpdateBlogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UpdateBlogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doUpdateBlog(c blogpb.BlogServiceClient, args []string) {
	//UpdateBlog
	defer Conn.Close()
	if len(args) != 4 {
		log.Fatalf("Enter 'ID' 'Author' 'Blog Title' and 'Content' with space seperation in the same order\n")
	}
	blogID := args[0]
	newreq := &blogpb.UpdateBlogRequest{
		Blog: &blogpb.Blog{
			Id:       blogID,
			AuthorId: args[1],
			Title:    args[2],
			Content:  args[3],
		},
	}

	updateRes, updateErr := c.UpdateBlog(context.Background(), newreq)
	if updateErr != nil {
		fmt.Printf("\nError while updating: %v", updateErr)
	}
	fmt.Printf("\nBlog Was Updated : %v", updateRes)

	// time.Sleep(15 * time.Second)
}
