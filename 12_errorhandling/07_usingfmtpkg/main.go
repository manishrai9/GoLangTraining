package main

import (
	"fmt"
	"log"
)

func main() {
	_, err := sqrt(-42.56)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, fmt.Errorf("sqrt of negative number: %v", num)
	}
	//actual implementation
	return num, nil
}
