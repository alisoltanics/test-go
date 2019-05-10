package main

import "fmt"

func main(){
	var name = "Ali"
	var age = 23
	var isCool = false
	const myVar = 1
	// Shorthand
	lname := "Soltani"
	myFloat := 2.1
	f, s := "f", "s"
	fmt.Println(name, lname, age, isCool)
	fmt.Println(myFloat, f, s)
	// prints type of var 
	fmt.Printf("%T\n", age)

}
