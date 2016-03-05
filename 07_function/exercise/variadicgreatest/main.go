package main

import "fmt"

func greatest(numbers ...int) int {
	var greatest int
	for _, num := range numbers {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}

func main() {
	greatestnum := greatest(4, 7, 10, 99, 15, 18, 27)
	fmt.Println(greatestnum)
}
