package main

import (
	"fmt"
	
)

func main() {
	//fmt.Println("Hello, World!")
	x := 5
	y := 10
	
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Printf("%d is equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is greater than %d\n", x, y)
	}
	color := "green"
	switch color {
	case "red":
		fmt.Println("Color is red")
	case "blue":
		fmt.Println("Color is blue")
	default:
		fmt.Println("Color is not red or blue")

	}

}