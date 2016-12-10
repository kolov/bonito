package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kolov/sardine/command"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "ENV_TOKEN",
		Name:   "token",
		Value:  "",
		Usage:  "",
	},
}

var Commands = []cli.Command{
	{
		Name:   "sometest",
		Usage:  "",
		Action: command.CmdSomeTest,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "droplets",
		Usage:  "",
		Action: command.CmdDroplets,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "",
		Action: command.CmdList,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "shutdown",
		Usage:  "",
		Action: command.CmdShutdown,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "start",
		Usage:  "",
		Action: command.CmdStart,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
