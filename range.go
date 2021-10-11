package main

import "fmt"

func main() {
	var a []int = []int{1, 3, 4, 5, 6, 7, 8, 4, 9, 6, 12}

	//Using the `range` to cycle through the elements in the slice
	//`i` represents the current index
	//`element` stands for the element and the current index
	for i, element := range a {
		fmt.Printf("\n Element %d: %d\n", i, element)
	}

	for i, element := range a {
		for _, element2 := range a[i+1:] {
			if element == element2 {
				fmt.Printf("\nDuplicate: %d", element2)
			}
		}
	}
}
