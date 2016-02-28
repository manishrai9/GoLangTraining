package main

import "fmt"

func main() {
	a := 43
	fmt.Println(a)  //43
	fmt.Println(&a) //memory address e.g. 0xc08203c1b8

	var b *int = &a
	fmt.Println(b)  //0xc08203c1b8
	fmt.Println(*b) //it is called de-referencing, i.e. 43
	//b is type of "int pointer"

	*b = 42        //change the value to 42
	fmt.Println(a) //42
}

//everything is PASS BY VALUE in GO
