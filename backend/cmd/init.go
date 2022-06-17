package cmd

import (
	"flag"
	"fmt"
	"io"
)

type Usage struct {
	name        string
	description string
}

type Info struct {
	name                    string
	description             string
	CommandUsages           []Usage
	ManagementCommandUsages []Usage
}

func (info *Info) printInfo(flagSet *flag.FlagSet) {
	var usages []Usage
	printUsages := func(w io.Writer, usages []Usage) {
		for i := 0; i < len(usages); i++ {
			fmt.Fprintf(w, "  %s\t%s\n", usages[i].name, usages[i].description)
		}
	}

	o := flagSet.Output()
	fmt.Fprintf(o, "\nUsage: %s\n", info.name)
	fmt.Fprintf(o, "\nDescription: %s\n", info.description)
	usages = info.CommandUsages
	if len(usages) > 0 {
		fmt.Fprintf(o, "\nCommands:\n")
		printUsages(o, info.CommandUsages)
	}
	usages = info.ManagementCommandUsages
	if len(usages) > 0 {
		fmt.Fprintf(o, "\nManagement Commands:\n")
		printUsages(o, info.ManagementCommandUsages)
	}

	fmt.Fprintf(o, "\nOptions:\n")
	flagSet.PrintDefaults()
}

type Job struct {
	FlagSet *flag.FlagSet
	Info    Info
	Help    bool
}

type Options struct {
	Info Info
	Help bool
	Job  Job
}

var Ops Options

func init() {
	// initial information
	Ops.Info = Info{
		name:        "portfolio",
		description: "portfolio api server",
		CommandUsages: []Usage{
			{name: "runserver", description: "Run api server"},
			{name: "version", description: "Show the app version information"},
		},
		ManagementCommandUsages: []Usage{
			{name: "job*", description: "Manage job scripts"},
		},
	}
	Ops.Job.Info = Info{
		name:        "job",
		description: "Run job scripts",
	}

	var flagSet *flag.FlagSet

	// initial flagset
	flagSet = flag.CommandLine
	flagSet.Init(Ops.Info.name, flag.ExitOnError)
	flagSet.Usage = func() { Ops.Info.printInfo(flagSet) }
	flagSet.BoolVar(&Ops.Help, "h", false, "Show how to use the app")

	// job sub commands
	Ops.Job.FlagSet = flag.NewFlagSet(Ops.Job.Info.name, flag.ExitOnError)
	flagSet = Ops.Job.FlagSet
	flagSet.Init(Ops.Job.Info.name, flag.ExitOnError)
	flagSet.Usage = func() { Ops.Job.Info.printInfo(flagSet) }
	flagSet.BoolVar(&Ops.Job.Help, "h", false, "Show how to use the command")
}
