package main

import "fmt"

type Human struct {
	name, surname string
	age           int
}

// Value reciever (Read only)
func (human Human) printFullName() {
	fmt.Println(human.name + " " + human.surname)
}

// Pinter reciever (To change the value)
func (human *Human) hasBirthDay() {
	human.age++
}

func main() {

	var mirali = Human{"Mirali", "Fikratzade", 55}
	fmt.Println(mirali.name)

	mirali.age = 88
	fmt.Println(mirali)

	mirali.printFullName()
	mirali.hasBirthDay()

	fmt.Println(mirali)

}
