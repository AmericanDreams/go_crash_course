package main

import "fmt"

func main() {
	var idNames = make(map[int]string)

	idNames[1] = "Mirali"
	idNames[2] = "John"

	fmt.Println(idNames)

	// Deleting from map
	delete(idNames, 1)

	fmt.Println(idNames)

	// Short hand
	ageNames := map[int]string{44: "Mirali", 66: "John"}
	fmt.Println(ageNames)
}
