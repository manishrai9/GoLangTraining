package main

import "fmt"

func main() {
	slice1 := []string{}        //shorthand style, will need to use append
	var slice2 []string         //var style, will need to use append(always)
	slice3 := make([]string, 4) //make style, best way to create slice , can be used with indexes, setting length and capacity

	//slice1 = append(slice1, "Manish")
	//slice2 = append(slice2, "Manish")
	//slice3[0] = "Manish"

	fmt.Println("shorthand style nil? ", slice1 == nil)
	fmt.Println("var style nil? ", slice2 == nil)
	fmt.Println("make style nil? ", slice3 == nil)

}
