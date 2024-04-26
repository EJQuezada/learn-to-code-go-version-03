package main

import (
	"fmt"
)

func main() {
	// nested loops
	for i := 0; i < 5; i++ { //outer loop
		fmt.Println("--")
		for j := 0; j < 5; j++ { //inner loop
			fmt.Printf("outer loop %v \t inner loop %v\n", i, j)
		}
	}
}
