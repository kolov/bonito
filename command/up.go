package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"regexp"
	"time"
	"github.com/kolov/bonito/common"
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

	record, err := QuerySnapshots()
	if err != nil {
		fmt.Println(err)
		return
	}

	if SnapshotTemplate != "" {
		if common.Verbose {
			fmt.Println("Loking up snapshots matching template [", SnapshotTemplate, "]")
		}

		var reName = regexp.MustCompile(SnapshotTemplate)

		matches := []Snapshot{}
		for _, snapshot := range record.Snapshots {
			if reName.MatchString(snapshot.Name) {
				matches = append(matches, snapshot)
			}
		}
		if len(matches) == 0 {
			fmt.Println("No snapshots found matching [", SnapshotTemplate, "]", "Available snapshots:")
			PrintSnapshots(record.Snapshots)
			return
		}

		fmt.Println("Found ", len(matches), " match(es):")
		if common.Verbose {
			PrintSnapshots(matches)
		}

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
		startDropletFromSnapshot(matches[0])
	}

}

func startDropletFromSnapshot(snapshot Snapshot) {

	region := snapshot.Regions[0]
	size := selectSize(snapshot.MinDISKSize)

	var keyIds []int = nil
	if common.Keys != "" {
		//split := strings.Split(Keys, ",")
		keys, err := ListKeys()
		if err != nil {
			common.PrintErrorAndExit(err)
			return
		}
		for _, key := range keys {
			if common.Keys == key.Name {
				keyIds = append(keyIds, key.Id)
			}
		}
	}

	name := common.DropletName;
	if name == "" {
		name = "bonito-" + time.Now().Format("2-1-2006-15-04")
	}

	body := StartDroplet{
		name,
		region,
		size,
		snapshot.Id,
		&keyIds,
		false,
		false,
		nil,
		false,
		nil,
		&[]string{"bonito"}}

	startDroplet(body)
}
/**
Select image to meet min size reqquirements. Now it just returns 2gb
 */
func selectSize(minDiskSize int) string {
	return "2gb"
}
