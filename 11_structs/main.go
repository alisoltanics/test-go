package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// age       int
	firstName, lastName string
	age int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " And I'm " + strconv.Itoa(p.age)
}

// HasBirthday method (pointer receiver)
func (p *Person) HasBirthday() {
	p.age++
}

// getMarried method (pointer receiver)
func (p *Person) getMarried(spouseLastName string) string{
	return p.firstName + " married with " + spouseLastName
}

func main() {
	person1 := Person{firstName: "ali", lastName: "sol", age: 22}
	// person1 := Person{"ali", "sol", 22}
	person1.HasBirthday()
	fmt.Println(person1)
	fmt.Println(person1.greet())
	fmt.Println(person1.getMarried("no one"))
}
