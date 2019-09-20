package main

import (
	"fmt"
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	"github.com/casbin/mysql-adapter"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)
/**
	这个文件主要进行的rbac API的学习。
	也就是https://casbin.org/docs/zh-CN/rbac-api这篇文档的内容
 */
var (
	modelFile = "E:/GoProjects/src/dev/casbindemo/demo1/rbac_model2.conf"
	//policy将会从数据库里面读取
	policyFile = "E:/GoProjects/src/dev/casbindemo/demo1/rbac_policy.csv"
)
func main() {
	a,err := gormadapter.NewAdapter("mysql","root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8",true)
	if err != nil{
		fmt.Println(a)
		fmt.Println("gormadapter.NewAdapter ",err.Error())
	}
	e,_ := casbin.NewEnforcer(modelFile,a)
	e.EnableLog(true)
	err = e.LoadPolicy()
	if err != nil{
		fmt.Println("e.LoadPolicy ",err.Error())
	}
}
/**
	权限相关操作
 */
func casbinPermission(e *casbin.Enforcer)  {
	// 1.给指定的用户添加权限，相当于是执行：INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('p', 'u_admin', '/article', 'GET', '', '', '');
	e.AddPermissionForUser("u_admin","/article","GET")
	// 2.删除某个用户指定的权限。相当于是执行delete from casbin_rule where p_type='p' and v0='u_admin' and v1='/article' and v2='GET'
	bdp1,err := e.DeletePermissionForUser("u_admin","/article","GET")
	if err != nil{
		fmt.Println("DeletePermissionForUser error",err.Error())
	}
	fmt.Println(bdp1)
	// 3.批量删除指定用户的所有权限。相当于是执行：delete from casbin_rule where p_type='p' and v0='u_admin'
	bdp2,err := e.DeletePermissionsForUser("u_admin")
	if err != nil{
		fmt.Println("DeletePermissionsForUser error",err.Error())
	}
	fmt.Println(bdp2)
	// 3.获取指定用户的所有权限，相当于是执行select v0,v1,v2 from casbin_rule where v0='u_admin' and p_type='p'
	permissions := e.GetPermissionsForUser("u_admin")
	fmt.Println(permissions)
	// 4.单纯删除权限记录。这里删除的权限和用户没有关联，单纯的是删除权限这条记录，相当于是执行：delete from casbin_rule where p_type='p' and v1='/article' and v2='GET'
	bdp3,err := e.DeletePermission("/article","GET")
	if err != nil{
		fmt.Println("DeletePermission error",err.Error())
	}
	fmt.Println(bdp3)
	// 5.获取用户的隐式权限。这个不但会直接查找该用户的权限，还会查找该用户的隶属角色所拥有的权限
	// 那下面的结果就是有下面SQL语句：
	// select v0,v1,v2 from casbin_rule where v0='u_admin' and p_type='p' # 直接查询是该用户的权限
	// select v1 from casbin_rule where v0='u_admin' and p_type='g'       # 查找该用户属于哪个角色，根据这个查询的结果来查询该角色所拥有的权限，那就有下面的语句
	// select v0,v1,v2 from casbin_rule where v0 in(XXX) and p_type='p'   #  v0 in(XXX)这里的XXX就是上面一条语句查询的结果
	permissions2,_ := e.GetImplicitPermissionsForUser("u_admin")
	fmt.Println(permissions2)
	// 5.判断某个用户是否有权限，相当于是执行：select * from casbin_rule where p_type='p' and v0='u_admin' and v1='/article' and v2='GET'
	bpu1 := e.HasPermissionForUser("u_admin","/article","GET")
	if bpu1 {
		fmt.Println("user has Permission")
	}else{
		fmt.Println("yser not Permission")
	}
	// 扩展的一个方法：获得隐式权限的用户
	fmt.Println(e.GetImplicitUsersForPermission("/list","GET"))
	// 6.判断是否有权限
	boolVal,err := e.Enforce("admin","/admin","GET")
	if err != nil{
		fmt.Println("Enforce ",err.Error())
	}
	fmt.Println(boolVal)
}

