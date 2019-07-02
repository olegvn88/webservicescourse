package main

import (
	"encoding/json"
	"fmt"
)

// structurs can have meta info - tags
type User struct {
	ID       int `json:"user_id, string"`
	Username string
	Address  string `json:",omitempty"` // don't write to result during packaging
	Company  string `json:"-"`          // don't necessary to serialize and deserialize
}

func main() {
	u := User{
		ID:       30,
		Username: "onester",
		Address:  "",
		Company:  "SoftServe",
	}
	result, _ := json.Marshal(u)
	fmt.Printf("json string %s\n", string(result))
}
