package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"regexp"
	"github.com/kolov/bonito/common"
)

func CmdUp(c *cli.Context) {
	if common.SnapshotTemplate == "" {
		fmt.Println("--snapshot must be provided.")
		cli.ShowCommandHelp(c, "up")
		return
	}

	record, err := queryAllSnapshots()
	if err != nil {
		fmt.Println(err)
		return
	}

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

	selected := record.Snapshots[0]

	if len(matches) > 1 {
		fmt.Printf("%d snapshot(s) match the given snapshot name [%s]:\n", len(matches), common.SnapshotTemplate)
		printSnapshots(matches)
		if !common.UseLatestSnapshot {
			fmt.Println("Specify exact name or use --latest")
			return
		} else {
			fmt.Println("Will use the latest")
			for _, snapshot := range matches {
				if snapshot.CreatedAt.After(selected.CreatedAt) {
					selected = snapshot
				}
			}
		}
	}


	startDropletFromSnapshot(selected)

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
			fmt.Println("Could not find keys named [", common.Keys, "]")
			return
		} else {
			fmt.Printf("Found %d keys matching [%s]", len(keys), common.Keys)
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

	fmt.Println("Selected snapshot", snapshot)
	fmt.Println("Droplet create command: ", body)
	startDropletFromCommand(body)
}
/**
Select image to meet min size requirements. Now it just returns 2gb
 */
func selectSize(minDiskSize int) string {
	if common.DropletSize != "" {
		return common.DropletSize
	} else {
		return "2gb"
	}
}

func startDropletFromCommand(cmd StartDroplet) {

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

	fmt.Println("Droplet starting....")

	waitUntilStarted(cmdResp.Droplet.Id)
}