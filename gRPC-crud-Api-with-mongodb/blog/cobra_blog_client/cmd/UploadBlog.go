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
	"bufio"
	"context"
	"fmt"
	// "grpc-go/blog/blogpb"
	"github.com/vamsikrishnasiddu/gRPC-go-Projects/GRPC-GO-COURSE/blog/blogpb"

	"log"
	"os"

	"github.com/spf13/cobra"
)

// UploadBlogCmd represents the UploadBlog command
var UploadBlogCmd = &cobra.Command{
	Use:   "UploadBlog",
	Short: "Upload a file",
	Long:  `To upload a file to db`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("UploadBlog called")
		doFileUpload(BlogClient, args)
	},
}

func init() {
	rootCmd.AddCommand(UploadBlogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// UploadBlogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// UploadBlogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doFileUpload(c blogpb.BlogServiceClient, args []string) {
	defer Conn.Close()
	if len(args) != 1 {
		log.Fatalf("Enter File name to upload with path\n")
	}
	file, err := os.Open(args[0])
	if err != nil {
		log.Fatalf("Error while opening file: %v", err)
	}
	defer file.Close()
	stream, err := c.FileUpload(context.Background())
	if err != nil {
		log.Fatalf("Error while uploading : %v", err)
	}
	defer stream.CloseSend()

	s := bufio.NewScanner(file)
	for s.Scan() {
		err = stream.Send(&blogpb.FileUploadRequest{
			FileName: args[0],
			Content:  []byte(s.Text()),
		})
		if err != nil {
			log.Fatalf("Error while sending data : %v", err)
			break
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving status : %v", err)
	}

	fmt.Printf("Message from server :%v", resp.Message)
}
