package main

import "fmt"

/*  
*** DATA TYPES ***

## Numbers
    uint8, uint16, uint32, uint64,
    int8,  int16,  int32,  int64
There are also 3 machine dependent integer types: *uint, int and uintptr*
They are machine dependent because their size depends on the type of architecture you are using

## Floating Point
	float32,   float64
	complex64, complex128
 
## Bollean
	true, false

## String
	string
*/

func main() {

	// Arrays automa value filling
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	// Array manual population
	x := [4]float64{
		98,
		93,
		77,
		82,
		// 83,
	}
	fmt.Println(x)

	// Maps are a dynamically growable associative array type,
	/*
	The make built-in function allocates and initializes an object of type slice, map, or channel(only).
	Like new, the first argument is a type, not a value.
	Unlike new, make's return type is the same as the type of its argument, not a pointer to it
	*/
	m := make(map[string]int)
	m["key"] = 10
	fmt.Println(m)

	if name, ok := m["Un"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println(name, ok)
	}
	if name, ok := m["key"]; ok {
		fmt.Println(name, ok)
	} else {
		fmt.Println(name, ok)
	}

	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map shorthand", n)

    value, present := n["bar"]
    fmt.Println("present:", present)
    fmt.Println("value:", value)

    delete(n, "bar")
    fmt.Println("map:", n)

    value, present = n["bar"]
    fmt.Println("present:", present)
    fmt.Println("value:", value)
}
