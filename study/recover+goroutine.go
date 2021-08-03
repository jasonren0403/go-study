package main

import (
	"fmt"
	"time"
)

func noProblem(){
	for i:=0;i<10;i++{
		fmt.Println("No problem")
		time.Sleep(time.Second)
	}
}

func yesProblem(){
	defer func(){
		if err:=recover();err!=nil{
			fmt.Println("yesProblem:",err)
		}
	}()
	var m map[int]string
	m[0] = "123"
}

func main() {
	go noProblem()
	go yesProblem()

	for q:=0;q<10;q++{
		time.Sleep(time.Second)
		fmt.Println("main()")
	}
}
