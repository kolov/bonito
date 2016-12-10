package command

import (
	"fmt"
	"log"
	"net/http"
	//	"net/url"
	"github.com/codegangsta/cli"
	//	"reflect"
	"os"
)




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
		log.Fatal("Error sending req: ", err)
		return
	}

	fmt.Print("Response: ")
	fmt.Println(resp)
	fmt.Print("Body: ")

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(resp.Body)
	//s := buf.String() // Does a complete copy of the bytes in the buffer.
	//fmt.Println(s)
	// Callers should close resp.Body
	// when done reading from it
	// Defer the closing of the body
	defer resp.Body.Close()

	// Fill the record with the data from the JSON
//	var record ListDropletsResponse

	// Use json.Decode for reading streams of JSON data
	//if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
	//	log.Println(err)
	//	return
	//}
	//
	//fmt.Println("record = ", record)
}
