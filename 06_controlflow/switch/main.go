package main

import "fmt"

func main() {
	switch "Manish" {
	case "Manish":
		fmt.Println("Hi Manish")
		//fallthrough //following case will also be executed
	case "Pratima":
		fmt.Println("Hi Pratima")
	case "Akash":
		fmt.Println("Hi Akash")
	default:
		fmt.Println("Default")

	}
}

//break isn't required
