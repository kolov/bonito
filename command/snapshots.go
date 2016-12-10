package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
	"github.com/kolov/sardine/common"
)

type Snapshot struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	Regions       []string      `json:"regions"`

	CreatedAt     string      `json:"created_at"`
	ResourceId    string      `json:"resource_id"`
	ResourceType  string      `json:"resource_type"`

	MinDISKSize   int      `json:"min_disk_size"`
	SizeGigabytes float32      `json:"size_gigabytes"`
}

type SnapshotList struct {
	Snapshots [] Snapshot   `json:"snapshots"`
}

func CmdSnapshots(c *cli.Context) {

	url := fmt.Sprintf("https://api.digitalocean.com/v2/snapshots?page=1&per_page=100")

	var record SnapshotList

	common.Query(url, &record)

	if len(record.Snapshots) != 0 {
		for i, v := range record.Snapshots {
			fmt.Println(i + 1, strings.Join(
				[]string{" [", v.Name, "] created at [", v.CreatedAt, "], type=", v.ResourceType}, ""))
		}
	} else {
		fmt.Println("No snapshots")
	}

	fmt.Println("Here's the rest ", record)

}