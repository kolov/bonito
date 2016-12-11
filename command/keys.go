package command

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/kolov/bonito/common"
	"encoding/json"
)

type Key struct {
	Id          int    `json:"id"`
	Fingerprint string `json:"fingerprint"`
	PublicKey   string    `json:"public_key"`
	Name        string    `json:"name"`
}

func (k Key)String() string {
	barr, _ := json.Marshal(k)
	return string(barr)
}

type KeysList struct {
	Keys []Key `json:"ssh_keys"`
}

func ListKeys() ([]Key, error) {
	url := fmt.Sprintf("https://api.digitalocean.com/v2/account/keys")

	var record KeysList

	err := common.Query(url, &record)
	if err == nil {
		return record.Keys, err
	} else {
		return nil, err
	}
}
func CmdListKeys(c *cli.Context) {

	keys, err := ListKeys()
	if err != nil {
		fmt.Println("error", err)
		return
	}

	if len(keys) != 0 {
		for i, v := range keys {
			fmt.Printf("%d. [%s]\n", i + 1, v.Name)
		}
	} else {
		fmt.Println("No SSH Keys")
	}

}
