package main

import "fmt"

const PI = 3.14

func sumar(numeros ...int) int {
  total := 0
  for _, numero := range numeros {
    total = numero + total
  }

  return total
}

func main() {
  fmt.Println(sumar(2))
  fmt.Println(sumar(2, 2))
  fmt.Println(sumar(5, 4 , 3))
}
