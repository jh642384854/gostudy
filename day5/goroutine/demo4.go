package main

import "fmt"

/**
	向channel里面写入各种数据类型
 */

//例子1：将int存入channel，并通过for range方式来进行遍历
func IntChannel()  {

	intChan := make(chan int,5)
	intChan <- 10
	intChan <- 11
	intChan <- 15

	//关闭channel
	close(intChan)
	/**
		遍历channel只能通过for  range的方式来实现。遍历channel的时候需要注意以下几点：
		①：通过for range遍历的时候，并没有key，这点需要特别注意。有以下两个细节问题需要注意：
			1) 在遍历时，如果 channel 没有关闭，则回出现 deadlock 的错误
			2) 在遍历时，如果 channel 已经关闭，则会正常遍历数据，遍历完后，就会退出遍历。
		②：在使用for range遍历的时候，需要先关闭channel，这个是前提条件。使用内置函数 close 可以关闭 channel, 当 channel 关闭后，就不能再向 channel 写数据了，但是仍然
可以从该 channel 读取数据
	 */
	for v := range intChan{
		fmt.Println("channel value:",v)
	}
}
//例子2：将string存入channel
func StringChannel()  {

	stringChan := make(chan string,3)
	stringChan <- "go"
	stringChan <- "java"
	stringChan <- "php"

	close(stringChan)

	for v := range stringChan{
		fmt.Println("stringChan value is ",v)
	}
}

//例子3：将map存入channel
func MapChannel()  {
	mapChan := make(chan map[string]string,3)
	map1 := make(map[string]string,3)
	map1["username"] = "zhangsan"
	map1["skills"] = "java,php,go"
	map1["address"] = "wuhan"

	map2 := make(map[string]string,3)
	map2["username"] = "lisi"
	map2["skills"] = "java,php,go,css"
	map2["address"] = "beijing"

	mapChan <- map1
	mapChan <- map2

	close(mapChan)
	for v := range mapChan {
		fmt.Println("mapChan value is :",v)
	}
}
//例子4：将结构体存入channel
type User struct {
	Username string
	Age int
}
func StructChannel()  {
	structChan := make(chan User,3)
	user1 := User{"zhangsan",25}
	user2 := User{"lisi",24}
	user3 := User{"zhaoliu",37}

	structChan <- user1
	structChan <- user2
	structChan <- user3

	close(structChan)

	for v := range structChan {
		fmt.Println("structChan value is :",v)
	}
}

//例子5：将指针存入channel
func PointChannel()  {
	pointChan := make(chan *User,3)

	user1:= &User{"zhangsan",25}
	user2 := &User{"lisi",24}
	user3 := &User{"zhaoliu",37}

	pointChan <- user1
	pointChan <- user2
	pointChan <- user3

	close(pointChan)

	for v := range pointChan {
		fmt.Println("pointChan value is :",v)
	}
}

//例子6：将空接口存入channel(即任意数据类型)
func InterfaceChannel()  {
	interfaceChan := make(chan interface{},3)

	intNum := 3
	str := "zhangsan"
	user := User{"zhaoliu",37}

	interfaceChan <- intNum
	interfaceChan <- str
	interfaceChan <- user

	//先取出两个数据
	<- interfaceChan
	<- interfaceChan
	//获取user信息
	userinfo := <- interfaceChan
	fmt.Println("userinfo username is :",userinfo.(User).Username)//这里使用了类型断言来做转换
}

//例子7：将slice存入到channel
func SliceChannel()  {
	sliceChannel := make(chan []string,3)

	arr := [...]string{"java","php","go","python"}

	slice1 := []string{"str1","str2"}
	slice2 := make([]string,2)
	slice3 := arr[1:3]

	sliceChannel <- slice1
	sliceChannel <- slice2
	sliceChannel <- slice3

	close(sliceChannel)
	for v := range sliceChannel {
		fmt.Println("sliceChan value is :",v)
	}
}

func main() {
	IntChannel()
	fmt.Println()

	StringChannel()
	fmt.Println()

	MapChannel()
	fmt.Println()

	StructChannel()
	fmt.Println()

	PointChannel()
	fmt.Println()

	InterfaceChannel()
	fmt.Println()

	SliceChannel()
	fmt.Println()
}