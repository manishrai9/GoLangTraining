package main

import "fmt"

func main() {
	fmt.Println(greet("Pratima ", "Rai "))
}

func greet(fname string, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
