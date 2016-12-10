package command

import (
	"github.com/codegangsta/cli"
	"fmt"
	"regexp"
)

var SnapshotTemplate string
var UseLatestSnapshot bool

func CmdUp(c *cli.Context) {
	if SnapshotTemplate == "" {
		fmt.Println("Mandatory --template not provided.")
		cli.ShowCommandHelp(c, "up")
		return
	}
	fmt.Println("Starting up from snapshot matching ", SnapshotTemplate)

	record := QuerySnapshots()

	var reName = regexp.MustCompile(SnapshotTemplate)

	matches := []Snapshot{}
	for _, snapshot := range record.Snapshots {
		if reName.MatchString(snapshot.Name) {
			matches = append(matches, snapshot)
		}
	}
	if len(matches) == 0 {
		fmt.Print("No snapshots found matching [", SnapshotTemplate, "]")
		fmt.Print("Available snapshots:")
		PrintSnapshots(record.Snapshots)
		return
	}

	fmt.Print("Found ", len(matches), " match(es)")

}
