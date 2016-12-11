package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
)

func sendRequest(url string, method string, body interface{}) (*http.Response, error) {

	var req *http.Request
	var err error

	if body != nil {
		barr, err := json.Marshal(body)
		if err != nil {
			fmt.Println("error marshalling", err)
			return nil, errors.New("marshalling")
		} else {
			req, err = http.NewRequest(method, url, bytes.NewBuffer(barr))
			//fmt.Println("Marshalled body:", string(barr))
		}

	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	token := os.Getenv("DO_TOKEN_BONITO")
	if token == "" {
		fmt.Println("Please, put your Digital Ocean Authorizarion token in an env var named DO_TOKEN_BONITO " +
			"and try again.")
		fmt.Println("See https://cloud.digitalocean.com/settings/api/tokens for more information")
		return nil, errors.New("Fatal: no auth token")
	}
	req.Header.Add("Authorization", "Bearer " + token)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	//fmt.Println("About to send request", req)
	resp, err := client.Do(req)
	//fmt.Println("Got response: ", resp)
	if err != nil {
		log.Fatal("Error sending req: ", err)
		return resp, err
	}

	return resp, err
}

func Query(url string, result interface{}) error {

	resp, err := sendRequest(url, "GET", nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(result)
}

func Post(url string, body interface{}) (*http.Response, error) {
	return sendRequest(url, "POST", body)
}
