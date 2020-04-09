package main

import (
	"os"

	"github.com/etherandrius/uqc/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
