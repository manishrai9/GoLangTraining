package main

import "fmt"

func main() {
	x := 1
	str := evaluator(x)
	fmt.Println(str)
}

func evaluator(n int) string {
	if n < 10 {
		return fmt.Sprint("num is less than 10")
	} else {
		return fmt.Sprint("num is greate than 10")
	}

}
