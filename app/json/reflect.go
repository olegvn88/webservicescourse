package main

import (
	"fmt"
	"reflect"
)

type User struct {
	ID int
	RealName string `unpack:"-"`
	Login string
	Flags int
}12

func PrintReflect(u interface{}) error {
	val := reflect.ValueOf(u).Elem()

	fmt.Printf("%T have %d fields:\n", u, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		fmt.Printf("\tname=%v. type=%v, value=%v, tag=`%v`\n",
			typeField.Name,
			typeField.Type.Kind(),
			valueField,
			typeField.Tag,
		)
	}

	return nil
}

func main(){
	u := &User {
		ID: 18,
		RealName: "onester",
		Flags: 30,
	}
	err := PrintReflect(u)
	if err != nil {
		panic(err)
	}
}
