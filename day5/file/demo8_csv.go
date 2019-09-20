package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

/**
	csv文件的导出和读取
 */

type Article struct {
	Id int
	Title string
	Content string
}

func main() {
	//将内容写入到csv文件里面。下面的路径会定位到GOPATH目录下面。也就是说，articlecsv.csv是在%GOPATH%目录下面。
	csvFile,err := os.Create("articlecsv.csv")
	if err != nil{
		panic(err)
	}
	defer csvFile.Close() //这个需要在上面的判断下面。如果已经发生了错误，最后还执行这个，就会有异常

	allArticles := []Article{
		{Id:1,Title:"title1",Content:"content1"},
		{Id:2,Title:"title2",Content:"content2"},
		{Id:3,Title:"title3",Content:"content3"},
		{Id:4,Title:"title4",Content:"content4"},
	}

	writer := csv.NewWriter(csvFile)
	for _, article := range allArticles {
		line := []string{strconv.Itoa(article.Id),article.Title,article.Content}
		err := writer.Write(line) //还有一个方法WriteAll()方法，只不过这个方法接收的一个二维数组字符串格式的参数
		if err != nil{
			panic(err)
		}
	}
	writer.Flush()


	//读取csv文件内容
	file,err := os.Open("articlecsv.csv")
	if err != nil{
		panic(err)
	}

	defer file.Close() //这个需要在上面的判断下面。如果已经发生了错误，最后还执行这个，就会有异常

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	recod,err := reader.ReadAll() //得到一个二维数组
	fmt.Println(recod)
	if err != nil{
		panic(err)
	}
	var articles []Article

	for _, item := range recod {
		id,_ := strconv.ParseInt(item[0],0,0)
		article := Article{
			Id:int(id),
			Title:item[1],
			Content:item[2],
		}
		articles = append(articles,article)
	}
	for _, article := range articles {
		fmt.Printf("id:%d,title:%v,content:%v \n",article.Id,article.Title,article.Content)
	}
}