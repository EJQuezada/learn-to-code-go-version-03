package main

import "fmt"

func main() {

	f := outer()
	fmt.Println(f()) //42
}

func outer() func() int {
	return func() int {
		return 42
	}
}
