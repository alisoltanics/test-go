package main

import "fmt"

func main(){
	ids := []int{22, 55, 88, 66, 78, 44}

	for i, id := range ids {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	for _, id := range ids {
		fmt.Printf("ID: %d\n", id)
	}

	numbers := map[string]string{"ali": "1424527", "ali2": "57476964"}
	for k, v := range numbers {
		fmt.Printf("%s: %s\n", k, v)
	}
}
