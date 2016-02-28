package main

import "fmt"

func zero(x *int) {
	//fmt.Println(x)
	*x = 0
}
func main() {
	x := 5
	//fmt.Println(&x)
	//fmt.Println(x)
	zero(&x)
	fmt.Println(x) //x will change to 0 now
}

//pass by value
