package main

import (
	"fmt"
)

func main() {
	var x float64 = 10
	var y float64 = 3

	//basic non-string data type conversion
	//e.g converting int to float
	var number int = 98;
	var convToFloat float64 =float64(number)

	var answer = x / y

	//Operations between values of a particular data type will
	//give a result in the same data type
	var no1 int = 9
	var no2 int = 2

	//The math package is used to do even more complex math operations like
	//log, exp e.t.c

	fmt.Printf("This division will give an integer %d", no1 / no2)

	fmt.Printf("\nX plus Y equals... %f", answer + convToFloat)
}