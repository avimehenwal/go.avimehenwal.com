/*
	Author	: avimehenwal
	Created	: 11-may-2017
	Purpose	: golang pointers

Poiinters Reference and Dereference

* => is a pointer to a variable and denotees its underlying value
& => Generates a pointer to its operands 

(*p).X == p.X where p is a pointer to a structure with attribute X

*/

package main
import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	dr  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1

	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j

	fmt.Println(v1, dr, v2, v3)		// {1 2} &{1 2} {1 0} {0 0}

}
