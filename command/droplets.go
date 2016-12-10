package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
	"github.com/kolov/sardine/common"
)

type Droplet struct {
	Id           int      `json:"id"`
	Name         string      `json:"name"`
	Memory       int      `json:"memory"`
	Vcpus        int      `json:"vcpus"`
	Locked       bool      `json:"locked"`
	Status       string      `json:"status"`
	Kernel       struct {
			     Id      int      `json:"id"`
			     Name    string      `json:"name"`
			     Version string      `json:"version"`
		     } `json:"kernel"`
	created_at   string      `json:"created_at"`
	Backup_ids   []int      `json:"backup_ids"`
	Snapshot_ids []int      `json:"snapshot_ids"`
	Image        struct {
			     Id             int      `json:"id"`
			     Name           string      `json:"name"`
			     Distribution   string      `json:"distribution"`
			     Slug           string      `json:"slug"`
			     Public         bool      `json:"public"`
			     Regions        []string      `json:"regions"`
			     Created_at     string      `json:"created_at"`
			     Min_disk_size  int      `json:"min_disk_size"`
			     Itype          string      `json:"type"`
			     Size_gigabytes float32      `json:"size_gigabytes"`
		     }  `json:"image"`
}

type DropletsList struct {
	Droplets [] Droplet   `json:"droplets"`
}

func CmdListDroplets(c *cli.Context) {

	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets?page=1&per_page=100")

	var record DropletsList

	common.Query(url, &record)

	if len(record.Droplets) != 0 {
		for i, v := range record.Droplets {
			fmt.Println(i + 1, strings.Join(
				[]string{" [", v.Name, "] created from image [", v.Image.Name, "]"}, ""))
		}
	} else {
		fmt.Println("No active droplets")
	}

}