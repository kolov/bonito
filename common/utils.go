package common

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"io/ioutil"
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
		}

	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		log.Fatal("NewRequest: ", err)
		return nil, err
	}

	token, err := getToken();
	if err != nil {
		log.Fatal("Token: ", err)
		return nil, err
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

func getToken() (string, error) {
	//token := os.Getenv("DO_TOKEN_BONITO")
	token := AuthToken
	if token == "" {
		fmt.Println("Please, put your Digital Ocean Authorizarion token in an env var named DO_TOKEN_BONITO " +
			"and try again.")
		fmt.Println("See https://cloud.digitalocean.com/settings/api/tokens for more information")
		return "", errors.New("Fatal: no auth token")
	}
	return token, nil
}

func Query(url string, result interface{}) error {

	resp, err := sendRequest(url, "GET", nil)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if ( !strings.HasPrefix(resp.Status, "2")) {
		return errors.New(resp.Status)
	}

	return json.NewDecoder(resp.Body).Decode(result)
}

func Post(url string, body interface{}) (*http.Response, error) {
	return sendRequest(url, "POST", body)
}

func PostAndParse(url string, body interface{}, result interface{}) error {

	resp, err := Post(url, body)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if ( !strings.HasPrefix(resp.Status, "2")) {
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
		return errors.New(fmt.Sprintf("%s: %s", resp.Status, b))
	}
	err = json.NewDecoder(resp.Body).Decode(result)
	return err

}

func PrintError(err error) error {
	fmt.Println("Error:", err)
	return err
}

func Timeid() string {
	return time.Now().Format("2-1-2006-15-04")
}

func ResponseOK(response *http.Response) bool {
	return strings.HasPrefix(response.Status, "2")
}