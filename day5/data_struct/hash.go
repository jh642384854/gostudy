package main

import (
	"errors"
	"fmt"
	"os"
)

/**
	实现过程：
	我们要存储一组数据，这个数据可以是一个用户对象(这里就是一个结构体对象)。这个用户对象包含了用户ID，用户名称等等信息
	我们可以这样来进行存储，通过散列的方式。我们定义一个定长的数组(假定为10)，需要将100个或更多存入这个数组中。
	这个时候你就会想了，长度只有10的数组怎么可能存放100个元素呢？其实我们这10个元素并不是实际存放的这些对象信息，而是这些对象的链表信息。
	我们会将这些用户信息，依据用户ID进行指定的散列算法，得到数组下标，将这些对象信息，散列的存放到长度为10的数组指定下标中。
 */

const HASH_TABLE_MAX  = 10

//定义的结构体对象
type Emploee struct {
 	Id int
 	Username string
 	next *Emploee
}

func (this *Emploee) ShowSelf()  {
	fmt.Printf("显示个人信息,Id:%d，Username:%s \n",this.Id,this.Username)
}

//定义EmploeeLink对象
type EmploeeLink struct {
	Head *Emploee
}

//定义插入链表的方法
func (this *EmploeeLink) InsertLinkVal(emploee *Emploee)  {
	//如果头结点为空，就直接把新节点指向头结点
	if this.Head == nil{
		this.Head = emploee
		return
	}
	//如果头结点不为空，就需要找到合适的位置进行插入(因为我们插入的顺序是有要求的，需要按照用户的id从小到大的规则)
	cur := this.Head
	var prev *Emploee = nil
	for {
		if cur != nil {
			if cur.Id > emploee.Id{
				break
			}
			prev = cur
			cur = cur.next
		}else {
			break
		}
	}
	//下面的这个需要画图就更好理解了
	prev.next = emploee
	emploee.next = cur
}

//显示某个链表的数据
func (this *EmploeeLink) ShowLinkVal(n int)  {
	//判断链表是否为空
	if this.Head == nil{
		fmt.Printf("数组下标为%d的链表内容为空",n)
		return
	}
	//链表不为空，则循环遍历链表
	cur := this.Head
	fmt.Printf("数组下标为：%d的链表内容如下",n)
	for  {
		if cur != nil{
			fmt.Printf("用户id：%d，用户名:%s \t",cur.Id,cur.Username)
			cur = cur.next
		}else{
			break
		}
	}
	fmt.Println()
}

//查询某个对象的值
func (this *EmploeeLink) GetEmploee(uid int) (emploee *Emploee,err error )  {
	var emploeeVal *Emploee
	//先判断是否为空
	if this.Head == nil{
		return emploeeVal,errors.New("没有找到该用户信息")
	}
	//遍历这个下标的元素
	cur := this.Head
	flag := false
	for  {
		if cur != nil{
			if cur.Id == uid{
				flag = true
				break
			}
			cur = cur.next
		}else{
			break
		}
	}
	if flag{
		return cur,nil
	}else{
		return nil,errors.New("没有找到该用户")
	}
}

//定义的hashtable表
type HashTable struct {
	Data [10]EmploeeLink
}

//向HashTable添加元素
func (this *HashTable) InsertHashTable(emploee *Emploee)  {
	index := hashAlgorithm(emploee.Id)
	this.Data[index].InsertLinkVal(emploee)
}

//显示HashTable的值
func (this *HashTable) ListHashTable()  {
	for i:=0;i<HASH_TABLE_MAX ;i++  {
		this.Data[i].ShowLinkVal(i)
		fmt.Println()
	}
}

//根据用户ID来查询用户信息
func (this *HashTable) GetHashValue(uid int) (emploee *Emploee,err error ) {
	index := hashAlgorithm(uid)
	return this.Data[index].GetEmploee(uid)
}

//自己定义的散列算法
func hashAlgorithm (userid int) int {
	return userid % HASH_TABLE_MAX
}

func main() {
	fmt.Println("HashTable的使用")
	var hashTable HashTable
	operator := ""
	userid := 0
	userName := ""
	findUid := 0
	for {
		fmt.Println("add 表示添加雇员")
		fmt.Println("list  表示显示雇员")
		fmt.Println("get  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&operator)
		switch operator {
		case "add":
			fmt.Println("请输入雇员ID")
			fmt.Scanln(&userid)
			fmt.Println("请输入雇员姓名")
			fmt.Scanln(&userName)
			emploee := &Emploee{
				Id:userid,
				Username:userName,
			}
			hashTable.InsertHashTable(emploee)
		case "list":
			hashTable.ListHashTable()
		case "get":
			fmt.Println("请输入用户ID")
			fmt.Scanln(&findUid)
			emploee,err := hashTable.GetHashValue(findUid)
			if err != nil{
				fmt.Println(err.Error())
			}else{
				emploee.ShowSelf()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("命令输入有误")
		}
	}
}
