package main

import "fmt"

func main() {
	age := 20
	name := "Atmik Shetty"
	height := 5.943333

	// Println adds spaces before printing the variables and adds a new line
	fmt.Println("age: ", age, "Height: ", height, "name :", name)

	// keeps printing in the same line and takes '%' like in C lanngauge
	fmt.Printf("age is %d\n", age)       //int
	fmt.Printf("height is %.2f", height) //float
	fmt.Printf("Name is %s\n", name)     //string

	fmt.Printf("Age: %d Name is %s and Height is %.2f \n", age, name, height)

	// Checking the datatype
	fmt.Printf("Type of name is %T", name)
}
