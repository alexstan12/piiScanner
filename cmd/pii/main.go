package main

import (
	"os"
	"piiScanner/pkg/cli"
)

func main() {
	if err := cli.NewDetectRootCmd().Execute(); err != nil {
		os.Exit(1)
	}
}