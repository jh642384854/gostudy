package other

//定义接口
type IUserService interface {
	GetUserByUid() int
	DeleteUserByUid() int
}

//定义接口实现类
type UserService struct {
	Name string
}
//创建一个构造函数，这个构造函数返回值是上面定义的接口类型，那这样的话，接口实现类必须要实现接口的所有方法
func NewUserService() IUserService {
	//这里要返回地址
	return &UserService{Name:"zhangsan"}
}

func (u *UserService) GetUserByUid() int {
	return 1
}

func (u *UserService) DeleteUserByUid() int {
	return 1
}