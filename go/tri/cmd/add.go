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
	"errors"
	"io/fs"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/younger-1/pkg/go/tri/todo"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add todo items",
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(viper.GetString("dataFile"))
		if err != nil {
			var pathError *fs.PathError
			if errors.As(err, &pathError) {
			} else {
				log.Printf("%v", err)
			}
		}

		for _, v := range args {
			item := todo.Item{Text: v}
			item.SetPriority(priority)
			items = append(items, item)
		}
		// fmt.Printf("%#v\n", items)
		err = todo.SaveItems(viper.GetString("dataFile"), items)
		if err != nil {
			log.Printf("%v", err)
		}
	},
}

var priority int

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")
}
