package main

import "fmt"

func half(n int) (int, bool) {
	//half := n / 2
	//var isHalf bool
	//if n%2 == 0 {
	//	isHalf = true
	//}
	//return half, isHalf
	return n / 2, n%2 == 0
}

func main() {
	h, even := half(2)
	fmt.Println("half of 2: ", h, " and is 2 even? ", even)
	//fmt.Println(half(1))
}
