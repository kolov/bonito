package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/kolov/bonito/common"
	"io/ioutil"
	"strings"
	"encoding/json"
)

type Droplet struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Memory       int    `json:"memory"`
	Vcpus        int    `json:"vcpus"`
	Locked       bool   `json:"locked"`
	Status       string `json:"status"`
	Kernel       struct {
			     Id      int    `json:"id"`
			     Name    string `json:"name"`
			     Version string `json:"version"`
		     } `json:"kernel"`
	created_at   string `json:"created_at"`
	Backup_ids   []int  `json:"backup_ids"`
	Snapshot_ids []int  `json:"snapshot_ids"`
	Image        struct {
			     Id             int      `json:"id"`
			     Name           string   `json:"name"`
			     Distribution   string   `json:"distribution"`
			     Slug           string   `json:"slug"`
			     Public         bool     `json:"public"`
			     Regions        []string `json:"regions"`
			     Created_at     string   `json:"created_at"`
			     Min_disk_size  int      `json:"min_disk_size"`
			     Itype          string   `json:"type"`
			     Size_gigabytes float32  `json:"size_gigabytes"`
		     } `json:"image"`
}

func (d Droplet)String() string {
	barr, _ := json.Marshal(d)
	return string(barr)
}

type DropletsList struct {
	Droplets []Droplet `json:"droplets"`
}

type DropletCommand struct {
	Type string `json:"type"`
}

type StartDroplet struct {
	Name              string    `json:"name"`
	Region            string    `json:"region"`
	Size              string    `json:"size"`
	Image             string    `json:"image"`
	SshKeys           *[]string `json:"ssh_keys"`
	Backups           bool      `json:"backups"`
	Ipv6              bool      `json:"ipv6"`
	UserData          *string   `json:"user_data"`
	PrivateNetworking bool      `json:"private_networking"`
	Volumes           *[]string `json:"volumes"`
	Tags              *[]string `json:"tags"`
}

func (sd StartDroplet) String() string {
	barr, _ := json.Marshal(sd)
	return string(barr)
}

func ListDroplets() ([]Droplet, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets?page=1&per_page=100")

	var record DropletsList

	err := common.Query(url, &record)
	if err == nil {
		return record.Droplets, err
	} else {
		return nil, err
	}
}
func CmdListDroplets(c *cli.Context) {

	droplets, err := ListDroplets()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	if len(droplets) != 0 {
		for i, v := range droplets {
			fmt.Println(i + 1, strings.Join(
				[]string{" [", v.Name, "] created from image [", v.Image.Name, "]"}, ""))
		}
	} else {
		fmt.Println("No active droplets")
	}

}

func startDroplet(body StartDroplet) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets")

	if Verbose {
		fmt.Println("Bonito will start the following droplet: ", body)
	}
	if !Force {
		fmt.Println("Are you sure? Type yes to continue or no to stop")
		if !common.Confirm() {
			return
		}
		fmt.Println("Proceeding... ")
	}
	resp, err := common.Post(url, body)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	htmlData, err := ioutil.ReadAll(resp.Body)
	if !strings.HasPrefix(resp.Status, "2") {
		fmt.Println("Error", resp.Status)
	} else {
		fmt.Println("Success")
	}
	fmt.Println("Response", string(htmlData))
}
