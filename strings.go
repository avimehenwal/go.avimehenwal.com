package main

import "fmt"

// this is a comment
func main() {
	fmt.Println("Hello World")
	fmt.Println("length", len("Hello World")) // length
	fmt.Println("Hello World"[10])   // str index byte representation
	fmt.Println("Hello " + "World") // concatinate

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)
	output := input * 2
	fmt.Println(output)
}
