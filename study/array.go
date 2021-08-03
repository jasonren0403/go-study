package main

import (
	"fmt"
)

func main(){
	arr := [3]int{1, 1, 42}
	arr2 := [...]int{3,4,5}
	arr3 := [...]int{2:400,0:124,1:345}
	printArray(arr)
	printArray(arr2)
	printArray(arr3)
	// 使用range简洁遍历数组，返回key,value
	for i, v := range arr {
		fmt.Println(i, "=>", v)
	}
	// 使用len(arr)获取数组长度
	// [m:n]创建切片数组(sliced)
	var sliced /*[]int*/ = arr[1:]
	// make(type,[size,[capacity]])
	//var sliced = make([]int,4,6)
	//var sliced = []int{1,1,4}
	for _, v := range sliced {
		fmt.Print(v, " ")
	}
	// 向切片中添加元素
	sliced = append(sliced, 114514, 7)
	// 第二个参数可以为切片，需要加三个点打散参数
	sliced = append(sliced, sliced...)
}

func printArray(arr [3]int){
	for i:=0;i<len(arr);i++{
		fmt.Println("i:",i,arr[i])
	}
}
