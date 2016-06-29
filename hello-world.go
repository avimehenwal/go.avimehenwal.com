package main
import "fmt"

// this is a comment
func main() {
    //var x string             // go variables <var name> <type>
    x := "Hello World"       // Var default values assignment. Skip type
    fmt.Println(x)
    x = "first "             // var x takes/assigned value of 'first'
    fmt.Println(x)
    x = x + "second"
    fmt.Println(x)
}
