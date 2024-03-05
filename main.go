package main

import (
	"fmt"
)

type Numbers interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func add[T Numbers](a, b, c T) T {
	return a + b + c
}

func main() {
	result := add(1.0, 2.0, 3.1)
	fmt.Println(result)
}
