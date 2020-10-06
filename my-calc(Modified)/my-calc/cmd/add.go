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
	Short: "add num1 num2 .... num n  ---> Adds the num1,num2,...num n",
	Long: `
		IntegersCase:
				add 1 2 .. n  adds the numbers and displays the result.
		FloatCase:
				add -f 1.2 3.4 5.5 ...n adds the nmbers and displays the result.
				`,
	Run: func(cmd *cobra.Command, args []string) {
		fstatus, _ := cmd.Flags().GetBool("float")
		if fstatus {
			addFloats(args)
		} else {
			addInts(args)

		}

	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().BoolP("float", "f", false, "Add Floating numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func addInts(args []string) {
	var sum int
	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println("Please use -f flag while using floating point numbers.")
			return
		}
		sum = sum + itemp
	}
	fmt.Printf("The addition of numbers is %s is %d", args, sum)
}

func addFloats(args []string) {
	var sum float64
	for _, ival := range args {
		itemp, err := strconv.ParseFloat(ival, 64)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Printf("The addition of numbers is %s is %f", args, sum)
}
