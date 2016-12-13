package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"regexp"
	"github.com/kolov/bonito/common"
)

func CmdUp(c *cli.Context) {
	if common.SnapshotTemplate == "" && common.SnapshotId == "" {
		fmt.Println("Either --template or --snapshotid must be provided.")
		cli.ShowCommandHelp(c, "up")
		return
	}
	if common.SnapshotTemplate != "" && common.SnapshotId != "" {
		fmt.Println("Only one of --template and --snapshotid must be provided.")
		cli.ShowCommandHelp(c, "up")
		return
	}

	record, err := querySnapshots()
	if err != nil {
		fmt.Println(err)
		return
	}

	if common.SnapshotTemplate != "" {
		if common.Verbose {
			fmt.Println("Loking up snapshots matching template [", common.SnapshotTemplate, "]")
		}

		var reName = regexp.MustCompile(common.SnapshotTemplate)

		matches := []Snapshot{}
		for _, snapshot := range record.Snapshots {
			if reName.MatchString(snapshot.Name) {
				matches = append(matches, snapshot)
			}
		}
		if len(matches) == 0 {
			fmt.Println("No snapshots found matching [", common.SnapshotTemplate, "]", "Available snapshots:")
			printSnapshots(record.Snapshots)
			return
		}

		fmt.Println("Found ", len(matches), " match(es):")
		printSnapshots(matches)

		selected := record.Snapshots[0]

		if len(matches) >= 1 {
			fmt.Println("More than one snapshot matches")
			if !common.UseLatestSnapshot {
				fmt.Println("Specify exact name or use --latest")
				return
			}
			fmt.Println("Will use the latest ")
			for _, snapshot := range matches {
				if snapshot.CreatedAt.After(selected.CreatedAt) {
					selected = snapshot
				}
			}
		} else if len(matches) == 0 {
			fmt.Println("No matching snapshots found")
			return
		}
		fmt.Println("Will start droplet from: ", selected)

		startDropletFromSnapshot(selected)
	}

	if common.SnapshotId != "" {
		matches := []Snapshot{}
		for _, snapshot := range record.Snapshots {
			if common.SnapshotId == snapshot.Id {
				matches = append(matches, snapshot)
			}
		}
		if len(matches) != 1 {
			fmt.Print("Expected 1 snapsot with Id ", common.SnapshotId, ", found " + string(len(matches)))
			fmt.Print("Available snapshots:")
			printSnapshots(record.Snapshots)
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
			common.PrintError(err)
			return
		}
		for _, key := range keys {
			if common.Keys == key.Name {
				keyIds = append(keyIds, key.Id)
			}
		}
		if len(keys) == 0 {
			fmt.Println("ccould not find keys named [", common.Keys, "]")
			return
		}

	}

	name := common.DropletName;
	if name == "" {
		name = "bonito-" + common.Timeid()
		if common.Verbose {
			fmt.Println("No droplet name provided, will use generated:[", name, "]")
		}
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

	CmdStartDroplet(body)
}
/**
Select image to meet min size reqquirements. Now it just returns 2gb
 */
func selectSize(minDiskSize int) string {
	return "2gb"
}

func CmdStartDroplet(cmd StartDroplet) {

	if common.Verbose {
		fmt.Println("Bonito will start the following droplet: ", cmd)
	}
	if !common.Force {
		fmt.Println("Are you sure? Type yes to continue or no to stop")
		if !common.Confirm() {
			return
		}
		fmt.Println("Proceeding... ")
	}
	cmdResp, err := startDroplet(cmd)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Dropet starting....")

	waitUntilStarted(cmdResp.Droplet.Id)
}