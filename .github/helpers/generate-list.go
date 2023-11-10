package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println([]string{})
		return
	}

	fmt.Println(strings.ReplaceAll(fmt.Sprintf("%q", args), " ", ", "))
}