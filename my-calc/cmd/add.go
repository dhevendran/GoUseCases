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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add <num-1> <num-2> ... --> Adds <num-n> --> Adds num-1, num-2 .... num-n",
	Long:  `add <num-1> <num-2> ... --> Adds <num-n> --> Adds num-1, num-2 .... num-n and displays the result`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
		addInt(args)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}

func addInt(args []string) {
	var sum int
	// iterate over the arguments
	// the first return value is index of args, we can omit it using _

	for _, ival := range args {
		// strconv is the library used for type conversion. for string
		// to int conversion Atio method is used.
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sum)
}