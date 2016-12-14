package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
	"github.com/kolov/bonito/common"
	"time"
	"encoding/json"
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

func (s Snapshot)String() string {
	barr, _ := json.Marshal(s)
	return string(barr)
}

type SnapshotList struct {
	Snapshots [] Snapshot   `json:"snapshots"`
}

func queryAllSnapshots() (SnapshotList, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/snapshots?page=1&per_page=100")
	var record SnapshotList
	err := common.Query(url, &record)
	return record, err
}
func queryDropletSnapshots(dropletId int) (SnapshotList, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/$d/snapshots", dropletId)
	var record SnapshotList
	err := common.Query(url, &record)
	return record, err
}

func CmdListSnapshots(_ *cli.Context) {

	record, err := queryAllSnapshots()

	if err != nil {
		fmt.Println(err)
		return
	}
	if len(record.Snapshots) != 0 {
		printSnapshots(record.Snapshots)
	} else {
		fmt.Println("No snapshots")
	}

}

func printSnapshots(snapshots []Snapshot) {
	for i, v := range snapshots {
		fmt.Println(i + 1, toString(v))
	}
}
func toString(v Snapshot) string {
	return fmt.Sprintf("[%s, id=%s] created at [%s], regions=[%s], mindisk=[%d]",
		v.Name, v.Id,
		v.CreatedAt.Format("2/1/2006 15:04"),
		strings.Join(v.Regions, ","),
		v.MinDISKSize)
}