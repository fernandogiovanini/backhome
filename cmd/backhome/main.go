package main

import (
	"os"

	cmd "github.com/fernandogiovanini/backhome/internal/command"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
