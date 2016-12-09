package command

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//	"net/url"
	"github.com/codegangsta/cli"
	//	"reflect"
)

type ListResponse struct {
	Valid bool   `json:"valid"`
}

func CmdList(c *cli.Context) {

	url := fmt.Sprintf("https://api.digitalocean.com/v2/droplets?page=1&per_page=100")

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return
	}

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

	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
	var record ListResponse

	// Use json.Decode for reading streams of JSON data
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}

	fmt.Println("Phone No. = ", record.Valid)
}
