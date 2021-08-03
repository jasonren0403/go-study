package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func writeOperation(str string,path string){
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err!=nil{
		fmt.Println("Error ,quit")
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, _ = writer.WriteString(str)
	_ = writer.Flush()
}

func read2(){
	content,err1 := ioutil.ReadFile("F:\\code\\go\\study\\proj1\\study\\f1")
	if err1 != nil{
		fmt.Println("error",err1)
	}
	fmt.Printf("%v",string(content))
}

func main() {
	f, err := os.Open("F:\\code\\go\\study\\proj1\\study\\f1")
	if err != nil {
		fmt.Println("error", err)
	}
	defer f.Close()
	// reader with buffer(4096)
	reader := bufio.NewReader(f)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束")
	writeOperation("12345","F:\\code\\go\\study\\proj1\\study\\f1")
	//fmt.Println("Open successfully",f)
	//Open successfully &{0xc000074780} , is a pointer
	//err = f.Close()
	//if err!=nil{
	//	fmt.Println("Close file error",err)
	//}
}
