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
		Name:   "letgo",
		Usage:  "Let a droplet go -shutdown, snapshot and destroy",
		Action: command.CmdShutdown,
		Flags:  []cli.Flag{
			cli.StringFlag{
				Name:   "name",
				Value: "",
				Usage:  "Droplet name",
				Destination: &common.DropletName,
			},
			cli.BoolFlag{
				Name:   "nosnapshot",
				Usage:  "Destroy without taking a snapshot",
				Destination: &common.NoSnapshot,
			},
			cli.BoolFlag{
				Name:   "verbose, v",
				Usage:  "Verbose output",
				Destination: &common.Verbose,
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
			cli.StringFlag{Name: "snapshot , s ",
				Usage:  "regex or name of the snapshot to use. More matches possible, see --latest",
				Destination: &common.SnapshotTemplate},
			cli.StringFlag{Name: "name ",
				Usage:  "Name of the droplet to create. If not provided, will use bonito-{timesttmp}",
				Destination: &common.DropletName},
			cli.BoolFlag{Name: "latest",
				Usage:  "Use the latest snapshot by more matches. If not set, wil stop by more matches",
				Destination: &common.UseLatestSnapshot},
			cli.StringFlag{
				Name:   "keys",
				Value: "",
				Usage:  "SSH key name(s) to initialize the droplet with",
				Destination: &common.Keys,
			},
			cli.StringFlag{
				Name:   "size",
				Value: "",
				Usage:  "size of the droplet to create. Defaukt is 2gb. Posslible values: 2gb, 4gb, 8gb, 16gb, m-16gb, 32gb, m-32gb, 48gb, m-64gb, 64gb, m-128gb, m-224gb",
				Destination: &common.DropletSize,
			},
			cli.BoolFlag{
				Name:   "verbose, v",
				Usage:  "Verbose output",
				Destination: &common.Verbose,
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
