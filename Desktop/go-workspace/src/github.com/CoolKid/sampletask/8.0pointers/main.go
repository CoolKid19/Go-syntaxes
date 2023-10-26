package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")

	a := 5
	b := &a

	fmt.Println(a, b)
	fmt.Printf("%T\n", b)

	// use * to read val from address
	fmt.Println(*b)
	fmt.Println(*&a)

	// change val with pointer

	*b = 10
	fmt.Println(a)
}