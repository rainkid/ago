package models

import (
	"dogo"
	"fmt"
	lib "lib"
)

type User struct {
	dogo.Model
	hash string
}

func NewUserModel() *User {
	user := &User{}
	user.TableName = "admin_user"
	user.hash = "A#a&(_=)"
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
		str := fmt.Sprintf("%d|%s|%s", uid, username, hash)
		cstr, err := lib.Encrypt(str, u.hash)
		destr, err := lib.Decrypt(cstr, u.hash)

		fmt.Println(destr)

		if err != nil {
			return 1004, ""
		}
		return 0, cstr
	}
	return 1003, ""
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
		return 1001, ""
	}

	if ok, _ := data["password"]; ok == nil {
		return 1002, ""
	}
	return 0, ""
}
