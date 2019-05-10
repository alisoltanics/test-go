package main

import "fmt"

func main(){
	// numbers := make(map[string]string)
	// numbers["ali"] = "0919554844"
	// numbers["ali2"] = "09544969"

	numbers := map[string]string{"ali": "1424527", "ali2": "57476964"}

	fmt.Println(numbers)
	fmt.Println(numbers["ali"])

	delete(numbers, "ali")
	fmt.Println(numbers)
}
