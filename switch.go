package main

import "fmt"

func main() {

	ans := 10

	switch ans {
	case 1, -1, 50:
		fmt.Println("1")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("not a case")
	}

}
