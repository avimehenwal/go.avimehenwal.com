package main

import "fmt"

/*
   GoLang Variables
   - Simple Variable
   - Default Variable
   - Multiple Variables
*/

func main() {
	//var x string           // go variables <var name> <type>
	x := "Hello World" // Var default values assignment. Skip type
	fmt.Println(x)
	x = "first " // var x takes/assigned value of 'first'
	fmt.Println(x)
	x = x + "second"
	fmt.Println(x)

	const (
		a  = 5
		pi = 3.14
	)
	fmt.Println(a, pi)
}
