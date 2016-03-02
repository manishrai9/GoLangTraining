package main

import "fmt"

func main() {
	total := avegare(12, 34, 56, 78)
	fmt.Println(total)
	//another way of passing variading params
	data := []float64{11, 12, 13, 14}
	t := avegare(data...)
	fmt.Println(t)
}

func avegare(sf ...float64) float64 { //variadic function taking varibale number of arguements
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
