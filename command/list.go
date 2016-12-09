package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//	"net/url"
	"github.com/codegangsta/cli"
	//	"reflect"
	"os"
)

type Droplet struct {
	id           int      `json:"id"`
	name         string      `json:"name"`
	memory       int      `json:"memory"`
	vcpus        int      `json:"vcpus"`
	locked       bool      `json:"locked"`
	status       string      `json:"status"`
	kernel       struct {
			     id      int      `json:"id"`
			     name    string      `json:"name"`
			     version string      `json:"version"`
		     } `json:"kernel"`
	created_at   string      `json:"created_at"`
	backup_ids   []string      `json:"backup_ids"`
	snapshot_ids []string      `json:"snapshot_ids"`
	image        struct {
			     id             int      `json:"id"`
			     name           int      `json:"name"`
			     distribution   int      `json:"distribution"`
			     slug           int      `json:"slug"`
			     public         bool      `json:"public"`
			     regions        []string      `json:"regions"`
			     created_at     string      `json:"created_at"`
			     min_disk_size  int      `json:"min_disk_size"`
			     itype          string      `json:"type"`
			     size_gigabytes int      `json:"size_gigabytes"`
		     }  `json:"image"`
}

type ListDropletsResponse struct {
	droplets [] Droplet   `json:"droplets"`
}

func CmdList(c *cli.Context) {

	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets?page=1&per_page=100")

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

	req.Header.Add("Authorization", "Bearer " + os.Getenv("DO_TOKEN_SARDINE"))
	req.Header.Add("Content-Type", "application/json")


	// For control over HTTP client headers,
	// redirect policy, and other settings,
	// create a Client
	// A Client is an HTTP client
	client := &http.Client{}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return
	}

	fmt.Print("Response: ")
	fmt.Println(resp)
	fmt.Print("Body: ")
	fmt.Println(resp.Body)
	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record ListDropletsResponse

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("record = ", record)
}
