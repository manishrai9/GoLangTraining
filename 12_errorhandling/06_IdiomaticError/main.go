package main

import (
	"errors"
	"log"
)

//ErrNorgateError - exported outside of package
var ErrNorgateError = errors.New("Norgate Math: sqrt of negative number")

func main() {
	_, err := sqrt(-42)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, ErrNorgateError
	}
	//actual implementation
	return num, nil
}
