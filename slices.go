package main

import "fmt"

func main() {

	var x [5]int = [5]int{1, 2, 3, 4, 5}

	///Creating a slice
	///`[:]` is the slice operator.
	///   Any index before the colon is where the slice should start from. No value before the colon means start at the beginning
	///   Any index after the colon is where the slice should end at, minus 1. No value after the colon means finish at the end, minus 1

	//Here, we are slicing out a piece of x
	var s []int = x[1:3]

	//Prints out [2,3]
	fmt.Println(s)

	//`cap(array/slice)` is used to determine the capacity of an array or slice
	//Prints 4
	fmt.Println(cap(s))

	//Re-extend the slice to fit all the element in the array
	//Prints [2,3,4,5]
	fmt.Println(s[:cap(s)])

	// ----------------------------------------------------------------------------------------

	//Creating a slice
	var slice []int = []int{5, 6, 7, 8}

	//Adding value to a slice using `append(slice, value)`
	newSlice := append(slice, 9)

	//Should print out a new slice with one more element: [5,6,7,8,9]
	fmt.Println(newSlice)

	//-------------------------------------------

	//Creating a slice with make
	madeSlice := make([]int, 5)
	fmt.Println(madeSlice)

}
