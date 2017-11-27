package main

import (
	"log"
	"os"

	"github.com/caioleonhardt/dinamo/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
