package models

import "fmt"

type Person struct {
	name string
	age  int
}

type Student struct{
	Person // inherits 'Person'
	score int
}

type Adult struct {
	Person // inherits 'Person'
	salary int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

func (p Person) GetName() string{
	return p.name
}

func (p Person) GetAge() int{
	return p.age
}

func (p Person) SetName(name string){
	p.name = name
}

func(p Person) SetAge(age int){
	p.age = age
}

// a private method
func (p Person) sayName() {
	fmt.Println(p.name)
}

func (p Person) Calc(n int) {
	res := 0
	for i := 0; i < n; i++ {
		res += i
	}
	fmt.Println(p.name, "can calculate to", res)
}

// a public method
func (p Person) GetInfo() string {
	return fmt.Sprint("Person ", "name:", p.name, " age:", p.age)
}

// passing struct pointer (&p).GetInfo2()
func (p *Person) GetInfo2() string {
	return (*p).name + " " + string(rune((*p).age))
}

func (p Person) String() string {
	return fmt.Sprint("Person ", "name:", p.name, " age:", p.age)
}
