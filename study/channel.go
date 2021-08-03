package main

import (
	"fmt"
	"math"
)

type PersonInfo2 struct {
	ID      int
	Name    string
	Address string
}

/* App 1 Data */
func writeData(intChan chan int) {
	for i := 0; i < 50; i++ {
		intChan <- i
		//time.Sleep(time.Second)
	}
	close(intChan)
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intChan
		if !ok {
			break
		}
		//time.Sleep(time.Second)
		fmt.Println("Read", v, "from readData()")
	}
	exitChan <- true
	close(exitChan)
}

func basic() {
	var IntChan chan int
	//var IntChan chan<- int   write only
	//var IntChan <-chan int   read only
	//var person chan PersonInfo
	//var person2 chan *PersonInfo
	IntChan = make(chan int, 3)
	fmt.Println(IntChan)  //0xc000078080
	fmt.Println(&IntChan) //0xc000006028
	IntChan <- 3          // write a constant
	var num = 3
	IntChan <- num // write a variable
	fmt.Printf("len %v,cap %v\n", len(IntChan), cap(IntChan))
	IntChan <- 114514
	//IntChan<- 1919810 // write too much will cause overflow
	num = <-IntChan
	fmt.Println(num) // 3

	IntChan = make(chan int, 100)
	for i := 0; i < 100; i++ {
		IntChan <- i * 2
	}
	close(IntChan) // close the channel before iteration
	for v := range IntChan {
		fmt.Println("v=", v)
	}
}

func chan_with_all_types() {
	var allChan chan interface{}
	allChan = make(chan interface{}, 3)
	allChan <- 10
	allChan <- "str"
	allChan <- PersonInfo2{Name: "Jason", ID: 114514, Address: "addr"}
	close(allChan) // cannot write to this channel anymore
	<-allChan
	<-allChan
	person := <-allChan
	//fmt.Printf("%v %T\n ",person,person) // interface{} with no methods
	a := person.(PersonInfo2)
	fmt.Println(a.Name)
}

func main() {
	// application 1 - read and write data
	var (
		chanInt  = make(chan int, 50)
		chanExit = make(chan bool, 1)
	)
	go writeData(chanInt)
	go readData(chanInt, chanExit)

	for {
		_, ok := <-chanExit
		if !ok {
			break
		}
	}
	// application 2 - count primes
	intChan := make(chan int, 1000)
	primeChan := make(chan int, 2000)

	exitChan := make(chan bool, 4)

	go putNum(intChan,100000)

	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan
		}
		close(primeChan)
	}()

	for {
		res, ok := <-primeChan
		if !ok {
			break
		}
		fmt.Println("Prime number is", res)
	}

}

/* App 2 */
func putNum(intChan chan int,maxnum int) {
	for i := 1; i <= maxnum; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	var (
		num  int
		flag bool
		ok   bool
	)
	for {
		num, ok = <-intChan
		if !ok {
			break
		}
		flag = isPrime(num)
		if flag {
			primeChan <- num
		}
	}
	fmt.Println("primeNum routine")
	exitChan <- true
}

func isPrime(n int) bool {
	flag := true
	for i := 2; i < int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			flag = false
			break
		}
	}
	return flag
}
