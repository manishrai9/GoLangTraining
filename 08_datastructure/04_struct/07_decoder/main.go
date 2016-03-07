package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

//Person type
type Person struct {
	First       string //as starts with Capital, it is exported, same goes for Last and Age
	Last        string
	Age         int
	notExported int
}

func main() {

	var p Person
	rdr := strings.NewReader(`{"First":"Manish","Last":"Rai","Age":31}`)
	json.NewDecoder(rdr).Decode(&p)

	fmt.Println(p.First)
	fmt.Println(p.Last)
	fmt.Println(p.Age)
}
