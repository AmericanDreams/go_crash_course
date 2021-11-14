package main

import "fmt"

func main() {
	ids := []int{33, 44, 55, 66, 77}

	fmt.Println(ids)

	// With index
	for i, id := range ids {
		fmt.Println(i, id)
	}

	// Without index
	for _, id := range ids {
		fmt.Println(id)
	}

	//Iteration of the map
	idNames := map[int]string{1: "Mirali", 2: "John"}

	for k, v := range idNames {
		fmt.Println(k, v)
	}

}
