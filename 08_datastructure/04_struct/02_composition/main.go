package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
}

type doublez struct {
	person
	tempbool bool
}

func (p person) getname() string {
	return p.fname + p.lname
}

func (d doublez) getname() string {
	return "My name is anthony gonsalvese!"
}

func main() {
	p1 := person{"Manish", "Rai", 31}
	p2 := doublez{
		person: person{"Pratima", "Rai", 30}, tempbool: true,
	}
	fmt.Println(p1.fname, p1.lname, p1.age)
	fmt.Println("Person name: ", p1.getname())
	fmt.Println("************************************")
	fmt.Println(p2.fname, p2.lname, p2.age)
	fmt.Println("2nd Person name: ", p2.getname())
}
