package command

import (
	"github.com/codegangsta/cli"
	"fmt"
	"regexp"
	"github.com/kolov/bonito/common"
	"strconv"
	"io/ioutil"
	"time"
	"net/http"
	"encoding/json"
)


func CmdShutdown(c *cli.Context) {
	fmt.Println("Shutting down called.")

	if SnapshotTemplate == "" && common.DropletName == "" {
		fmt.Println("Either --template or --name must be provided")
		return
	}

	if SnapshotTemplate != "" && common.DropletName != "" {
		fmt.Println("One of --template or --name must be provided")
		return
	}
	droplets, err := QueryDroplets()
	if err != nil {
		common.PrintError(err)
		return
	}

	var matchingDroplet Droplet

	if SnapshotTemplate != "" {
		fmt.Println("Looking for droplets started from image matching [", SnapshotTemplate, "]")

		var reName = regexp.MustCompile(SnapshotTemplate)

		matches := []Droplet{}
		for _, droplet := range droplets {
			if reName.MatchString(droplet.Image.Name) {
				matches = append(matches, droplet)
			}
		}

		if len(matches) != 1 {
			fmt.Println("Expected 1 droplet matching image [", SnapshotTemplate, "], found", strconv.Itoa(len(matches)))
			return
		}
		matchingDroplet = matches[0]
	}
	if common.DropletName != "" {
		fmt.Printf("Looking for droplets named [%s]\n", common.DropletName)

		matches := []Droplet{}
		for _, droplet := range droplets {
			if common.DropletName == droplet.Name {
				matches = append(matches, droplet)
			}
		}

		if len(matches) != 1 {
			fmt.Printf("Expected 1 droplet matching name [%s], found %d\n",
				common.DropletName, len(matches))
			printDroplets(droplets)
			return
		}
		matchingDroplet = matches[0]
	}
	if common.Verbose {
		fmt.Println("Will shutdown ", matchingDroplet)
	}

	shutdown(matchingDroplet)

}

func shutdownDroplet(dropletId int) (*http.Response, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%d/actions", dropletId);
	return common.Post(url, DropletCommand{"shutdown"})
}

func snapshotDroplet(dropletId int, name string) (*http.Response, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%d/actions", dropletId);
	return common.Post(url, NamedDropletCommand{"snapshot", name})
}

func shutdown(droplet Droplet) {

	resp, err := shutdownDroplet(droplet.Id)

	if err != nil {
		common.PrintError(err)
		return
	}
	if resp.Status == "201 Created" {
		fmt.Println("Shuddown in progress. WWaiting for droplet to shutdown...")
		var snapshotBase = droplet.Image.Name
		waitShutdownAndSnapshot(droplet.Id, snapshotBase + "1")

	}
	barr, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(barr))

}

func waitShutdownAndSnapshot(id int, snapshotname string) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	var phase = "shutdown"
	for {
		select {
		case <-ticker.C:
			fmt.Println("tick!")
			if phase == "shutdown" {
				droplet, err := queryDroplet(id)
				if ( err != nil) {
					fmt.Println(err)
				} else {
					fmt.Printf("[%s] status [%s] id=%d\n", droplet.Name, droplet.Status, droplet.Id)
					if droplet.Status == "off" {
						phase = "asksnapshot"
						fmt.Println("Droplet shut down successfully. Starting snapshot")
						ticker.Stop()
					}
				}
			} else if phase == "asksnapshot" {
				fmt.Printf("requesting snapshot")
				resp, err := snapshotDroplet(id, snapshotname)
				if err != nil {
					var response ActionResponse
					err := json.NewDecoder(resp.Body).Decode(response)
					if err != nil {
						phase = "waitsnapshot"
					} else {
						common.PrintError(err)
					}
				}
			} else if phase == "waitsnapshot" {
				fmt.Println("Waitingsnapshot tofinish")
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}

}
