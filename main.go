package main

import (
	"fmt"
)

func main() {
	b := bucket{}
	b.optimize(1.000, 0.001)
	fmt.Println(b, b.volume(), b.surfaceArea())
}
