package main

import (
	"encoding/json"
	"fmt"
)

func square(x *float64) {
	*x = *x * *x
	fmt.Println(*x)
}
func main() {
	x := 1.5
	square(&x)
	parseJson(str)
}

var str = `{"Id": 1, "Name":"Oleg"}`

type Person struct {
	Name string
	Id   int32
}

func parseJson(jsonData string) {
	p := &Person{}
	jsonDat := []byte(jsonData)
	json.Unmarshal(jsonDat, p)
	fmt.Println(p.Id)
	fmt.Println(p.Name)

	var m map[string]int32
	m = make(map[string]int32)
	m["Map"] = 30
	d := m["Map"]
	fmt.Println(len(m))
	fmt.Println(d)

	pars := Person{"Oleg", 30}
	pars.SetName("Nikita")
	fmt.Println(pars.Name)

	var acc Account = Account{
		Person:      Person{"Oleg", 2},
		AccountName: "Account Name",
	}
	acc.SetName("Nesterov")
	fmt.Println(acc)
	s1 := newSlice([]int{1, 2})
	s1.Add(10)
	fmt.Println(s1.Count(), s1)
}

type newSlice []int

func (sl *newSlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *newSlice) Count() int {
	return len(*sl)

}

func PersonFunc(name string, id int32) []int {
	num := []int{2, 4, 5}

	return num
}

//изменяет структуру , для которой вызван
func (p *Person) SetName(name string) {
	// адрес на тип
	p.Name = name
}

//не изменяет структуру , для которой вызван
func (p Person) UpdateName(name string) {
	// p Person -передается копия этого значения
	p.Name = "New structure"
}

type Account struct {
	Person
	AccountName string
}
