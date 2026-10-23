package main

import (
	"os"

	"github.com/stonkol/wallah/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		printErrorMessage("Error executing command: ")
		printErr(err)
		os.Exit(1)
	}
}
