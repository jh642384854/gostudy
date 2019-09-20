package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

/**
	gob二进制文件的导出和读取
 */

 type Post struct {
	 Id int
	 Title string
	 Content string
 }
//存储数据
func store(data interface{},filname string)  {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil{
		panic(err)
	}
	err = ioutil.WriteFile(filname,buffer.Bytes(),0600)
	if err != nil{
		panic(err)
	}
}

 //读取数据
func load(data interface{},filename string)  {
	raw,err := ioutil.ReadFile(filename)
	if err != nil{
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err  = dec.Decode(data)
	if err != nil{
		panic(err)
	}
}

func main() {
	post := Post{
		Id:1,
		Title:"gob demo",
		Content:"gob demo content",
	}
	store(post,"gobstore")
	var postRead Post
	load(&postRead,"gobstore")
	fmt.Println(postRead)
}