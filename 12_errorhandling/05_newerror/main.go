package main

import (
	"errors"
	"log"
)

func main() {
	_, err := sqrt(-42)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("sqrt of negative number")
	}
	//actual implementation
	return num, nil
}
