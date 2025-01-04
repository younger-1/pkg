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
	"fmt"
	"io/fs"
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/younger-1/code-playground/go/tri/todo"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done",
	Short: "Mark items as done",
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(dataFile)
		if err != nil {
			var pathError *fs.PathError
			if errors.As(err, &pathError) {
				fmt.Println("No todo yet, please add first")
			} else {
				log.Printf("%v", err)
			}
		}
		for _, label := range args {
			i, err := strconv.Atoi(label)
			if err != nil {
				log.Printf("%v is not a valid label\n", label)
				continue
			}

			i -= 1
			if i < 0 || i >= len(items) {
				log.Printf("Label %v doesn't match any item\n", label)
				continue
			}

			if !invertFlag {
				items[i].Done = true
				fmt.Printf("%q is marked done\n", items[i].Text)
			} else {
				items[i].Done = false
				fmt.Printf("%q is marked todo\n", items[i].Text)
			}
		}
		err = todo.SaveItems(dataFile, items)
		if err != nil {
			log.Printf("%v", err)
		}
	},
}

var invertFlag bool

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	doneCmd.Flags().BoolVarP(&invertFlag, "invert", "i", false, "Invert done status")
}
