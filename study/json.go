package main

import (
	"encoding/json"
	"fmt"
)

type result struct{
	Success  bool     `json:"success"`
	Contents []content `json:"contents"`
}

type content struct{
	IntNum int `json:"intNum"`
	Strings string `json:"string"`
}

func main(){
	var slice []content
	content1 := content{IntNum: 4,Strings:"234235"}
	r := result{
		true,
		make([]content,4),
	}
	r.Contents = append(slice,content1)
	r.Contents = append(r.Contents,content{114514,"abcde"})
	data,err := json.Marshal(r)
	if err!= nil{
		fmt.Println("Error,",err)
	}
	fmt.Println(string(data))

	// map marshalling
	//var a map[string]interface{}
	a := make(map[string]interface{})
	a["success"]=true
	a["array"]=[2]string{"中","c"}
	data, err = json.Marshal(a)
	if err!=nil{
		fmt.Println("Error",err)
	}
	fmt.Println(string(data))
}
