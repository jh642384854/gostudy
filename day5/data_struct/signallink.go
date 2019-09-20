package main

import "fmt"

/**
	单链表(带头结点)核心要点：
	必须有一个头结点，通过这个头结点来进行元素的添加、删除、查看
 */

type SingalUser struct {
	No int
	Username string
	Age int
	next *SingalUser
}
//往节点插入数据
func InsertLink(hedenode *SingalUser,newnode *SingalUser)  {
	temp := hedenode
	//因为不知道链表有多大，所以需要进行无限for循环进行遍历
	for  {
		if temp.next == nil{ //表示找到最后
			break
		}
		temp = temp.next    //让temp不断的指向下一个结点
	}
	temp.next = newnode
}

//有序插入数据，按照结构体的NO大小来插入
func InsertLinkOrder(hedenode *SingalUser,newnode *SingalUser)  {
	temp := hedenode
	flag := false
	for {
		if temp.next == nil{
			break
		}else if hedenode.next.No > newnode.No{
			break
		}else if hedenode.next.No == newnode.No{
			flag = true
			break
		}
		temp = temp.next
	}
	if flag{
		fmt.Println("重复NO，不允许添加")
	}else{
		newnode.next = temp.next
		temp.next = newnode
	}
}

//从节点中删除数据
func DeleteLink(hedenode *SingalUser,no int)  {
	temp := hedenode
	flag := false
	if temp.next == nil{
		fmt.Println("当前链表没有元素")
		return
	}
	for {
		if temp.next == nil{
			break
		}
		//找到要删除结点的no，和temp的下一个结点的no比较
		if temp.next.No == no{
			flag = true
		}
		temp = temp.next
	}
	if flag{
		temp.next = temp.next.next   //将节点执行下下个节点，跳过中间的这个节点，就相当于是删除了(最后会被垃圾回收)
	}else{
		fmt.Println("要删除的节点不存在")
	}
}
//循环节点
func ListLink(hedenode *SingalUser)  {
	temp := hedenode

	if temp.next == nil{
		fmt.Println("当前链表没有元素")
		return
	}
	for  {
		if temp == nil{
			break
		}
		fmt.Println(temp)
		fmt.Printf("no:%d,username:%v,age:%d \n",temp.No,temp.Username,temp.Age)
		fmt.Println()
		temp = temp.next
	}
}


func main() {
	head := &SingalUser{}

	user1 := &SingalUser{
		No:1,
		Username:"zhangsan",
		Age:15,
	}

	user2 := &SingalUser{
		No:2,
		Username:"lisi",
		Age:25,
	}

	/*InsertLink(head,user1)
	InsertLink(head,user2)*/
	InsertLinkOrder(head,user2)
	InsertLinkOrder(head,user1)


	ListLink(head)
}