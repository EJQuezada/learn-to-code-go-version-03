package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)
	fmt.Printf("%d \t %b - %v", 1<<20, 1<<20, math.Pow(2, 20))
}
// shifted 100 places
// 1267650600228229401496703205376
// 18446744073709551615