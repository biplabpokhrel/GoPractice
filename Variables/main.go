package main

import "fmt"

func main() {
	var name string
	fmt.Println(name)

	firstName := "Biplab"
	fmt.Println(firstName)

	var lastName = "Pokhrel"
	fmt.Println(lastName)

	var id, fullname = 1, "Biplab Pokhrel"
	fmt.Printf("%d\t%v\n", id, fullname)

	number1, number2 := 1, 99
	fmt.Printf("%d\t%d\n", number1, number2)
}
