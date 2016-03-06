package main

import "fmt"

func main() {
	//2 way a map can be created
	//1: initialise to empty map
	m := map[string]int{}
	//2: make an empty map
	//m1 := make(map[string]int)
	m["manish"] = 1
	m["pratima"] = 2
	m["random"] = 3

	//fmt.Println("map is: ", m)
	fmt.Println("------------------------")
	fmt.Println("Iterating through map")

	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	fmt.Println("------------------------")
	//fmt.Println("m1 is: ", m1)

	val, ok := m["manish"]

	fmt.Println("value is: ", val, " value present?", ok)
	fmt.Println("------------------------")
	//to delete from map
	delete(m, "random")

	fmt.Println("map after deleting random key is: ", m)
	fmt.Println("------------------------")
	//declare and initialise in same line with following syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("another map:", n)
	fmt.Println("------------------------")
}
