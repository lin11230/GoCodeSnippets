package main

import "fmt"

func swap(x, y string) (string, string, string) {
 return y, x, y
}

func main(){
  a, b, c := swap("hello","world")
  fmt.Println(a, b, c)
}
