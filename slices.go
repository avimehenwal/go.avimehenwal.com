/*
	Author	: avimehenwal
	Created	: 10-may-2017
	Purpose	: golang slices and array datatypes

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

https://blog.golang.org/go-slices-usage-and-internals

# SLICE
A slice is a descriptor of an array segment. It consists of a pointer to the array,
the length of the segment, and its capacity (the maximum length of the segment).

re-slicing a slice doesn't make a copy of the underlying array. The full array will be kept in memory
until it is no longer referenced. Occasionally this can cause the program to hold all the data in
memory when only a small piece of it is needed.

*/

package main
import "fmt"

func main() {

	// Arrays  value filling from loop
	var arr [5]int
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	fmt.Println(arr)

	// Array manual population with and without range
	x := [4]float64{98, 93, 77, 82}
	b := [...]string{"Penn", "Teller"}
	fmy.Println(x,b)

	// SLICE
	letters := []string{"a", "b", "c", "d"}

	//make allocates an array and returns a slice that refers to that array
	func make([]T, len, cap) []T

	s := make([]byte, 5)
	len(s) == 5
	cap(s) == 5

	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	// b[1:4] == []byte{'o', 'l', 'a'}, sharing the same storage as b

	// b[:2] == []byte{'g', 'o'}
	// b[2:] == []byte{'l', 'a', 'n', 'g'}
	// b[:] == b

	// A slice can also be formed by "slicing" an existing slice or array
	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // a slice referencing the storage of x

	// Dynamcally adding to slices
	p := []byte{2, 3, 5}
	p = AppendByte(p, 7, 11, 13)
	// p == []byte{2, 3, 5, 7, 11, 13}
	func append(s []T, x ...T) []T

	a := make([]int, 1)
	// a == []int{0}
	a = append(a, 1, 2, 3)
	// a == []int{0, 1, 2, 3}
}