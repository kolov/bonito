package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
	"github.com/kolov/bonito/common"
	"time"
	"strconv"
)

type Snapshot struct {
	Id            string      `json:"id"`
	Name          string      `json:"name"`
	Regions       []string      `json:"regions"`

	CreatedAt     time.Time      `json:"created_at"`
	ResourceId    string      `json:"resource_id"`
	ResourceType  string      `json:"resource_type"`

	MinDISKSize   int      `json:"min_disk_size"`
	SizeGigabytes float32      `json:"size_gigabytes"`
}

type SnapshotList struct {
	Snapshots [] Snapshot   `json:"snapshots"`
}

func querySnapshots() (SnapshotList, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/snapshots?page=1&per_page=100")
	var record SnapshotList
	err := common.Query(url, &record)
	return record, err
}

func CmdListSnapshots(_ *cli.Context) {

	record, err := querySnapshots()

	if err != nil {
		fmt.Println(err)
		return
	}
	if len(record.Snapshots) != 0 {
		PrintSnapshots(record.Snapshots)
	} else {
		fmt.Println("No snapshots")
	}

}

func PrintSnapshots(snapshots []Snapshot) {
	for i, v := range snapshots {
		fmt.Println(i + 1, toString(v))
	}
}
func toString(v Snapshot) string {
	return strings.Join(
		[]string{" [", v.Name, "] created at [",
			v.CreatedAt.Format("2/1/2006 15:04"),
			"], regions=[", strings.Join(v.Regions, ","),
			"], mindisk=[", strconv.Itoa(v.MinDISKSize), "]"}, "")
}