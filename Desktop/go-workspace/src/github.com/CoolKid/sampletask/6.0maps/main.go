package main

import "fmt"

func main() {
	// emails := make(map[string]string)

	// // assign key values
	// emails["Bob"] = "bobgmail.com"
	// emails["Sharon"] = "sharongmail.com"
	// emails["Mike"] = "mikegmail.com"

	// declare map and add key values

	emails := map[string]string{"Bob": "bobgmail.com", "Sharon": "sharongmail.com"}

	fmt.Println(emails)
	fmt.Println(emails["Bob"])

	// delete from map

	delete(emails, "Bob")
	fmt.Println(emails)




}