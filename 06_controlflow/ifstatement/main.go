package main

import "fmt"

func main() {
	for i := 0; i <= 5; i++ {
		if i%2 == 0 {
			fmt.Println(i, "is even number")
		} else {
			fmt.Println(i, "is odd number")
		}
	}
}
