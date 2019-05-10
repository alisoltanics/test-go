package main

import "fmt"

func main(){
	a := 1
	b := &a

	fmt.Println(a, b)
	fmt.Println(*b)
}
