package main

import (
	"os"

	"github.com/ngc7293/plink/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
