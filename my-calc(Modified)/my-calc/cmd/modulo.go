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

// moduloCmd represents the modulo command
var moduloCmd = &cobra.Command{
	Use:   "modulo",
	Short: "modulo num1 num2 gives the remainder of num1 and num2",
	Long: `
		div 1 2  gives the remainder of 1%2 which is equal to 1.
		remainder is calculated using "%" operator.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		moduloNums(args)
	},
}

func init() {
	rootCmd.AddCommand(moduloCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// moduloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// moduloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func moduloNums(args []string) {
	a, _ := strconv.Atoi(args[0])
	b, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("modulo is not defined for floating point numbers.")
		return
	}

	fmt.Printf("the modulo of %s is %d.", args, a%b)
}
