package main

import "fmt"

func main() {
	s := greet("Pratima ", "Rai")
	fmt.Println(s)
	m := greet1("Manish ", "Rai")
	fmt.Println(m)
}

func greet(fname string, lname string) string {
	return fmt.Sprint(fname, lname)
}

func greet1(fname string, lname string) (s string) {
	s = fmt.Sprint(fname, lname)
	return
}
