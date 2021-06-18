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

// pathScanCmd represents the pathScan command
var pathScanCmd = &cobra.Command{
	Use:   "pathScan",
	Short: "A brief description of your command",
	Long: ` For example:
	diskspacecheck full --path "/Users/NAME/Downloads/"
	diskspacecheck full --path "/Users/NAME/Downloads/" --output=/PATH/TO/output.json`,
	Run: func(cmd *cobra.Command, args []string) {
		output, outErr := cmd.Flags().GetString("output")
		if outErr != nil {
			fmt.Println(outErr)
		}
		path, pathErr := cmd.Flags().GetString("path")
		if pathErr != nil {
			fmt.Println(pathErr)
		}
		size, _ := cmd.Flags().GetInt64("size")

		fmt.Println("full called")
		fmt.Println("Here are the arguments of card command : " + strings.Join(args, ","))
		fmt.Println("Value of the flag output: " + output + " path:" + path)
		fmt.Printf("Value of globals: size: %d \n", size)
		scanResults, _ := scan(path)
		fmt.Printf("Number of scan results: %d", len(scanResults))
		// for _, scandata := range scanResults {

		// }
		// fmt.Printf("Another global:%d \n", *testSize)
		// walkDir := NewBasicScanner(path, output)
		// largeFileScanner := NewLargeFileFinder(path, output, 110)
		// scanners := []scanner{}
		// scanners = append(scanners, walkDir)
		// scanners = append(scanners, largeFileScanner)

		// for _, scan := range scanners {
		// 	scan.test()
		// 	scan.scan()
		// }
		// scannedfiles := largeFileScanner.getScannedFiles()
	},
}

func init() {
	rootCmd.AddCommand(pathScanCmd)
	pathScanCmd.PersistentFlags().StringP("output", "o", "", "Write the output to a file (default format: json) WIP")
	pathScanCmd.PersistentFlags().StringP("path", "p", ".", "Folder to scan")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pathScanCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pathScanCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
