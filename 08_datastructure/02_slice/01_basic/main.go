package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d", "e"}
	fmt.Println(slice)
	fmt.Println(slice[2:4]) //slicing a slice`
	fmt.Println(slice[3])   //index access
	fmt.Println("slice"[3]) //index access
}
