package calculator

import (
  "fmt"
)

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) int {
  if y != 0 {
    return x / y
  } else {
    err := fmt.Errorf("Error: denominator cannot be %d", y)
    fmt.Println(err.Error())
  }
}
