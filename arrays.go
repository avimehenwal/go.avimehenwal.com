package main
import "fmt"

/*  GoLang Arrays and Maps  */

func main() {
  // Arrays
  var arr [5]int
  for i := 0; i < len(arr); i++ {
    arr[i] = i
  }
  fmt.Println(arr)

  x := [4]float64{
     98,
     93,
     77,
     82,
     // 83,
  }
  fmt.Println(x)

  // Maps
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
}
