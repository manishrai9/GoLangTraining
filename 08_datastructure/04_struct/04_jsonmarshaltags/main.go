package main

import (
	"encoding/json"
	"fmt"
)

//Person type
type Person struct {
	First string //as starts with Capital, it is exported
	Last  string `json:"-"`            //this will not be exported, see tags for more detail
	Age   int    `json:"wisdom-score"` //this will be exported with name wisdom-score
}

func main() {
	p := Person{"Manish", "Rai", 31}
	bs, _ := json.Marshal(p)
	fmt.Println(string(bs)) //{"First":"Manish","wisdom-score":31}
}
