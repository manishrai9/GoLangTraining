package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}
		if i >= 10 {
			break
		}
		fmt.Println(i)

	}

	// another way for loop
	//i = 0
	//for i <= 10 {
	//	fmt.Println(i)
	//	i++
	//}
}
