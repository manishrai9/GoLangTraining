package main

import (
	"encoding/json"
	"fmt"
)

//Person type
type Person struct {
	First       string //as starts with Capital, it is exported, same goes for Last and Age
	Last        string
	Age         int
	notExported int //this will not be exported,so small case
}

func main() {
	p := Person{"Manish", "Rai", 31, 007}
	bs, _ := json.Marshal(p)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
