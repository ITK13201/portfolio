package cmd

import (
	"flag"
	"fmt"
	"io"
)

type Runserver struct {
	FlagSet *flag.FlagSet
}

type Version struct {
	FlagSet *flag.FlagSet
}

type Options struct {
	Runserver Runserver
	Version   Version
	Help      bool
}

var ops Options

func (opts *Options) printUsage(w io.Writer) {
	dump := func(w io.Writer, name string, description string) {
		fmt.Fprintf(w, "  %s\t%s\n", name, description)
	}
	dump(w, "runserver", "Run api server")
	dump(w, "version", "Show the app version information")
}

func init() {
	flag.CommandLine.Init("portfolio", flag.ExitOnError)

	ops.Runserver.FlagSet = flag.NewFlagSet("runserver", flag.ExitOnError)
	ops.Version.FlagSet = flag.NewFlagSet("version", flag.ExitOnError)
	flag.BoolVar(&ops.Help, "h", false, "Show how to use the app")

	flag.CommandLine.Usage = func() {
		o := flag.CommandLine.Output()
		fmt.Fprintf(o, "\nUsage: %s\n", flag.CommandLine.Name())
		fmt.Fprintf(o, "\nDescription: portfolio api server\n")
		fmt.Fprintf(o, "\nCommands:\n")
		ops.printUsage(o)
		fmt.Fprintf(o, "\nOptions:\n")
		flag.PrintDefaults()
	}
}
