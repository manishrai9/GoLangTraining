package main

import (
	"encoding/json"
	"os"
)

//Person type
type Person struct {
	First       string //as starts with Capital, it is exported, same goes for Last and Age
	Last        string
	Age         int
	notExported int
}

func main() {
	p := Person{"Manish", "Rai", 31, 8}
	json.NewEncoder(os.Stdout).Encode(p)
}
