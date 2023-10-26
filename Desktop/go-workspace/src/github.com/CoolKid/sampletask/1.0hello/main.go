package main

import "fmt"

func main()  {
	//var name int32 = 23
	const age int32 = 23
	name := "CoolKid"
	var size float32 = 1.3
	fmt.Printf("%T\n", size)
	fmt.Println("Hello", name, age)
	fmt.Println("Hello World")
	fmt.Printf("Hello World, %T\n", name)

	game, mission := "GTA", "GTA5"
	fmt.Println(game, mission)
}