package model

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginCheck() (User,bool) {
	user := User{
		ID:1,
		Name:"zhangsan",
		Phone:"16102363254",
	}
	return user,true
}