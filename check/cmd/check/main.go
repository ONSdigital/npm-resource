package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/idahobean/npm-resource"
	"github.com/idahobean/npm-resource/check"
)

func main() {

	NPM := check.NewNPM()

	var request check.Request
	if err := json.NewDecored(os.Stdin).Decode(&request); err != nil {
		fatal("reading request from stdin", err)
	}

	var err error
	if requset.Srouce.PackageName == "" {
		err = errors.New("package name")
	}
	if err != nil {
		fatal("parameter required", err)
	}

	response, err := command.Run(request)
	if err != nil {
		fatal("running command", err)
	}

	if err := json.NewEncoder(os.Stdout).Encode(response); err != nil {
		fatal("writing response to stdout", err)
	}
}

func fatal(message string, err error) {
	fmt.Fprintf(os.Stderr, "error %s: %s\n", message, err)
	os.Exit(1)
}
