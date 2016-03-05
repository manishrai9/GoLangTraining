package main

import "fmt"

func main() {
	slice := make([]int, 0, 3)
	fmt.Println("length : ", len(slice))
	fmt.Println("Capacity : ", cap(slice))

	for i := 0; i < 20; i++ {
		slice = append(slice, i)
		fmt.Print("length : ", len(slice))
		fmt.Println(" and Capacity : ", cap(slice), " and value is ", slice[i])
	}

}
