package main

import (
	"fmt"
	"os"

	"github.com/did-twit/did-twit-cli/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
