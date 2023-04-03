package main

import (
  "fmt"
)

const prefixHello = "Hello, "

func Hello(nome string) string {
  return prefixHello + nome
}

func main() {
  fmt.Println(Hello("mundo"))
}
