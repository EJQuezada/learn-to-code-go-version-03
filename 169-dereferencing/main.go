package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	fmt.Println(&x)

	y := &x
	fmt.Printf("%v\t%T\n", y, y)
	fmt.Println(*y)
	fmt.Println(*&x)

	*y = 43 //this memory address only because *
	fmt.Println(x)
	fmt.Println(*y)
}
