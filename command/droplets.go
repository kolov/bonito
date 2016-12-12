package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/kolov/bonito/common"
	"encoding/json"
	"time"
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
type DropletsResponse struct {
	Droplet Droplet `json:"droplet"`
}

type DropletCommand struct {
	Type string `json:"type"`
}
type NamedDropletCommand struct {
	Type string `json:"type"`
	Name string `json:"name"`
}

type StartDroplet struct {
	Name              string    `json:"name"`
	Region            string    `json:"region"`
	Size              string    `json:"size"`
	Image             string    `json:"image"`
	SshKeys           *[]int `json:"ssh_keys"`
	Backups           bool      `json:"backups"`
	Ipv6              bool      `json:"ipv6"`
	UserData          *string   `json:"user_data"`
	PrivateNetworking bool      `json:"private_networking"`
	Volumes           *[]string `json:"volumes"`
	Tags              *[]string `json:"tags"`
}

type ActionResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Type   string `json:"type"`
}

func (sd StartDroplet) String() string {
	barr, _ := json.Marshal(sd)
	return string(barr)
}

func queryDroplet(id int) (Droplet, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%d", id)

	var record DropletsResponse
	err := common.Query(url, &record)
	return record.Droplet, err

}
func QueryDroplets() ([]Droplet, error) {
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

	droplets, err := QueryDroplets()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	printDroplets(droplets)

}

func printDroplets(droplets []Droplet) {
	if len(droplets) != 0 {
		for i, v := range droplets {
			fmt.Printf("%d. [%s] created from image [%s] statuse=[%s] id=%d\n",
				i + 1, v.Name, v.Image.Name, v.Status, v.Id)
		}
	} else {
		fmt.Println("No active droplets")
	}
}

func startDroplet(body StartDroplet) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets")

	if common.Verbose {
		fmt.Println("Bonito will start the following droplet: ", body)
	}
	if !common.Force {
		fmt.Println("Are you sure? Type yes to continue or no to stop")
		if !common.Confirm() {
			return
		}
		fmt.Println("Proceeding... ")
	}

	var dropletResponse DropletsResponse

	err := common.PostAndParse(url, body, &dropletResponse)

	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("Success")

	waitUntilStarted(dropletResponse.Droplet.Id)
}

func waitUntilStarted(id int) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Checking status of droplet %d", id)
			droplet, err := queryDroplet(id)
			if ( err != nil) {
				fmt.Println(err)
			} else {
				fmt.Printf("Droplet [id=%d, name= %s] has status [%s]\n",
					droplet.Id, droplet.Name, droplet.Status)
				if droplet.Status == "active" {
					fmt.Println("Droplet started successfulyl")
					ticker.Stop()
				}
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}
