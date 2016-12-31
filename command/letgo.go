package command

import (
	"github.com/codegangsta/cli"
	"fmt"
	"github.com/kolov/bonito/common"
	"time"
)

func CmdKill(c *cli.Context) {
	common.NoSnapshot = true
	CmdLetgo(c)
}

func CmdLetgo(_ *cli.Context) {
	if common.DropletName == "" {
		fmt.Println("--name must be provided")
		return
	}

	fmt.Printf("Looking for droplets named [%s]\n", common.DropletName)
	droplets, err := queryDroplets()
	if err != nil {
		common.PrintError(err)
		return
	}

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

	if !common.Silent {
		fmt.Println("Will shutdown ", matches[0])
	}

	letgo(matches[0])
}

func letgo(droplet Droplet) {

	if common.NoSnapshot {
		if !common.Force {
			fmt.Println("Are you sure you want do destroy this droplet without snapshot? Type yes to continue or no to stop")
			if !common.Confirm() {
				return
			}
			fmt.Println("Proceeding... ")
		}

	}

	snapshotName := common.SnapshotName
	if snapshotName == "" {
		snapshotName = nextSnapshotName(droplet)
		if !common.Silent {
			fmt.Println("No snapshot name was provide. Wil use [", snapshotName, "]")
		}
	}
	if !common.NoSnapshot {
		_, err := shutdownDroplet(droplet.Id)
		if err != nil {
			common.PrintError(err)
			return
		}
		fmt.Println("Shutdown in progress. Waiting for droplet to shutdown...")
		waitShutdown(droplet.Id)
		fmt.Println("Droplet shut down successfully. Requesting snapshot...")

		actionResp, err1 := snapshotDroplet(droplet.Id, snapshotName)
		if err1 != nil {
			common.PrintError(err1)
			return
		}
		fmt.Println("Snapshot requested. This can take quite a while.")
		if !common.Silent {
			fmt.Println("Snapshot action response", actionResp)
		}
		waitSnapshot(actionResp.Action.Id)
		fmt.Println("Snapshot taken successfully. Destoying starts.")
	}
	destroyAndReport(droplet)
}

func destroyAndReport(droplet Droplet) {
	actionResp, err := destroyDroplet(droplet.Id)
	if err != nil {
		common.PrintError(err)
		return
	}
	fmt.Println("Destroy, requested, got this response", actionResp);
}

func waitShutdown(id int) {
	start := time.Now().Unix()
	ticker := time.NewTicker(5 * time.Second)
	quit := make(chan struct{})

	for {
		select {
		case <-ticker.C:

			droplet, err := queryDroplet(id)
			if ( err != nil) {
				fmt.Println(err)
			} else {
				fmt.Printf("Droplet [id=%d name=%s] status [%s] after %d sec of waiting...\n",
					droplet.Id, droplet.Name, droplet.Status, time.Now().Unix() - start)
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
	start := time.Now().Unix()
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
				fmt.Printf("Action %d not completed in %d sec. Waiting...\n",
					actionId, time.Now().Unix() - start)
			}
		case <-quit:
			ticker.Stop()
			return
		}
	}

}

func nextSnapshotName(droplet Droplet) string {
	return droplet.Name + "-" + common.Timeid();
}