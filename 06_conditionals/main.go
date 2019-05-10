package main

import "fmt"

func main(){
	x := 5
	y := 9

	if x > y {
		fmt.Println("hi")
	} else if x == 5 {
		fmt.Println("oh")
	} else {
		fmt.Println("bye")
	}

	switch x {
		case 5:
			fmt.Println("ok")
		case 2:
			fmt.Println("not ok")
		default:
			fmt.Println("hmm")
	}
		
}
