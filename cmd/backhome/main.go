package main

import (
	"fmt"
	"os"

	cmd "github.com/fernandogiovanini/backhome/internal/command"
)

func main() {
	if err := cmd.Execute(); err != nil {
		// TODO: How to properly hande error here?
		// What kind of error end up here considering
		// cobra framework.
		os.Exit(1)
	}
	fmt.Fprint(os.Stdout, "\n")
}
