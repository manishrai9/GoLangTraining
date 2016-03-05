package main

import "fmt"

func main() {
	slice := make([]int, 3, 5)

	slice[0] = 1

	slice[1] = 4
	slice[2] = 6
	//slice[3] = 9 //will throw error at runtime, use append method
	slice = append(slice, 9)

	fmt.Println("length : ", len(slice))
	fmt.Println("Capacity : ", cap(slice))
}
