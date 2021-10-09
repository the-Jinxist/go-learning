package main

import "fmt"

func main() {

	//Valid for-loop in go
	for {

		//Immediately exits the loop
		break

		//Goes to the next iteration immediately
		continue
	}

	//Another valid for loop in go
	x := 0
	for x < 5 {
		fmt.Println(x)
		x++
	}

	//Another valid for loop in go
	for y := 0; y < 5; y++ {
		fmt.Println(y)
	}

	//Testing continue statement
	for z := 0; z <= 1000; z++ {
		if z != 0 && z%9 == 0 {
			fmt.Println(z)
			continue
		}

	}

	for e := 0; e <= 1000; e++ {
		if e != 0 && e%9 == 0 {
			fmt.Println(e)
			break
		}

	}

}
