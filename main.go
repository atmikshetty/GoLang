package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello Atmik Shetty.")

	var name string = "Atmik"
	var age = 18
	fmt.Println(name)
	fmt.Print(age)

	var dimensions float64 = 90.876
	fmt.Println(dimensions)

	var claim bool = true
	claim = true
	fmt.Print(claim)

	const pi = 3.14
	fmt.Println(pi)

	goat := "Messi"
	fmt.Print(goat)

	// Starting with capital letter: Public can be used anywhere
	// Starting with small letter: private can be used only within the file
}
