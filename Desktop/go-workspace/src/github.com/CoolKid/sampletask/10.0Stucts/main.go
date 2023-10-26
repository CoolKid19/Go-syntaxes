package main

import ("fmt"
		"strconv"
)

// Define person struct

type Person struct {
	// firstname string
	// lastname string
	// city string
	// gender string
	// age int

	// alternative

	firstname, lastname, city, gender string
	age int

}

// Greeting method (value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstname + " " + p.lastname + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)

func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)

func (p *Person) getMarried(spouseLastname string) {
	if p.gender == "m" {
		return
	} else {
		p.lastname = spouseLastname
	}
}


func main() {

	// Init person using struct
	//person1 := Person{firstname: "sam", lastname: "wheeler", city: "Boston", age: 24}

	// alternative
	person1 := Person{"sam", "wheeler", "Boston", "m", 24}

	fmt.Println(person1)

	fmt.Println(person1.firstname)

	fmt.Println(person1.greet())

	person1.hasBirthday()
	person1.getMarried("williams")
	fmt.Println(person1.greet())
}