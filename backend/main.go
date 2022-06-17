package main

import (
	"flag"
	"fmt"
	"github.com/ITK13201/portfolio/backend/cmd"
	_ "github.com/ITK13201/portfolio/backend/cmd"
	"github.com/ITK13201/portfolio/backend/cmd/server"
	"github.com/ITK13201/portfolio/backend/config"
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
		fmt.Printf("%s v%s\n", flag.CommandLine.Name(), config.Version)
	case "batch":
		if err := cmd.Ops.Job.FlagSet.Parse(os.Args[2:]); err != nil {
			if err != flag.ErrHelp {
				fmt.Fprintf(os.Stderr, "error: %s\n", err)
			}
			os.Exit(2)
		}
		if cmd.Ops.Job.Help {
			cmd.Ops.Job.FlagSet.Usage()
			os.Exit(0)
		}

	default:
		fmt.Printf("not defined: %s\n", arg)
		flag.CommandLine.Usage()
		os.Exit(2)
	}
}
