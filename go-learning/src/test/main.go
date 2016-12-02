package main

import "fmt"
import "test/handler"

func init()  {
  fmt.Println("main init")
}

func main()  {
  fmt.Println("main main")
  fmt.Println(handler.Data)
  fmt.Println(handler.User)

}
