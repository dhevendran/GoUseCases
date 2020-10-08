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

// ReadBlogCmd represents the ReadBlog command
var ReadBlogCmd = &cobra.Command{
	Use:   "ReadBlog",
	Short: "Read in MongoDB",
	Long:  `This function is to read the data from MongoDB`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ReadBlog called")
		doReadBLog(BlogClient, args)
	},
}

func init() {
	rootCmd.AddCommand(ReadBlogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ReadBlogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ReadBlogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doReadBLog(c blogpb.BlogServiceClient, args []string) {
	//Read Blog
	defer Conn.Close()
	if len(args) != 1 {
		log.Fatalf("Enter 'ID'\n")
	}

	blogID := args[0]
	readBlogReq := &blogpb.ReadBlogRequest{BlogId: blogID}
	readRes, errread := c.ReadBlog(context.Background(), readBlogReq)
	if errread != nil {
		log.Fatalf("\nError while reading: %v", errread)
	}

	fmt.Printf("\nBlog Was Read : %v", readRes)
}
