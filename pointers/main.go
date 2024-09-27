package main

import "fmt"

func main() {
	fmt.Println("Init function")

	// var ptr *int
	// fmt.Println("Value of ptr: ", ptr)

	// myNumber := 23

	// var ptr *int = &myNumber

	// fmt.Println("Value of myNumber: ", ptr)
	// fmt.Println("Value of myNumber: ", myNumber)
	// fmt.Println("Value of myNumber: ", *ptr)
	// fmt.Println("Value of myNumber: ", &myNumber)

	// *ptr = *ptr * 2

	// fmt.Println("Value of myNumber: ", myNumber)
	// fmt.Println("Value of myNumber: ", ptr)
	// fmt.Println("Value of myNumber: ", *ptr)

	//Array

	var arr [5]int

	arr[0] = 1
	arr[1] = 12
	arr[2] = 13
	arr[3] = 14

	fmt.Println("Value of arr: ", arr)
	fmt.Println(" Len  of arr: ", len(arr))

	//slices on go
	var slice = []int{}
	slice = append(slice, 1)

	fmt.Println("Value of slice: ", slice)
}
