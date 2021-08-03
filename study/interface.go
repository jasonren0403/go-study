package main

import(
	"fmt"
	"proj1/models"
)
func main() {
	p := models.Phone{}
	c := models.Computer{}
	c2 := models.Camera{}


	c.Work(p)
	c.Work(c2)

	var x interface {}
	// can assign any value to empty interface
	x = p
	// type assertion, if failed to convert the program will panic
	p = x.(models.Phone)
	// with check
	if y,ok := x.(models.Phone);ok{
		fmt.Println("ok, continue execution")
		fmt.Println(y)
	}else{
		fmt.Println("Convert failed")
	}
}
