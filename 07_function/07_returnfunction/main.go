package main

import "fmt"

func greeter() func() string {
	return func() string {
		return "Hello Manish!"
	}
}

func main() {
	//anonymous function
	greet := greeter()
	fmt.Println(greet())
}
