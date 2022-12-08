package main
import "fmt"
//Type assertions allow us to access the data and data type of values stored by the interface.


func main() {

  // create an empty interface
  var a interface {}

  // store integer to an empty interface
  a = 10

  // type assertion
  interfaceValue := a.(int)

  fmt.Println(interfaceValue)
}
