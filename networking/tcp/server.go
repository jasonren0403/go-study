package main

import (
	"fmt"
	"io"
	"net"
)

func process(c net.Conn){
	defer c.Close()

	for{
		buf := make([]byte,512)
		n,err := c.Read(buf) // read from conn, blocks the connection
		if err==io.EOF{
			fmt.Println("Client exited")
			return
		}
		// Include n for content read!
		if n>0{
			fmt.Print("from client[",c.RemoteAddr(),"]:")
			fmt.Print(string(buf[:n]))
		}
	}
}

func main() {
	fmt.Println("Start listening")
	listen,err := net.Listen("tcp","127.0.0.1:888")
	if err!=nil{
		fmt.Println("listen error",err)
		return
	}
	//fmt.Printf("listen=%v\n", listen)
	defer listen.Close()

	for{
		// wait for client
		conn, err := listen.Accept()
		if err!=nil{
			fmt.Println("Accept() errored",err)
		}else{
			fmt.Printf("Accept one conn=%v\n", conn)
		}
		// open a goroutine for client
		fmt.Println("Server is waiting for your[",conn.RemoteAddr().String(),"] input")
		go process(conn)
	}
}

