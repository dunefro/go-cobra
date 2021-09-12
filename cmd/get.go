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
	"io"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Command to get the desired Cat",
	Long:  `This command will call https://cataas.com/ to get the desired cat`,
	Run: func(cmd *cobra.Command, args []string) {
		var catName string
		if len(args) >= 1 && args[0] != "" {
			catName = args[0]
		}
		URL := "https://cataas.com/cat/" + catName
		fmt.Println("Trying to get the image of a", catName, "cat")
		response, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()
		if response.StatusCode == 200 {
			out, err := os.Create(catName + ".png")
			if err != nil {
				fmt.Println(err)
			}
			defer out.Close()
			// This writes the output to a file
			_, err = io.Copy(out, response.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("Pictue of a", catName, "cat is saved at file", out.Name())
		} else {
			fmt.Println("Error in gettig an image of", catName, "cat")
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
