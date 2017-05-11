/*
	Author	: avimehenwal
	Created	: 11-may-2017
	Purpose	: golang Method Indirection

Choosing a value or pointer receiver
There are two reasons to use a pointer receiver.
1. The first is so that the method can modify the value that its receiver points to.
2. The second is to avoid copying the value on each method call. This can be more efficient
   if the receiver is a large struct

In general, all methods on a given type should have either value or pointer receivers, but not a mixture of both. (We'll see why over the next few pages.)
*/


package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
