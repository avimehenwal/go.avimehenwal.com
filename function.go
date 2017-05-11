/*
	Author	: avimehenwal
	Created	: 10-may-2017
	Purpose	: golang functions

# GO FUNCTIONS

- Multiple returns (explicit)
- Anonymous Functions
- Variadic Functions
- Closures
- Recursions

Functions are values too. They can be passed around just like other values.
Function values may be used as function arguments and return values.

*/

package main
import "fmt"

// returned function _closes over_ the variable `i` to form a closure.
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func f() (int, int) {
	return 5, 6
}

// Variadic functions
func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func main() {
	x, y := f()
	fmt.Println("x=", x, "y=", y)
	fmt.Println("sum=", add(1, 2, 3, 4))

	// Variadic functions with sliced parameters
	xs := []int{1, 2, 3}
	fmt.Println(add(xs...))

    // We call `intSeq`, assigning the result (a function) to `nextInt`.
    // This function value captures its own `i` value, which will be updated each time we call `nextInt`.
    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // To confirm that the state is unique to that
    // particular function, create and test a new one.
    newInts := intSeq()
    fmt.Println(newInts())
}