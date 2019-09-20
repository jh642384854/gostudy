package main

import "fmt"

type DubleUser struct {
	No int
	Username string
	Age int
	prev *DubleUser  //这个表示指向前一个结点
	next *DubleUser  //这个表示指向下一个结点
}

//添加元素
func Push(headnode *DubleUser,newnode *DubleUser)  {
	//1. 创建一个辅助结点
	temp := headnode
	//遍历节点，找到最后一个节点
	for {
		//2. 先找到该链表的最后这个结点
		if temp.next == nil{
			break
		}
		temp = temp.next   //这一点很重要，保证结点会往下移懂，最后正常循环完毕
	}
	//3. 将newHeroNode加入到链表的最后
	temp.next = newnode
	newnode.prev = temp
}
//按照顺序来插入
func PushOrder(headnode *DubleUser,newnode *DubleUser)  {
	temp := headnode
	flag := false
	//遍历节点，找到符合条件的节点
	for  {
		if temp.next == nil{
			break
		}else if temp.next.No > newnode.No{
			break
		}else if temp.next.No == newnode.No{
			flag = true
		}
		temp = temp.next   //这一点很重要，保证结点会往下移懂，最后正常循环完毕
	}
	if flag{
		fmt.Println("重复NO，不允许添加")
	}else{
		//1.处理新节点的前后节点
		newnode.next = temp.next
		newnode.prev = temp

		//2.处理原始结点的前后节点
		if temp.next != nil{
			temp.next.prev = newnode
		}
		temp.next = newnode
	}

}

//删除元素
func Pop(headnode *DubleUser,no int)  {
	temp := headnode
	flag := false
	//遍历节点，找到要删除的元素
	for  {
		if temp.next == nil{
			break
		}else if temp.next.No == no{
			flag = true
			break
		}
		temp = temp.next   //这一点很重要，保证结点会往下移动，最后正常循环完毕
	}
	if flag{
		temp.next = temp.next.next
		if temp.next != nil{
			temp.next.prev = temp
		}
	}else{
		fmt.Println("没有找到可以删除的元素")
	}
}

//显示数据(正序输出)
func List(headnode *DubleUser)  {
	temp := headnode
	if temp.next == nil{
		fmt.Println("空数据")
		return
	}
	for {
		//fmt.Println(temp.next) //注意：这里打印的是temp.next，而不是temp，这二者是有区别的。
		fmt.Printf("[%d , %s , %d]\n", temp.next.No, temp.next.Username, temp.next.Age)
		fmt.Println()
		temp = temp.next
		if temp.next == nil{
			break
		}
	}
}

//显示数据(倒序输出)
func ListDesc(headnode *DubleUser)  {
	temp := headnode
	if temp.next == nil{
		fmt.Println("空数据")
		return
	}
	//通过下面的forx循环，我们已经找到了链表中的最后一个元素，所以下面输出的时候，就可以直接用temp.No的方式来输出，而不需要用temp.prev.No这样的方式
	for {
		if temp.next == nil{
			break
		}
		temp = temp.next
	}
	for  {
		fmt.Printf("[%d , %s , %d]\n", temp.No, temp.Username, temp.Age)
		temp = temp.prev //将temp结点往前移动
		if temp.prev == nil{
			break
		}
	}
}

func main() {
	doubleUser := &DubleUser{}

	user1 := &DubleUser{
		No:1,
		Username:"zhangsan",
		Age:15,
	}

	user2 := &DubleUser{
		No:2,
		Username:"lisi",
		Age:25,
	}

	user3 := &DubleUser{
		No:3,
		Username:"wangwu",
		Age:25,
	}

	List(doubleUser)

	Push(doubleUser,user1)
	Push(doubleUser,user2)
	Push(doubleUser,user3)

	List(doubleUser)
	fmt.Println()
	Pop(doubleUser,2)
	fmt.Println("逆序输出")
	ListDesc(doubleUser)
}