package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	pers := Person{1, "Oleg"}
	//pers := new(Person) // retuns pointer to a structure
	pers.SetName("Oleg Nesterov")
	//(&pers).SetName("Oleg Nesterov")
	//fmt.Printf("update person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "onester",
		Person: Person{
			Id:   2,
			Name: "Oleg Nesterov",
		},
	}
	acc.Person.SetName("Test")

	//fmt.Printf("%#v \n", acc)

	sl := MySlice([]int{1, 2})
	sl.add(3)
	fmt.Println(sl.Count(), sl)
}
