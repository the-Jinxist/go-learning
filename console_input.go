package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Implicit declaration, hehehe
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the year of your birth: ")

	scanner.Scan()

	input, _ := strconv.ParseInt(scanner.Text(), 0, 64)
	
	fmt.Printf("You wil be %d years old at the end of 2020", 2021 - input)
}
