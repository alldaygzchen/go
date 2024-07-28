package main

import "fmt"

func main() {
	result := add(1, 2)
	fmt.Println(result)
}

func add[TYPE int | float64 | string](a, b TYPE) TYPE {
	return a + b
}