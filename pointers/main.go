package main

import "fmt"

func main() {
	fmt.Println("Init function")

	// var ptr *int
	// fmt.Println("Value of ptr: ", ptr)

	myNumber := 23

	var ptr *int = &myNumber

	fmt.Println("Value of myNumber: ", ptr)
	fmt.Println("Value of myNumber: ", myNumber)
	fmt.Println("Value of myNumber: ", *ptr)
	fmt.Println("Value of myNumber: ", &myNumber)

	*ptr = *ptr * 2

	fmt.Println("Value of myNumber: ", myNumber)
	fmt.Println("Value of myNumber: ", ptr)
	fmt.Println("Value of myNumber: ", *ptr)

}
