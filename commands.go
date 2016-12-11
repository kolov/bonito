package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/kolov/sardine/command"
)

var AuthToken string
var ListOrder string

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
		Name:   "list",
		Usage:  "list all snapshots or droplets",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "order , o ",
				Usage:  "order by n(ame) or d(ate) `FIELD`",
				Destination: &ListOrder},
		},
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
		Name:   "up",
		Usage:  "starts a droplet from a snapshot",
		Action: command.CmdUp,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template , t ",
				Usage:  "regex or fullname of the snapshot to use. More matches allowd, see --latest",
				Destination: &command.SnapshotTemplate},
			cli.StringFlag{Name: "snapshotid , sid ",
				Usage:  "id of the snapshot to use. Exact match expected",
				Destination: &command.SnapshotId},
			cli.BoolFlag{Name: "latest",
				Usage:  "Use the latest snapshot by more matches. ",
				Destination: &command.UseLatestSnapshot},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
