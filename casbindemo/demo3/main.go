package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

/**
	这个主要是学习管理 API。也就是https://casbin.org/docs/zh-CN/management-api这篇文档的内容
 */
var (
	modelFile = "E:/GoProjects/src/dev/casbindemo/demo1/rbac_model2.conf"
	policyFile = "E:/GoProjects/src/dev/casbindemo/demo1/rbac_policy.csv"
)

func main() {
	a,err := gormadapter.NewAdapter("mysql","root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8",true)
	if err != nil{
		fmt.Println(a)
		fmt.Println("gormadapter.NewAdapter ",err.Error())
	}
	e,err := casbin.NewEnforcer(modelFile,a)
	if err != nil{
		fmt.Println(e)
		fmt.Println("casbin.NewEnforcer Error",err.Error())
	}


	fmt.Println(e.GetAllActions())

}


func base(e *casbin.Enforcer)  {
	// 1.获取当前casbin的适配器
	adapter := e.GetAdapter()
	fmt.Println(adapter)

	// 2.获取当前的Model相关属性
	model := e.GetModel()
	fmt.Println(model)
	// 相当于是执行：SELECT DISTINCT(v1) FROM casbin_rule WHERE p_type = 'p';
	fmt.Println(e.GetAllObjects())
	// 相当于是执行：SELECT v1 FROM casbin_rule WHERE p_type = 'g';
	fmt.Println(e.GetAllRoles())
}

func policy(e *casbin.Enforcer)  {
	// 1.添加规则(权限)的几种方法
	// ①、添加一条规则(或是权限)。在rbac里面的AddPermissionForUser()方法其实也就是调用了AddPolicy()这个方法
	e.AddPolicy("/user","GET")
	// 虽然可以指定ptype，但是这里也只能写p，可以将AddNamedPolicy()理解为AddPolicy()方法的一个别名。
	e.AddNamedPolicy("p","/user","POST")
	// ②、这个是将权限进行分组。或者更好理解的是，将用户和角色对应起来。比如下面我们就可以将group1认为是一个用户，data2_admin是一个角色，那这样的意思就是group1是data2_admin这个用户角色
	// 在rbac里面的AddRoleForUser()方法其实也就是调用了AddGroupingPolicy()这个方法
	e.AddGroupingPolicy("group1", "data2_admin")
	// 这个方法和AddNamedPolicy()方法类似，虽然也可以指定ptype，但是这个参数值也只能写g。这个方法的功能和AddGroupingPolicy()一样
	e.AddNamedGroupingPolicy("g","group2","data3_admin")

	// 2.获取规则(权限)的几种方法
	// ①、获取单独权限
	policy1 := e.GetPolicy()
	fmt.Println(policy1)
	// 虽然可以指定ptype，但是这里也只能写p，可以将GetNamedPolicy()理解为GetPolicy()方法的一个别名。
	policy2 := e.GetNamedPolicy("p")
	fmt.Println(policy2)
	// ②、获取权限分组列表
	policy3 := e.GetGroupingPolicy()
	fmt.Println(policy3)
	// 虽然可以指定ptype，但是这里也只能写g，可以将GetNamedGroupingPolicy()理解为GetGroupingPolicy()方法的一个别名。
	policy4 := e.GetNamedGroupingPolicy("g")
	fmt.Println(policy4)

	// ③、获取指定角色或用户的权限
	// 在RBAC的API中，GetPermissionsForUser()方法就是调用的GetFilteredPolicy()方法。
	policy5 := e.GetFilteredPolicy(0,"zhangsan")
	fmt.Println(policy5)

	// 获取指定用户是什么角色
	policy6 := e.GetFilteredGroupingPolicy(0,"group1")
	fmt.Println(policy6)

	// GetFilteredNamedGroupingPolicy()方法就是GetFilteredGroupingPolicy()方法的一个别名，可以单独指定ptype，不过这个只能传递参数g
	policy7 := e.GetFilteredNamedGroupingPolicy("g",0,"group1")
	fmt.Println(policy7)

	// ④、判断是否有权限
	boolValue1 := e.HasPolicy("/user","POST")//这里传递的参数，是严格区分大小写的。如果将POST写成post，就会返回false
	fmt.Println(boolValue1)

	boolValue2 := e.HasNamedPolicy("p","/user","POST")
	fmt.Println(boolValue2)

	boolValue3 := e.HasGroupingPolicy("group1","data2_admin")
	fmt.Println(boolValue3)

	boolValue4 := e.HasNamedGroupingPolicy("g","group1","data2_admin")
	fmt.Println(boolValue4)

	// ⑤、删除权限
	/**
		RemoveGroupingPolicy
		removeFilteredPolicy
		RemovePolicy
		RemoveFilteredPolicy
		RemoveNamedPolicy
		RemoveFilteredNamedPolicy
		RemoveFilteredGroupingPolicy
		RemoveNamedGroupingPolicy
		RemoveFilteredNamedGroupingPolicy
	 */

}
