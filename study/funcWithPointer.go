package main

import "fmt"

func main(){
	a1 := 20
	fmt.Println("a1=", a1)
	_fun(a1)
	fmt.Println("~a1=", a1)
	_fun2(&a1)
	fmt.Println("~~a1=", a1)

	// anonymous function
	a2 := func(a1 int, a2 int) int{
		return a1+a2
	}(10,20)

	fmt.Println("a2=",a2)

	a3 := func(a1 int,a2 int)int{
		return a1-a2
	}
	fmt.Println("a3=",a3(40,20))
}

func _fun(n1 int) {
	n1 = n1 + 10
	fmt.Println("n1 in _fun():", n1)
}

func _fun2(n1 *int) {
	*n1 = *n1 + 10
	fmt.Println("n1 in _fun2()", *n1)
}
