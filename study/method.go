package main

import (
	"fmt"
	"proj1/models"
)

func main() {
	// industry mode to create new struct
	a := models.NewPerson("Jason",114514)
	//a.sayName()
	fmt.Println(a.GetInfo())
	a.Calc(100)
	fmt.Println(a)
	b := models.Student{} // <===> sames as models.Adult
	b.Person.SetName("abcd") // <==> b.SetName()
	b.SetAge(42) // b.Person.SetAge()
}
