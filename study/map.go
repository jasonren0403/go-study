package main

import (
	"encoding/json"
	"fmt"
)
type PersonInfo struct {
	ID      int `json:"id"`  // <- use `` to specify "tags" for json marshal return
 	Name    string `json:"name"`
	Address string `json:"address"`
}

func main(){
	var a map[string]string
	a = make(map[string]string,10)
	a["a1"] = "12345"
	myMap := map[string]PersonInfo{
		"1234": {
			ID:      114514,
			Name:    "Jack",
			Address: "abcd",
		},
	}
	val, ok := myMap["1234"]
	if ok {
		fmt.Println(val)
	}
	// struct usage
	var person PersonInfo = PersonInfo{0, "a", "b"}
	person.Address = "abcde"
	//fmt.Println(person)
	jsonStr,err := json.Marshal(person)  // json.Marshal returns ([]byte,err)
	if err!=nil{
		fmt.Println("处理错误",err)
	}else {
		fmt.Println(string(jsonStr))
	}
	// use pointer and new
	var p2 *PersonInfo = new(PersonInfo)
	(*p2).Address = "addr"
	// it's ok to omit *
	p2.ID = 3
	var p3 *PersonInfo = &PersonInfo{}
	(*p3).Name = "Jason"
	p3.Address = "cacp"

}
