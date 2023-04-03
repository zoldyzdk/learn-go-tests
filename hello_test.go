package main

import (
  "testing"
)

func TestHello(t *testing.T) {
  result := Hello("Isagi")
  expected := "Hello, Isagi"

  if result != expected {
    t.Errorf("Result '%s', expected '%s'", result, expected)
  }
}
