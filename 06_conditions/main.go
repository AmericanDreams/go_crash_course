package main

import "fmt"

func main() {
	age := 66
	color := "red"

	if age > 60 {
		fmt.Println("You are old bro")
	} else {
		fmt.Println("Good!")
	}

	switch color {
	case "red":
		{
			fmt.Println("Red color")
			break
		}
	case "black":
		{
			fmt.Println("Black color")
			break
		}
	default:
		{
			fmt.Println("Default color")
		}
	}

}
