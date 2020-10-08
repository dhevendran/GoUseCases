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

// CreateBlogCmd represents the CreateBlog command

var CreateBlogCmd = &cobra.Command{
	Use:   "CreateBlog",
	Short: "Create in MongoDB",
	Long:  `This function is to create the data to MongoDB`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("CreateBlog called")
		doCreateBlog(BlogClient, args, cmd)
	},
}

func init() {
	CreateBlogCmd.Flags().StringP("author", "a", "", "Add an author")
	CreateBlogCmd.Flags().StringP("title", "t", "", "A title for the blog")
	CreateBlogCmd.Flags().StringP("content", "c", "", "The content for the blog")
	CreateBlogCmd.MarkFlagRequired("author")
	CreateBlogCmd.MarkFlagRequired("title")
	CreateBlogCmd.MarkFlagRequired("content")
	rootCmd.AddCommand(CreateBlogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// CreateBlogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// CreateBlogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func doCreateBlog(c blogpb.BlogServiceClient, args []string, cmd *cobra.Command) {
	//Create blog
	// if len(args) != 3 {
	// 	log.Fatalf("Enter 'Author' 'Blog Title' and 'Content' with space seperation in the same order\n")
	// }

	author, err := cmd.Flags().GetString("author")
	title, err := cmd.Flags().GetString("title")
	content, err := cmd.Flags().GetString("content")
	if err != nil {
		log.Fatalf("Unnexpected error:%v", err)
	}
	defer Conn.Close()

	fmt.Println("Creating the blog")
	req := &blogpb.CreateBlogRequest{
		Blog: &blogpb.Blog{
			AuthorId: author,
			Title:    title,
			Content:  content,
		},
	}
	bloRes, err := c.CreateBlog(context.Background(), req)
	if err != nil {
		log.Fatalf("Unnexpected error:%v", err)
	}
	fmt.Printf("\nCreated blog response: %v", bloRes)
}
