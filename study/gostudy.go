package main

import (
	"fmt"
	"proj1/main"
	"strconv"
)

func main() {
	//var op1,op2 float64
	//var opnd byte
	//fmt.Println("Input first operand")
	//_, _ = fmt.Scanln(&op1)
	//fmt.Println("Input second operand")
	//_, _ = fmt.Scanln(&op2)
	//fmt.Println("Input operator(+/-/*//)")
	//_, _ = fmt.Scanf("%c",&opnd)
	//pointer()

	//result := helloWorld.Calc(op1,op2,opnd)
	//fmt.Println("result=",result)
	fmt.Println("from helloworld.go", "Model:", helloWorld.Model)
	helloWorld.SayHelloWorld()
}

func returningMulValues() (val1, val2, val3 string) {
	return "a", "b", "c"
}

func vars_and_values() {
	/* 变量声明 注意类型在变量名之后，无需分号作为结束符 */
	var _ int
	/* :=编译器自动推导，但是左边不能是声明过的变量 */
	v3 := "hey"
	v2 := 2
	/* 用圆括号将若干个声明变量放在一起 */
	var (
		v114 int = 1
		v250     = ""
		//vbool = true
	)
	/* 用下划线 _ 代表匿名变量 */
	_, v250, _ = returningMulValues()
	/* Go 语言支持多重赋值 */
	v2, v114 = v114, v2
	// const关键字定义常量
	const hahaha, zero = "hey", 0.0
	// 一组枚举常量
	const (
		Sunday = iota
		//iota 是一个可被编译器修改的常量，每一个const关键字出现时被重置为0，在下一次const出现前，每次iota所代表的的数字都会加1
		Monday /* 不写默认等同于前一个表达式，也就是iota */
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf("Hello world %d %s %s\n", v2, v3, v250)
}

func str_and_char() {
	var c1 byte = 'c'
	fmt.Printf("c1=%c\n", c1)
	var str string = "Hello world"
	ch := str[0] // 取第一个字符
	//str2 := `abc\n<23>` 反引号，位于其中的字符全部以原生形式输出
	str += "h1" + "h2" + // 多行拼接，加号要留在上层
		"he"
	// len(s) 字符串长度
	fmt.Printf("The first character of %s is %c. \n", str, ch)
	fmt.Printf("The length of the string is %d\n", len(str))
}

func strConversion() {
	str1 := "true"
	bol, _ := strconv.ParseBool(str1)
	fmt.Printf("%t\n", bol)
	str2 := "64"
	n1, _ := strconv.ParseInt(str2, 10, 64)
	fmt.Printf("%d\n", n1)
}

func pointer() {
	var i int = 10
	// use &var to get the address of a variable
	fmt.Println("Address of i is", &i)
	// ptr is a pointer variable pointing to int
	var ptr *int = &i // now points to i
	fmt.Printf("ptr=%v\n", ptr)
	fmt.Printf("ptr is pointing to %v\n", *ptr)
}
