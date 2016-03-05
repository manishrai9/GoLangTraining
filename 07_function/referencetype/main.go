package main

import "fmt"

func changeMe(z []string) {
	z[0] = "Manish"
	z[1] = "Pratima"
	fmt.Println(z)
}

func main() {
	z := make([]string, 2, 10)
	fmt.Println(z)
	changeMe(z)

}
