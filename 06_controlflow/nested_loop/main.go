package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Println(i * j)
		}
	}
}
