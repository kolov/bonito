package command

import (
	"github.com/codegangsta/cli"
	"fmt"
	"regexp"
	"github.com/kolov/bonito/common"
	"strconv"
	"time"
)

func CmdShutdown(_ *cli.Context) {
	if common.SnapshotTemplate == "" && common.DropletName == "" {
		fmt.Println("Either --template or --name must be provided")
		return
	}

	if common.SnapshotTemplate != "" && common.DropletName != "" {
		fmt.Println("One of --template or --name must be provided")
		return
	}
	droplets, err := queryDroplets()
	if err != nil {
		common.PrintError(err)
		return
	}

	var matchingDroplet Droplet

	if common.SnapshotTemplate != "" {
		fmt.Println("Looking for droplets started from image matching [", common.SnapshotTemplate, "]")

		var reName = regexp.MustCompile(common.SnapshotTemplate)

		matches := []Droplet{}
		for _, droplet := range droplets {
			if reName.MatchString(droplet.Image.Name) {
				matches = append(matches, droplet)
			}
		}

		if len(matches) != 1 {
			fmt.Println("Expected 1 droplet matching image [", common.SnapshotTemplate, "], found", strconv.Itoa(len(matches)))
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

func shutdown(droplet Droplet) {

	resp, err := shutdownDroplet(droplet.Id)

	if err != nil || !common.ResponseOK(resp) {
		common.PrintError(err)
		return
	}
	fmt.Println("Shutdown in progress. WWaiting for droplet to shutdown...")
	waitShutdown(droplet.Id)
	fmt.Println("Droplet shut down successfully. Starting snapshot")
	var snapshotBase = droplet.Image.Name
	actionResp, err1 := snapshotDroplet(droplet.Id, snapshotBase + "1")
	if err1 != nil {
		common.PrintError(err1)
		return
	}
	fmt.Println("Waiting for snaphot to finish. THis can take quite a while.")
	fmt.Println(actionResp)
	waitSnapshot(actionResp.Action.Id)
	fmt.Println("Snapshot taken successfully. Starting delete")
	destroyDroplet(droplet.Id)

}

func waitShutdown(id int) {
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:

			droplet, err := queryDroplet(id)
			if ( err != nil) {
				fmt.Println(err)
			} else {
				fmt.Printf("Droplet [id=%d name=%s] status [%s] \n",
					droplet.Id, droplet.Name, droplet.Status)
				if droplet.Status == "off" {
					close(quit)
				}
			}

		case <-quit:
			ticker.Stop()
			return
		}
	}

}

func waitSnapshot(actionId int) {
	ticker := time.NewTicker(20 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:
			aresp, err := queryAction(actionId)
			if err != nil {
				common.PrintError(err)
			} else {
				if aresp.Action.CompletedAt != "" {
					close(quit)
				}
				fmt.Sprint("Action %s not completed yet. Waiting...", actionId)
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}

}
