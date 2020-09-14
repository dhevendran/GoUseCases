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
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// subCmd represents the sub command
var subCmd = &cobra.Command{
	Use:   "sub",
	Short: "sub <num-1> <num-2> --> Subtracts num-1 by num-2",
	Long:  `sub <num-1> <num-2> --> Subtracts num-1 by num-2 and displays the result`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("sub called")
		subInt(args)
	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func subInt(args []string) {
	// iterate over the arguments
	// the first return value is index of args, we can omit it using _

	if len(args) != 2 {
		fmt.Println("Minimum two arguments are needed")
		os.Exit(1)
	}

	itemp1, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
	}
	itemp2, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(err)
	}
	result := itemp1 - itemp2
	fmt.Printf("Sub of numbers %s is %d", args, result)
}
