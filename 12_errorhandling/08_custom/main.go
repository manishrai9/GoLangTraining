package main

import (
	"fmt"
	"log"
)

//NorgateMathError Custom type
type NorgateMathError struct {
	lineNum int
	err     error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured at line: %v , error is:  %v", n.lineNum, n.err)
}
func main() {
	_, err := sqrt(-42.56)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, &NorgateMathError{24, fmt.Errorf("square root of negative number: %v", num)}
	}
	//actual implementation
	return num, nil
}
