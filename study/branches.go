package main

import (
	"fmt"
)

func main() {
	/* go语言循环只有for关键字，且不用写圆括号 */

	for i := 0; i < 10; i++ {
		fmt.Println("Hello 任子恒.", "Printed for", i+1, "times")
	}

	/* 以上等价于这种写法 */
	j := 0
	for j < 10 {
		fmt.Println("Hello 任子恒.", "Printed for", j+1, "times")
		j++
	}

	/* 想要无限循环这么写 */
	var k int = 0
	for {
		if k < 10 {
			fmt.Println("Hello 任子恒.", "Printed for", k+1, "times")
			k++
		} else {
			break
		}
	}

	/* range 遍历字符串和数组 */
	var str string = "hello world中文"
	for index, val := range str {
		//fmt.Println("index", index, "val", val)
		fmt.Printf("index = %d, val=%c\n", index, val)
	}
}

func switchBranch() {
	var choice int
	fmt.Print("Input choice")
	_, _ = fmt.Scanln(&choice)
	switch choice {
	case 1, 2:
		fmt.Println("1,2")
	case 3:
		fmt.Println("3")
	default:
		fmt.Println("default")
	}
	// switch不跟表达式，类似于if-else
	switch {
	case choice == 1:
		fmt.Println("1")
	case choice == 2:
		fmt.Println("2")
		fallthrough //next case will be evaluated and executed
	case choice < 114514:
		fmt.Println("3")
	default:
		fmt.Println("no match")
	}
}

func singleBranch() {
	var age int
	fmt.Print("Input your age...>>")
	_, _ = fmt.Scanf("%d", &age)
	if age > 18 {
		fmt.Println("You should be responsible for yourself;")
	}
}

func doubleBranch() {
	var age int
	fmt.Print("Input your age>>")
	_, _ = fmt.Scanf("%d", &age)
	if age > 18 {
		fmt.Println("You should be responsible for yourself;")
	} else {
		fmt.Println("放过你了")
	}
}

func multiBranch() {
	var point int
	fmt.Print("Input your point>>")
	_, _ = fmt.Scanln(&point)
	if point >= 90 {
		fmt.Println("pt larger than 90")
	} else if point < 90 && point > 80 {
		fmt.Println("pt in (80,90)")
	} else {
		fmt.Println("others")
	}
}
