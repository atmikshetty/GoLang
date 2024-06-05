package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello.. whats your name???")
	// var name string
	// fmt.Scan(&name) // only reads until the first whitespace character
	// fmt.Println("Hello Mr.", name)

	// Using the Bufio package
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // read until the new line
	fmt.Println("Hello Mr.", value)
}
