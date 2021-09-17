package main

import "fmt"

func main(){

	//Conditional operators
	// < , <=, ==, !=,  =>, >
	//RHS and LHS values/variables must be of the same type

	val := 1 < 5
	fmt.Printf("%t", val)

	//Comaparing Strings
	name := "Fave"
	lastName := "Fave"

	same := name == lastName
	fmt.Printf("\n%t", same)

}