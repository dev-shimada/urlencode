package main

import (
	"fmt"
	"os"

	"github.com/dev-shimada/urlencode/cmd"
)

// main is the entry point of the application
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
