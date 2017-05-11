/*
	Author	: avimehenwal
	Created	: 11-may-2017
	Purpose	: golang Interfaces

An interface type is defined as a set of method signatures.
A value of interface type can hold any value that implements those methods.

One of the most ubiquitous interfaces is Stringer defined by the fmt package.

# Go Interfaces
	- set of go methods on a specific type
	- what actions type can execute
	- Like Python classes
	- everything in Go is passed by value

* Go is procedural language and NOT object roriented
* Go doesnt have classes

http://govspy.peterbe.com/

*/

package main
import (
	"fmt"
)

// interface
type Animal interface {
	Speak() string
}

// argument types for interfaces
type Dog struct{}
type JavaProgrammer struct{}
type Llama struct{}
type Cat struct{}

// interface type implementation
func (d Dog) Speak() string {
	return "Woof!"
}

func (c Cat) Speak() string {
	return "Meow!"
}

func (l Llama) Speak() string {
	return "?????"
}

func (j *JavaProgrammer) Speak() string {
	return "Design patterns!"
}

func main() {
	animals := []Animal{Dog{}, Cat{}, Llama{}, new(JavaProgrammer)}
	for st, animal := range animals {
		fmt.Println(st, animal)
		fmt.Println(animal.Speak())
	}
}
