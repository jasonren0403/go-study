package main

import (
	"flag"
	"fmt"
	"os"
)

func main_usage(){
	fmt.Println("The length of argument is",len(os.Args))
	for i:=0;i<len(os.Args);i++{
		fmt.Printf("args[%d]:%s\n",i,os.Args[i])
	}
}

func main(){
	// use package 'flag' to support param
	var user string
	var port int
	var pwd string
	var host string
	flag.StringVar(&user,"u","","username, defaults to empty")
	flag.StringVar(&pwd,"pwd","","password, defaults to empty")
	flag.StringVar(&host,"h","localhost","host, defaults to localhost")
	flag.IntVar(&port,"port",3306,"port, defaults to 3306")
	// Notice: parse them!
	flag.Parse()
	fmt.Printf("User = %v, Port = %v, pwd = %v, host = %v",user,port,pwd,host)
}
