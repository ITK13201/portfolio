package main

import (
	"flag"
	"fmt"
	_ "github.com/ITK13201/portfolio/backend/cmd"
	"github.com/ITK13201/portfolio/backend/cmd/server"
	"os"
)

func main() {
	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		if err != flag.ErrHelp {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
		}
		os.Exit(2)
	}
	
	arg := flag.Arg(0)
	switch arg {
	case "":
		flag.CommandLine.Usage()
	case "runserver":
		server.Run()
	case "version":
		fmt.Printf("%s v%s\n", flag.CommandLine.Name(), "0.0.1")
	default:
		fmt.Printf("not defined: %s\n", arg)
		flag.CommandLine.Usage()
		os.Exit(2)
	}
}
