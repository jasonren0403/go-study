package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:888")
	if err != nil {
		fmt.Println("Dial error", err)
		return
	}
	//fmt.Printf("conn succ = %v\n",conn)
	// send single line content
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Input your message >> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read string error")
		}
		line = strings.Trim(line, " \r\n")
		if line == "exit" {
			fmt.Println("client exit")
			break
		}
		_, err = conn.Write([]byte(line + "\n"))
		if err != nil {
			fmt.Println("Send error", err)
		}
	}
}
