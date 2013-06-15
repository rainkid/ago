package models

import (
	"fmt"
	lib "libs"
)

type User struct {
	Model
	hash string
	code int64
}

func NewUserModel() *User {
	user := &User{}
	user.TableName = "admin_user"
	user.hash = "A#a&(_=)"
	user.code = 1000
	user.Fields = []string{
		"uid",
		"username",
		"password",
		"hash",
		"email",
		"registertime",
		"registerip",
		"groupid",
	}
	return user
}

func (u *User) Login() (int64, string) {
	username := u.Data["username"]
	if username != nil {

		row, _ := u.Where("username", username).Get()

		var uid int
		var username string
		var password string
		var hash string
		var email string
		var registertime int64
		var registerip string
		var groupid int64

		row.Scan(&uid, &username, &password, &hash, &email, &registertime, &registerip, &groupid)
		fmt.Println(uid, username, hash)
		str := fmt.Sprintf("%d|%s|%s", uid, username, hash)
		cstr, err := lib.Encrypt(str, u.hash)
		destr, err := lib.Decrypt(cstr, u.hash)

		fmt.Println(destr)

		if err != nil {
			return u.code + 1, ""
		}
		return 0, cstr
	}
	return u.code + 2, ""
}

func (u *User) IsLogin(str string) bool {
	if str != "" {
		return true
	}
	return false
}

func (u *User) Valid() (int64, string) {

	data := u.GetData()
	if ok, _ := data["username"]; ok == nil {
		return u.code + 3, ""
	}

	if ok, _ := data["password"]; ok == nil {
		return u.code + 4, ""
	}
	return 0, ""
}
