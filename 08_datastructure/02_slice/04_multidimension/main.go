package main

import "fmt"

func main() {
	records := make([][]string, 0)

	student1 := make([]string, 4)
	student1[0] = "Manish"
	student1[1] = "Rai"
	student1[2] = "Maths"
	student1[3] = "99"

	student2 := make([]string, 4)
	student2[0] = "Pratima"
	student2[1] = "Rai"
	student2[2] = "Science"
	student2[3] = "99"

	records = append(records, student1)
	records = append(records, student2)

	fmt.Println("records : ", records)

}
