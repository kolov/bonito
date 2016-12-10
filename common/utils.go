package common

import (
	"log"
	"os"
	"fmt"
	"net/http"
	"encoding/json"
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
func Query(url string, result *DropletsList) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	token := os.Getenv("DO_TOKEN_SARDINE")
	if token == "" {
		fmt.Println("Please, put your Digital Ocean Authorizarion token in an env var named DO_TOKEN_SARDINE " +
			"and try again.")
		fmt.Println("See https://cloud.digitalocean.com/settings/api/tokens for more information")
		return
	}
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending req: ", err)
		return
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		log.Println(err)
		return
	}
}

