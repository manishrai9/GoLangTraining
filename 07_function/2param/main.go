package main

import "fmt"

func main() {
	greet("Manish", "Rai")
}

func greet(fname string, lname string) {
	fmt.Println(fname, lname)
}
