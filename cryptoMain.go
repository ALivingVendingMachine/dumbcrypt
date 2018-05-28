package main

import (
  "fmt"
  "dumbcrypt/cryptoFuncs"
)

func main()  {
  fmt.Println("Hello")
  var hundred int = 1
  hundred = hundred << 1
  fmt.Println(hundred)
  fmt.Println(cryptoFuncs.FastModExp(3, 218, 1000))
  //var a bool = cryptoFuncs.RabinMillerPrimality(6, 2)
  //fmt.Println(a)
  //a = cryptoFuncs.RabinMillerPrimality(7, 2)
  //fmt.Println(a)
  a := cryptoFuncs.RabinMillerPrimality(3215031751, 2)
  fmt.Println(a)
}
