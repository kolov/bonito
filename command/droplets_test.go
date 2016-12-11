package command

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCmdDroplets(t *testing.T) {
	// Write your code here
}

func ExampleReverse() {
	b, _ := json.Marshal(StartDroplet{"sardine",
		"ams1",
		"2gb",
		"iid",
		nil,
		false,
		false,
		nil,
		false,
		nil,
		&[]string{"sardine"},
	})
	fmt.Println(string(b))
	// Output: {"name":"sardine","region":"ams1","size":"2gb","image":"iid","ssh_keys":null,"backups":false,"ipv6":false,"user_data":null,"private_networking":false,"volumes":null,"tags":["sardine"]}

}