/**
	角色相关操作
 */
func casbinRole(e *casbin.Enforcer)  {
	// 1.将某个用户添加到某个角色中，当然一个用户可以属于多个角色
	e.AddRoleForUser("u_admin","r_admin")  //相当于是执行 INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', 'u_admin', 'r_admin', '', '', '', '');
	e.AddRoleForUser("u_admin","r2_admin") //相当于是执行 INSERT INTO `casbin_rule` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`) VALUES ('g', 'u_admin', 'r2_admin', '', '', '', '');

	// 2.根据用户名称来获取属于哪些角色
	fmt.Println(e.GetRolesForUser("u_admin"))
	// 3.获取指定角色包含哪些用户
	fmt.Println(e.GetUsersForRole("r_admin"))

	// 4.判断用户是否属于某个角色
	fmt.Println(e.HasRoleForUser("u_admin","r_admin"))
	fmt.Println(e.HasRoleForUser("u_admin","r2_admin"))
	// 5.删除角色.这个会删除p_type='g' and v0='r_admin'的记录。也就是会执行delete from casbin_rule where p_type='g' and v0='r_admin'的数据
	boolVal4,err := e.DeleteRole("r_admin")
	if err != nil{
		fmt.Println("e.DeleteRole ",err.Error())
	}
	fmt.Println(boolVal4)

	// 6.当用户有多个角色时候，可以删除用户的某个角色。这个就相当于是执行delete from casbin_rule where p_type='g' and v0='u_admin' and v1='r2_admin'
	boolVal5,err := e.DeleteRoleForUser("u_admin","r2_admin")
	if err != nil{
		fmt.Println("e.DeleteRole ",err.Error())
	}
	fmt.Println(boolVal5)

	// 7.删除指定用户所有角色。这个就会把这个用户绑定的所有角色都删除。相当于是执行delete from casbin_rule where p_type='g' and v0='u_admin'
	boolVal6,err := e.DeleteRolesForUser("u_admin")
	if err != nil{
		fmt.Println("e.DeleteRole ",err.Error())
	}
	fmt.Println(boolVal6)

	// 8.获取所有角色，这个相当于是执行：select v1 from casbin_rule where p_type='g'
	roles := e.GetAllRoles()
	fmt.Println(roles)
}
func casbinTest(e *casbin.Enforcer)  {

	//下面的这种判断方式，依然会返回false
	boolVal2,err := e.Enforce("r_admin","/admin","GET")
	if err != nil{
		fmt.Println("Enforce ",err.Error())
	}
	fmt.Println(boolVal2)

	// 9.删除用户(这个真不知道有啥用？)
	boolVal3,err := e.DeleteUser("u_admin")
	if err != nil{
		fmt.Println("e.DeleteUser ",err.Error())
	}
	fmt.Println(boolVal3)
}

/**
	下面的这个方法有点问题
 */
func mysqlAdaper()  {
	mysqlAdapter := mysqladapter.NewAdapter("mysql","root:jianghua@tcp(127.0.0.1:3306)/godb?charset=utf8")
	casbinEnforce,err := casbin.NewEnforcer(modelFile,mysqlAdapter)
	if err != nil{
		fmt.Println("mysqladapter.NewAdapter Error",err.Error())
	}
	fmt.Println(casbinEnforce)
}

/**
	casbin的中间件拦截器
 */
func casbinInterceptor(e *casbin.Enforcer) gin.HandlerFunc  {
	return func(c *gin.Context) {
		//获取请求的URL
		obj := c.Request.URL.RequestURI()
		//获取请求的方法
		act := c.Request.Method
		//获取用户角色(注意是角色，不是用户名)
		sub := "abc123"
		//判断是否有权限
		boolVal,err := e.Enforce(sub,obj,act)
		if err != nil{
			fmt.Println("casbin Enforce error",err.Error())
		}
		if boolVal{
			fmt.Println("has promission")
			c.Next()
		}else{
			fmt.Println("no promission")
			c.Abort()
		}
	}
}