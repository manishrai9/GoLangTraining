package main

import "fmt"

func main() {
	max := max(7)
	fmt.Printf("Max of 7 is %d\n", max)

}

func max(x int) int {
	return 42 + x
}

//dont do this, use separate var name
