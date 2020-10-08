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
	"os"

	"github.com/spf13/cobra"
)

// DownloadBlogCmd represents the DownloadBlog command
var DownloadBlogCmd = &cobra.Command{
	Use:   "DownloadBlog",
	Short: "Download a file",
	Long:  `Download a file`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DownloadBlog called")
		doFileDownload(BlogClient, args)
	},
}

func init() {
	rootCmd.AddCommand(DownloadBlogCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// DownloadBlogCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// DownloadBlogCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func doFileDownload(c blogpb.BlogServiceClient, args []string) {
	defer Conn.Close()
	if len(args) != 1 {
		log.Fatalf("Enter 'File Name to Download'\n")
	}
	req := &blogpb.FileDownloadRequest{
		FileName: args[0],
	}
	res, err := c.FileDownload(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greet rpc: %v", err)
	}

	file, err := os.OpenFile(args[0], os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		log.Fatalf("Error while recieve streaming: %v", err)
	}
	defer file.Close()

	_, err = file.WriteString(string(res.Content))
	if err != nil {
		log.Fatalf("Error while writing to file: %v", err)
	}

}
