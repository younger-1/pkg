/*
Copyright © 2025 Xavier Young

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
	"log"

	"github.com/spf13/cobra"
	"github.com/younger-1/code-playground/go/tri/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add new todo items",
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems("./todo.txt")
		if err != nil {
			log.Printf("%v", err)
		}

		for _, v := range args {
			items = append(items, todo.Item{Text: v})
		}
		// fmt.Printf("%#v\n", items)
		err = todo.SaveItems("./todo.txt", items)
		if err != nil {
			log.Printf("%v", err)
		}
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
