package main

import (
	"fmt"
	"os"
	"github.com/codegangsta/cli"
	"github.com/kolov/bonito/command"
	"github.com/kolov/bonito/common"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		EnvVar: "DO_TOKEN_BONITO",
		Name:   "token, t",
		Value:  "",
		Usage:  "Authentication token. Must be provided here or as ",
		Destination: &common.AuthToken,
	},

}

var Commands = []cli.Command{
	{
		Name:   "list",
		Usage:  "lists all snapshots, droplets or keys",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "order , o ",
				Usage:  "order by n(ame) or d(ate) `FIELD`",
				Destination: &common.ListOrder},
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
			{
				Name:  "keys",
				Usage: "list keys",
				Action: command.CmdListKeys,
			},
		},
	},
	{
		Name:   "shutdown",
		Usage:  "Stops, archives and deletes a droplet",
		Action: command.CmdShutdown,
		Flags:  []cli.Flag{
			cli.StringFlag{Name: "template , t ",
				Usage:  "regex or fullname of the snapshot used to start the droplet. More matches " +
					"NOT allowed",
				Destination: &command.SnapshotTemplate},
			cli.StringFlag{Name: "snapshotid , sid ",
				Usage:  "id of the snapshot to use. Exact match expected",
				Destination: &command.SnapshotId},
			cli.BoolFlag{
				Name:   "verbose, v",
				Usage:  "Verbose output",
				Destination: &common.Verbose,
			},
			cli.StringFlag{
				Name:   "name",
				Value: "",
				Usage:  "Droplet name",
				Destination: &common.DropletName,
			},
			cli.BoolFlag{
				Name:   "force, f",
				Usage:  "Don't ask confirmation",
				Destination: &common.Force,
			},
		},
	},
	{
		Name:   "up",
		Usage:  "starts a droplet from a snapshot",
		Action: command.CmdUp,
		Flags: []cli.Flag{
			cli.StringFlag{Name: "template , t ",
				Usage:  "regex or fullname of the snapshot to use. More matches allowed, see --latest",
				Destination: &command.SnapshotTemplate},
			cli.StringFlag{Name: "snapshotid , sid ",
				Usage:  "id of the snapshot to use. Exact match expected",
				Destination: &command.SnapshotId},
			cli.BoolFlag{Name: "latest",
				Usage:  "Use the latest snapshot by more matches. ",
				Destination: &command.UseLatestSnapshot},
			cli.BoolFlag{
				Name:   "verbose, v",
				Usage:  "Verbose output",
				Destination: &common.Verbose,
			},
			cli.StringFlag{
				Name:   "keys",
				Value: "",
				Usage:  "SSH key name(s) to initialize the droplet with",
				Destination: &common.Keys,
			},
			cli.StringFlag{
				Name:   "name",
				Value: "",
				Usage:  "Droplet name",
				Destination: &common.DropletName,
			},
			cli.BoolFlag{
				Name:   "force, f",
				Usage:  "Don't ask confirmation",
				Destination: &common.Force,
			},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
