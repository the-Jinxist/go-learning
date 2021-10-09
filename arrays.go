package main

import "fmt"

func main() {
	//Arrays can only store types of the same value
	//Arrays are fixed in length, their length has to be specified when it is initialized
	var array [5]int

	array[0] = 100
	array[4] = 300
	fmt.Println(array)

	//Explicitly defining elements in the array
	arr := [5]int{2, 3, 4}
	fmt.Println(arr)

	//len(array) returns the length of the array
	fmt.Println(len(arr))

	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}

	fmt.Println(sum)

	//A multidimensional array
	arr2D := [2][3]int{{1, 2, 0}, {1, 2, 5}}
	fmt.Println(arr2D)
	fmt.Println(arr2D[1][0])

}
