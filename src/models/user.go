package models

import (
	"bytes"
	"fmt"
	utils "libs/utils"
	"strings"
)

type User struct {
	Model

	hash string
	code int64
}

func NewUserModel() *User {
	return &User{
		Model: Model{TableName: "admin_user", PrimaryKey: "uid"},
		hash:  "A#a&(_=)",
		code:  1000,
	}
}

func (u *User) Login() (int64, string) {
	username, ulen := u.GetData("username")
	password, plen := u.GetData("password")

	if ulen == 0 || plen == 0 {
		return u.code + 2, ""
	}

	result, err := u.Where("username =? ", utils.ItoString(username)).Get()
	if err != nil {
		return u.code + 3, ""
	}

	temp := utils.MD5(utils.ItoString(password))
	t := utils.MD5(fmt.Sprintf("%x", temp) + fmt.Sprintf("%s", result["hash"]))

	if !bytes.Equal([]byte(string(fmt.Sprintf("%x", t))), utils.ItoByte(result["password"])) {
		return u.code + 4, ""
	}

	str := fmt.Sprintf("%d|%s|%s", result["uid"], result["username"], result["hash"])
	cstr, err := utils.Encrypt(str, u.hash)

	if err != nil {
		return u.code, ""
	}
	return 0, cstr
}

func (u *User) IsLogin(str string) (bool, map[string]interface{}) {
	if str == "" {
		return false, nil
	}
	destr, err := utils.Decrypt(str, u.hash)
	if destr == "" || err != nil {
		return false, nil
	}

	info := strings.Split(destr, "|")
	if len(info) != 3 {
		return false, nil
	}
	result, err := u.Where("uid = ? AND username = ?", info[0], info[1]).Get()

	if err != nil {
		return false, nil
	}

	if !bytes.Equal(utils.ItoByte(result["hash"]), []byte(info[2])) {
		return false, nil
	}
	return true, result
}

func (u *User) Valid() (int64, string) {
	if _, length := u.GetData("username"); length == 0 {
		return u.code + 1, ""
	}

	if _, length := u.GetData("password"); length == 0 {
		return u.code + 2, ""
	}
	return 0, ""
}
