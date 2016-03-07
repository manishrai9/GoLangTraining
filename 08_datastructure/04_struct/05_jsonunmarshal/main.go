package main

import (
	"encoding/json"
	"fmt"
)

//Person type
//can use tags as well
type Person struct {
	First string //as starts with Capital, it is exported, same goes for Last and Age
	Last  string
	Age   int
	//Age int `json:"Wisdom score"`
}

func main() {
	var p1 Person
	fmt.Println(p1.First, p1.Last, p1.Age)
	bs := []byte(`{"First":"Pratima","Last":"Rai","Age":30}`)
	//bs := []byte(`{"First":"Pratima","Last":"Rai","Wisdom score":30}`)

	json.Unmarshal(bs, &p1)
	fmt.Println("************After Unmarshal***********")
	fmt.Println(p1.First, p1.Last, p1.Age)
	fmt.Printf("%T\n", p1)
}
