package main
import "fmt"

/*  Structures and Interfaces  */

func main() {
  type Circle struct {
    x float64
    y float64
    r float64
  }

  // Go Methods area() reciever <name> <type>
  func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
  }

  c := Circle{x: 0, y: 0, r: 5}
  fmt.Println(c)
  fmt.Println(c.area())

  // slice of structure
  s := []struct {
    i int
    b bool
  }{
    {2, true},
    {3, false},
    {5, true},
    {7, true},
    {11, false},
    {13, true},
  }
  fmt.Println(s)
}
