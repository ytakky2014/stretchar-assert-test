package main

import (
	"fmt"
)

func main() {
	fmt.Println("test")
	fmt.Println(sum(1, 4))
}

func sum(a, b int) int {
	return a + b
}
