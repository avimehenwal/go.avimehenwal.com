package main
import "fmt"

/*
   GoLang Arrays and Maps
*/
func main() {
  var arr [5]int
  for i := 0; i < len(arr); i++ {
    arr[i] = i
  }
  fmt.Println(arr)

  // Maps
  m := make(map[string]int)
  m["key"] = 10
  fmt.Println(m)
}
