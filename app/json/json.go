package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Phone    string
}

var jsonStr = `{"id": 716, "username": "onester", "phone": "063"}`

func main() {

	data := []byte(jsonStr)

	u := &User{}
	json.Unmarshal(data, u) // u - variable where we want to put result
	fmt.Printf("struct:\n\t%#v\n\n", u)

	u.Phone = "2345532" // change phone number and create json
	result, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}
	fmt.Printf("json string:\n\t%s\n", string(result))
}
