package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"regexp"
)

var SnapshotTemplate string
var SnapshotId string
var UseLatestSnapshot bool

func CmdUp(c *cli.Context) {
	if SnapshotTemplate == "" && SnapshotId == "" {
		fmt.Println("Either --template or --snapshot must be provided.")
		cli.ShowCommandHelp(c, "up")
		return
	}
	if SnapshotTemplate != "" && SnapshotId != "" {
		fmt.Println("Only one of --template and --snapshot must be provided.")
		cli.ShowCommandHelp(c, "up")
		return
	}

	record := QuerySnapshots()

	if SnapshotTemplate != "" {
		fmt.Println("Starting up from snapshot matching ", SnapshotTemplate)

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

		fmt.Println("Found ", len(matches), " match(es)")

		if len(matches) != 1 {
			fmt.Println("--latest not supported yet. Please, specify the full name")
		}

		startDropletFromSnapshot(matches[0])
	}

	if SnapshotId != "" {
		matches := []Snapshot{}
		for _, snapshot := range record.Snapshots {
			if SnapshotId == snapshot.Id {
				matches = append(matches, snapshot)
			}
		}
		if len(matches) != 1 {
			fmt.Print("Expected 1 snapsot with Id ", SnapshotId, ", found " + string(len(matches)))
			fmt.Print("Available snapshots:")
			PrintSnapshots(record.Snapshots)
			return
		}
		startDropletFromSnapshot(SnapshotId)
	}

}

func startDropletFromSnapshot(snapshot Snapshot) {
	region := snapshot.Regions[0]
	startDroplet(snapshot.Id, region)
}
