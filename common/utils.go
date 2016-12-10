package common

import (
	"log"
	"os"
	"fmt"
	"net/http"
	"encoding/json"
)

func Query(url string, result interface{}) {
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

