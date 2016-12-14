package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/kolov/bonito/common"
	"encoding/json"
	"time"
)

type Network struct {
	IpAddress string    `json:"ip_address"`
	Netmask   string    `json:"netmask"`
	Gateway   string    `json:"gateway"`
	Type      string    `json:"type"`
}

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
	CreatedAt    string `json:"created_at"`

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
	Networks     struct {
			     V4 []Network `json:"v4"`
			     V6 []Network `json:"v6"`
		     } `json:"networks"`
	Features     []string `json:"features"`
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
	Action struct {
		       Id           int    `json:"id"`
		       Status       string    `json:"status"`
		       Type         string    `json:"type"`
		       StartedAt    string    `json:"started_at"`
		       CompletedAt  string `json:"completed_at"`
		       ResourceId   int      `json:"resource_id"`
		       ResourceType string      `json:"resource_type"`
		       // "region":{"name":"Frankfurt 1","slug":"fra1","sizes":["512mb","1gb","2gb","4gb","8gb","16gb","m-16gb","32gb","m-32gb","48gb","m-64gb","64gb","m-128gb","m-224gb"],"features":["private_networking","backups","ipv6","metadata","install_agent","storage"],"available":true},"region_slug":"fra1"}}
	       }  `json:"action"`
}

func (ar ActionResponse)String() string {
	barr, _ := json.Marshal(ar)
	return string(barr)
}

func (sd StartDroplet) String() string {
	barr, _ := json.Marshal(sd)
	return string(barr)
}

func queryAction(actionId int) (ActionResponse, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/actions/%d", actionId)

	var record ActionResponse
	err := common.Query(url, &record)
	return record, err

}
func queryDroplet(id int) (Droplet, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%d", id)

	var record DropletsResponse
	err := common.Query(url, &record)
	return record.Droplet, err

}
func queryDroplets() ([]Droplet, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets?page=1&per_page=100")

	var record DropletsList

	err := common.Query(url, &record)
	if err == nil {
		return record.Droplets, err
	} else {
		return nil, err
	}
}

func CmdListDroplets(_ *cli.Context) {

	droplets, err := queryDroplets()
	if err != nil {
		fmt.Println("error", err)
		return
	}
	printDroplets(droplets)
}

func printDroplets(droplets []Droplet) {
	if len(droplets) != 0 {
		for i, v := range droplets {
			fmt.Printf("%d. [%s] created from image [%s] statuse=[%s] ip=%s\n",
				i + 1, v.Name, v.Image.Name, v.Status, v.Networks.V4[0].IpAddress)
		}
	} else {
		fmt.Println("No active droplets")
	}
}

func waitUntilStarted(id int) {
	ticker := time.NewTicker(11 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			fmt.Printf("Waiting for droplet %d to become [active]\n", id)
			droplet, err := queryDroplet(id)
			if ( err != nil) {
				fmt.Println(err)
			} else {
				fmt.Printf("Droplet [id=%d, name= %s] has status [%s]\n",
					droplet.Id, droplet.Name, droplet.Status)
				if droplet.Status == "active" {
					fmt.Println("Droplet started successfully:")
					fmt.Printf("[%s] created from image [%s] statuse=[%s] ip=%s\n",
						droplet.Name, droplet.Image.Name, droplet.Status, droplet.Networks.V4[0].IpAddress)
					close(quit)
				}
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func startDroplet(body StartDroplet) (DropletsResponse, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets")
	var dropletResponse DropletsResponse
	err := common.PostAndParse(url, body, &dropletResponse)
	return dropletResponse, err

}

func executeDropletCommand(dropletId int, command DropletCommand) (ActionResponse, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%d/actions", dropletId);
	resp := ActionResponse{}
	err := common.PostAndParse(url, command, &resp)
	return resp, err
}
func shutdownDroplet(dropletId int) (ActionResponse, error) {
	return executeDropletCommand(dropletId, DropletCommand{"shutdown"})
}

func snapshotDroplet(dropletId int, name string) (ActionResponse, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%d/actions", dropletId);
	resp := ActionResponse{}
	err := common.PostAndParse(url, NamedDropletCommand{"snapshot", name}, &resp)
	return resp, err
}

func destroyDroplet(dropletId int) (ActionResponse, error) {
	return executeDropletCommand(dropletId, DropletCommand{"destroy"})
}

