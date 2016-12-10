package command

import (
	"github.com/codegangsta/cli"
	"fmt"
)

var SnapshotTemplate string
var UseLatestSnapshot bool

func CmdUp(c *cli.Context) {
	if SnapshotTemplate == "" {
		fmt.Println("Mandatory --template not provided.")
		cli.ShowCommandHelp(c, "up")
		return
	}
	fmt.Println("Starting up from image matching ", SnapshotTemplate)

	record := QuerySnapshots()

	fmt.Println(record)

}
