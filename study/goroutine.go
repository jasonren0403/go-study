package main

import (
	"fmt"
	_ "runtime"
	"sync"
	"time"
)
var(
	myMap = make(map[uint]uint,10)
	// should declare a synchornized lock
	lock sync.Mutex
)

func test(n uint){
	var (
		res uint=1
		j uint = 1
	)
	for j=1;j<=n;j++{
		res *= j
	}
	lock.Lock() // lock before write
	myMap[n] = res
	lock.Unlock() // unlock after write
}

func main(){
	//test()
	//fmt.Println("cpu number:",runtime.NumCPU())
	//set CPU numbers, no need to run this after go 1.8
	//runtime.GOMAXPROCS(2)
	for i:=1;i<=200;i++{
		go test(uint(i)) // starts a 'goroutine'
		// warning: only starting this without lock will cause concurrent map writes
	}
	time.Sleep(10*time.Second)
	lock.Lock() // lock before read
	for i,v:=range myMap{
		fmt.Printf("map[%d]=%d\n",i,v)
	}
	lock.Unlock() // unlock after read
	//for i:=0;i<10;i++{
	//	fmt.Println("hello world in main()",i)
	//	time.Sleep(time.Second)
	//}
}
