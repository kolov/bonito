package command

import (
	"github.com/codegangsta/cli"
	"fmt"
	"regexp"
	"github.com/kolov/bonito/common"
	"strconv"
	"io/ioutil"
)

func CmdShutdown(c *cli.Context) {
	fmt.Println("Shutting down called.")

	if SnapshotTemplate == "" && DropletName == "" {
		fmt.Println("Either --template or --name must be provided")
		return
	}

	if SnapshotTemplate != "" && DropletName != "" {
		fmt.Println("One of --template or --name must be provided")
		return
	}
	droplets, err := ListDroplets()
	if err != nil {
		common.PrintErrorAndExit(err)
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
	if DropletName != "" {
		fmt.Printf("Looking for droplets named [%s]\n", DropletName)

		matches := []Droplet{}
		for _, droplet := range droplets {
			if DropletName == droplet.Name {
				matches = append(matches, droplet)
			}
		}

		if len(matches) != 1 {
			fmt.Printf("Expected 1 droplet matching name [%s], found %d\n",
				DropletName, len(matches))
			return
		}
		matchingDroplet = matches[0]
	}
	if Verbose {
		fmt.Println("Will shutdown ", matchingDroplet)
	}
	shutdown(matchingDroplet)

}

func shutdown(droplet Droplet) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets/%d/actions", droplet.Id);

	fmt.Println("Will use url", url)
	resp, err := common.Post(url, DropletCommand{"shutdown"})
	if err != nil {
		common.PrintErrorAndExit(err)
		return
	}
	if resp.Status == "201 Created" {
		fmt.Println("Shuddown in progress: ")
	}
	barr, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(barr))

}
