package helloWorld

import "fmt"

func Calc(opnd1 float64, opnd2 float64, op byte) float64 {
	var res float64
	switch op {
	case '+':
		res = opnd1 + opnd2
	case '-':
		res = opnd1 - opnd2
	case '*':
		res = opnd1 * opnd2
	case '/':
		if opnd2 == 0 {
			panic("opnd2 cannot be 0 in divide")
		}
		res = opnd1 / opnd2
	default:
		panic("Invalid operator")
	}
	return res
}

func invisible(){
	fmt.Println("I am not visible to other Packages.")
}