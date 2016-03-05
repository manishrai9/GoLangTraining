package main

import "fmt"

func main() {
	var str [58]string

	fmt.Println("empty string array : ", str)
	fmt.Println("string array length : ", len(str))

	for i := 65; i <= 122; i++ {
		str[i-65] = string(i)
	}
	fmt.Println("string array : ", str)
}

//doesn't change size
