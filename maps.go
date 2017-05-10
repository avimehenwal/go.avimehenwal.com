/*
	Author	: avimehenwal
	Created	: 10-may-2017
	Purpose	: golang slices datatypes

*/

package main
import "fmt"

func main() {
	
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
