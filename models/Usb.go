package models

import "fmt"

type Usb interface{
	Start()
	Stop()
}

type Phone struct{

}

type Camera struct{

}

type Computer struct{

}

// implements interface usb
func (p Phone) Start(){
	fmt.Println("The phone start")
}

func (p Phone) Stop(){
	fmt.Println("The phone stop")
}

func (p Camera) Start(){
	fmt.Println("The camera stop")
}

func (p Camera) Stop(){
	fmt.Println("The camera stop")
}

// an interface variable
func (p Computer) Work(usb Usb){
	usb.Start()
	usb.Stop()
}
