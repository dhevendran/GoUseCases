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

// divCmd represents the div command
var divCmd = &cobra.Command{
	Use:   "div",
	Short: "divide num1 num2 -- divides the num1 and num2",
	Long: `
	IntegersCase:
	         div num1 num2   divides the numbers and displays the result.
    FloatCase:
	         div -f num1 num2 divides the nmbers and displays the result.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println(args)
		fval, _ := cmd.Flags().GetBool("float")
		if fval {
			divFloats(args)
		} else {
			divInts(args)
		}

	},
}

func init() {
	rootCmd.AddCommand(divCmd)

	divCmd.Flags().BoolP("float", "f", false, "divides float numbers")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func divInts(args []string) {
	if len(args) != 2 {
		fmt.Println("Maximum two arguments are needed.")
		os.Exit(1)
	}
	a, err1 := strconv.Atoi(args[0])

	if err1 != nil {
		fmt.Println("Use -f for floating numbers")
		return
	}
	b, err2 := strconv.Atoi(args[1])
	if err2 != nil {
		fmt.Println("Use -f for floating numbers")
		return
	}

	if b == 0 {
		fmt.Println("Cannot divide by 0")
	} else {
		fmt.Printf("the division of %s is %d.", args, a/b)
	}
}

func divFloats(args []string) {
	if len(args) != 2 {
		fmt.Println("Maximum two arguments are needed.")
		os.Exit(1)
	}
	a, _ := strconv.ParseFloat(args[0], 64)
	b, _ := strconv.ParseFloat(args[1], 64)

	if b == 0 {
		fmt.Println("Cannot divide by 0")
	} else {
		fmt.Printf("the division of %s is %f.", args, a/b)
	}
}
