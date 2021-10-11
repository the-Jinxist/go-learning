package main

import "fmt"

func main() {

	//Map allows us to store key-value pairs
	mp := map[string]int{
		"monkey": 0,
		"eat":    6,
		"orange": 9,
	}

	//To change and access values
	mp["monkey"] = 3

	fmt.Println(mp)
}
