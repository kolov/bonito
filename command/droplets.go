package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"strings"
	"github.com/kolov/sardine/common"
)



func CmdDroplets(c *cli.Context) {

	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets?page=1&per_page=100")

	var record common.DropletsList

	common.Query(url, &record)

	if len(record.Droplets) != 0 {
		for i, v := range record.Droplets {
			fmt.Println(i + 1, strings.Join(
				[]string{" [", v.Name, "] created from image [", v.Image.Name, "]"}, ""))
		}
	} else {
		fmt.Println("No active droplets")
	}

	fmt.Println("Here's the rest ", record)

}