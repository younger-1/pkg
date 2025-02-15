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
	"os"
	"sort"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/younger-1/pkg/go/tri/todo"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List todo items",
	Run: func(cmd *cobra.Command, args []string) {
		items, err := todo.ReadItems(viper.GetString("dataFile"))
		if err != nil {
			var pathError *fs.PathError
			if errors.As(err, &pathError) {
				fmt.Println("No todo yet, please add first")
			} else {
				log.Printf("%v", err)
			}
		}

		// log.Printf("%+v", items)
		sort.Sort(todo.ByPri(items))
		// log.Printf("%+v", items)

		w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
		for _, v := range items {
			if allFlag || v.Done == doneFlag {
				fmt.Fprintf(w, "%s\t%s\t%s\t%s\t\n", v.Label(), v.PrettyDone(), v.PrettyP(), v.Text)
			}
		}
		w.Flush()
	},
}

var (
	doneFlag bool
	allFlag  bool
)

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	listCmd.Flags().BoolVar(&doneFlag, "done", false, "Show 'done' items")
	listCmd.Flags().BoolVar(&allFlag, "all", false, "Show all items")
}
