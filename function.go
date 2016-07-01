package main
import "fmt"

/*  Functions  */

func f() (int, int) {
  return 5, 6
}

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
  fmt.Println("sum=", add(1,2,3,4))

  xs := []int{1,2,3}
  fmt.Println(add(xs...))
}
