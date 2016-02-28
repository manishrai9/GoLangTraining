package main

import "fmt"

//constant declaration
const a = "Manish"

//another way of declaring constants
const (
	pi       = "3.14"
	language = "GO"
)

const (
	A = iota //0
	B = iota //1
	C = iota //2
)

func main() {
	const b = 10
	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("PI - ", pi)
	fmt.Println("Language - ", language)

	fmt.Println("A - ", A)
	fmt.Println("B - ", B)
	fmt.Println("C - ", C)
}
