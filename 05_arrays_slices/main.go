package main

import "fmt"

func main() {

	// Arrays
	var names [2]string
	names[0] = "Mirali"

	// Short hand
	ages := [2]int{55, 66}

	// Slices (They are basicallay arrays but we dont need to know exact size in advance)
	var ages2 = []int{44, 66, 77}

	fmt.Println(names)
	fmt.Println(ages)

	fmt.Println(ages2)
}
