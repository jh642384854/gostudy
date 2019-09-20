package main

/**
	结构体、map、slice数据类型转换为json的示例

	注意事项：在将json转换为map、slice、结构体对象的时候，一定要做到类型匹配，不然会出错的。
 */
import (
	"encoding/json"
	"fmt"
)

type User struct {
	Username string `json:"username"`
	Age int `json:"age"`
	Skill []string `json:"skills"`
}

//示例1：将结构体转换为json字符串
func MarshalStruct() string {
	user := User{
		Username:"张三",
		Age:15,
		Skill:[]string{"java","php","go"},
	}
	byteStr,error := json.Marshal(user)
	if error != nil{
		fmt.Printf("结构体转换为json失败")
	}
	return string(byteStr)
}

//示例2：将map转换为json字符串
func MarshalMap() string  {

	var userMap map[string]interface{}
	userMap = make(map[string]interface{})

	userMap["username"] = "李四"
	userMap["age"] = 24
	userMap["skills"] = []string{"java","php","go","Python"}

	byteStr,error := json.Marshal(userMap)
	if error != nil{
		fmt.Printf("map转换为json失败")
	}
	return string(byteStr)
}


//示例3：将slice(切片)转换为json
func Marsha1Slice() string  {

	var userSlice []map[string]interface{}

	map1 := make(map[string]interface{})
	map1["username"] = "wangwu"
	map1["age"] = 23
	map1["skills"] = "linux"
	userSlice = append(userSlice,map1)

	map2 := make(map[string]interface{})
	map2["username"] = "wangwu"
	map2["age"] = 23
	map2["skills"] = []string{"java","php","go","Python"}

	userSlice = append(userSlice,map2)

	byteStr,error := json.Marshal(userSlice)
	if error != nil{
		fmt.Printf("slice转换为json失败")
	}
	return string(byteStr)
}


//json转换为结构体、map、slice的示例
//示例1：将json转换为结构体
func UnMarsha1Struct() User{
	var user User
	structMarshal := MarshalStruct()
	error := json.Unmarshal([]byte(structMarshal),&user) //注意这里要传递地址
	if error != nil{
		fmt.Printf("json转换结构体失败")
		return User{}
	}
	return user
}

//示例2：将json转换为map
func UnMarsha1Map() map[string]interface{} {
	var usermap map[string]interface{}  //注意，这里的map在转换为json的时候，并不需要使用make()来初始化一下。

	mapMarsha1 := MarshalMap()
	err := json.Unmarshal([]byte(mapMarsha1),&usermap) //注意这里要传递地址
	if err != nil {
		fmt.Printf("json转换map失败")
		return nil
	}

	return usermap
}

//示例3：将json转换为slice
func UnMarsha1Slice() []map[string]interface{} {
	var userSlice []map[string]interface{}

	sliceMarsha1 := Marsha1Slice()

	err := json.Unmarshal([]byte(sliceMarsha1),&userSlice) //注意这里要传递地址
	if err != nil {
		fmt.Printf("json转换slice失败")
		return nil
	}
	return userSlice
}

//测试结构体、map、slice转换为json的方法
func TestMarsha1()  {
	structMarshal := MarshalStruct()
	fmt.Println(structMarshal)

	fmt.Println()

	mapMarsha1 := MarshalMap()
	fmt.Println(mapMarsha1)

	fmt.Println()

	sliceMarsha1 := Marsha1Slice()
	fmt.Println(sliceMarsha1)
}

func TestUnMarsha1()  {
	user := UnMarsha1Struct()
	fmt.Println(user)

	fmt.Println()

	userMap := UnMarsha1Map()
	fmt.Println(userMap)

	fmt.Println()

	userSlice := UnMarsha1Slice()
	fmt.Println(userSlice)

}

func main() {

	TestMarsha1()

	fmt.Println()

	TestUnMarsha1()

}
