package main

import "fmt"

func main() {

	//Explicit declaration
	var sentence string = "A string"

	//Implicit declaration
	var number = 256.98

	//Implicit variable declaration may cause the compiler to assign the wrong type
	//This only happens in very, very rare occasions

	//Expression assignment operator
	//Basically assigns the value without the 'var' operator
	speed := 500.45

	fmt.Printf("Number type %T", number)
	fmt.Printf("Number type %T", speed)
	fmt.Printf("Number type %T", sentence)

	//Omo, before assigning values, Go actually gives the variable a default
	//value, case in point:
	var defaultValueChecker bool

	fmt.Println(defaultValueChecker)
}