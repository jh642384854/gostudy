package main

import (
	"encoding/json"
	"fmt"
	"github.com/chenhg5/collection"
	"strconv"
	"strings"
)

func main() {
	data := make(map[string][]int)
	data["ids"] = []int{1,2,3}
	jsondata,_ := json.Marshal(data)
	data2 := make(map[string][]int)
	if err := json.Unmarshal(jsondata,&data2); err != nil{
		fmt.Println(err.Error())
	}
	fmt.Println(string(jsondata))
	fmt.Println(data2["ids"])
	idstr := ""
	for _, value := range data2["ids"] {
		idstr += strconv.Itoa(value) + ","
	}
	fmt.Println(strings.TrimRight(idstr,","))
	fmt.Println(collection.Collect(data2["ids"]).ToJson())
}