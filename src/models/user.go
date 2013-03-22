package models

import (
	// "fmt"
	"dogo"
)

type User struct{
	dogo.Model
}

func NewUserModel() *User{
	user := &User{}
	user.Table("userinfo");
	return user
}

func (u *User) Valid() (int64, string) {
	
	data := u.GetData()
	if ok, _ := data["username"]; ok == nil{
		return -1, "用户名不能为空"
	}

	if ok, _ := data["password"]; ok == nil {
		return -1, "密码不能为空"
	}
	return 0, ""
}