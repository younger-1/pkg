package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/younger-1/code-playground/go/simple/morestrings"
)

func main() {
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}
