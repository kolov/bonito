package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kolov/sardine/command"
)

var AuthToken string

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "DO_TOKEN_SARDINE",
		Name:   "token, t",
		Value:  "",
		Usage:  "Authentication token. Must be provided here or as ",
		Destination: &AuthToken,
	},
}

var Commands = []cli.Command{
	{
		Name:   "sometest",
		Usage:  "used during develpment for random tests. Ignore.",
		Action: command.CmdSomeTest,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "images",
		Usage:  "Show all images",
		Action: command.CmdImages,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "snapshots",
		Usage:  "Show all snapshots",
		Action: command.CmdListSnapshots,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "droplets",
		Usage:  "",
		Action: command.CmdListDroplets,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "list",
		Usage:  "",
		//Action: command.CmdList,
		Flags:  []cli.Flag{},
		Subcommands: []cli.Command{
			{
				Name:  "snapshots",
				Usage: "alist snapshots",
				Action: command.CmdListSnapshots,
			},
			{
				Name:  "droplets",
				Usage: "list droplets",
				Action: command.CmdListDroplets,
			},
		},
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
