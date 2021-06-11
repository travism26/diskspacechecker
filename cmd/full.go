/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"strings"

	"github.com/spf13/cobra"
)

// fullCmd represents the full command
var fullCmd = &cobra.Command{
	Use:   "full",
	Short: "This will run an full scan on the computer looking for disk space issues",
	Long: ` For example:
	diskspacecheck full --config CONFIG_FILE
	diskspacecheck full --config CONFIG_FILE --output=/PATH/TO/output.json`,
	Run: func(cmd *cobra.Command, args []string) {
		output, outErr := cmd.Flags().GetString("output")
		if outErr != nil {
			fmt.Println(outErr)
		}
		path, pathErr := cmd.Flags().GetString("path")
		if pathErr != nil {
			fmt.Println(pathErr)
		}

		fmt.Println("full called")
		fmt.Println("Here are the arguments of card command : " + strings.Join(args, ","))
		fmt.Println("Value of the flag output: " + output + " path:" + path)
		walkDir := NewBasicScanner(path, output)
		largeFileScanner := NewLargeFileFinder(path, output, 110)
		scanners := []scanner{}
		scanners = append(scanners, walkDir)
		scanners = append(scanners, largeFileScanner)

		for _, scan := range scanners {
			scan.test()
			scan.scan()
		}
	},
}

func init() {
	rootCmd.AddCommand(fullCmd)
	// 	fullCmd.PersistentFlags().StringP("config", "", "Path to config file")
	fullCmd.PersistentFlags().StringP("output", "o", "", "Write the output to a file (default format: json)")
	fullCmd.PersistentFlags().StringP("path", "p", "~/", "Write the output to a file")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fullCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fullCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
