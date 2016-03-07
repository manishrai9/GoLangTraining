package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	age   int
}

func (p person) getname() string {
	return p.fname + p.lname
}

func main() {
	p1 := person{"Manish", "Rai", 31}
	fmt.Println(p1.fname, p1.lname, p1.age)
	fmt.Println("Person name: ", p1.getname())
}
