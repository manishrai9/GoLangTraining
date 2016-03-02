package main

import "fmt"

func main() {
	data := []float64{11, 12, 13, 14}
	t := avegare(data)
	fmt.Println(t)
}

func avegare(sf []float64) float64 { //variadic function taking varibale number of arguements
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
