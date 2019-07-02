package main

import (
	"encoding/json"
	"fmt"
)

var jsonString = `[
{"id": 18, "username": "onester", "phone": 3333},
{"id": "18", "address": "none", "company": "SoftServe"}
]`

func main() {
	data := []byte(jsonString)

	var user1 interface{}
	json.Unmarshal(data, &user1)
	fmt.Printf("unpacked in empty interface:\n%#v\n\n", user1)

	//return
	user2 := map[string]interface{}{
		"id":       18,
		"username": "oleg",
	}

	var user2i interface{} = user2

	result, _ := json.Marshal(user2i)

	fmt.Printf("json string from map: %#s\t\t", result)
}
