package main

import "fmt"

func main() {

	a := 5
	b := &a // now b keeps the value of memori location of a

	fmt.Println(a, b)

	c := *b // With help of * , we get the value which is stored given address. Basically reverse of &

	fmt.Println(c)

	*b = 10 // Also we can change the value which is stored there
	fmt.Println(a)

	// By default everhting is passed by value in the go. So we uses the pointers if we wanna pass by ref

}
