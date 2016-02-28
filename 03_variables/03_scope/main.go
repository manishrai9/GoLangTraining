package main

import "fmt"

//package level scope
var a int = 10

func main() {
	fmt.Println(a)
	//method level scope
	y := 20
	fmt.Println(y)
	foo()
	{

		fmt.Println(a)
		y := 10
		fmt.Println(y)

	}
	//fmt.Println(y) //outside scope
}

func foo() {
	fmt.Println(a)
	//fmt.Println(y)
}
