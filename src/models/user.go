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

	cookieStr string
}

func NewUserModel() *User {
	return &User{
		Model: Model{TableName: "admin_user", PrimaryKey: "uid"},
		hash:  "A#a&(_=)",
		code:  1000,
	}
}

func (u *User) WithCookie(cookieStr string) *User {
	u.cookieStr = cookieStr
	return u
}

func (u *User) WithHash(hash string) *User {
	u.hash = hash
	return u
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

	flag, s := u.Password(utils.ItoString(password), fmt.Sprintf("%s", result["hash"]))
	if !flag {
		return u.code + 3, ""
	}

	if !bytes.Equal([]byte(s), utils.ItoByte(result["password"])) {
		return u.code + 4, ""
	}

	str := fmt.Sprintf("%d|%s|%s", result["uid"], result["username"], result["hash"])
	cstr, err := utils.Encrypt(str, u.hash)

	if err != nil {
		return u.code, ""
	}
	return 0, cstr
}

func (u *User) CheckPasswd(password string) (bool, string) {
	if len(password) == 0 {
		return false, "原始密码不能为空."
	}

	flag, info := u.GetLoginUser()
	if !flag {
		return false, "用户未登录."
	}

	result, err := u.Where("username =? ", info[1]).Get()
	if err != nil {
		return false, "获取用户信息失败."
	}

	flag, str := u.Password(password, fmt.Sprintf("%s", result["hash"]))
	if !flag {
		return false, "密码加密失败."
	}
	if !bytes.Equal([]byte(str), utils.ItoByte(result["password"])) {
		return false, "原始密码不正确."
	}
	return true, ""
}

func (u *User) Password(password, hash string) (bool, string) {
	temp := utils.MD5(password)
	str := utils.MD5(fmt.Sprintf("%x", temp) + hash)
	return true, string(fmt.Sprintf("%x", str))
}

func (u *User) GetLoginUser() (bool, []string) {
	if u.cookieStr == "" {
		return false, nil
	}
	destr, err := utils.Decrypt(u.cookieStr, u.hash)
	if destr == "" || err != nil {
		return false, nil
	}

	info := strings.Split(destr, "|")
	if len(info) != 3 {
		return false, nil
	}
	return true, info
}

func (u *User) IsLogin() (bool, map[string]interface{}) {
	flag, info := u.GetLoginUser()
	if !flag {
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
	username, ulen := u.GetData("username")
	password, plen := u.GetData("password")
	email, elen := u.GetData("email")
	groupid, _ := u.GetData("groupid")
	r_password, rplen := u.GetData("r_password")

	if username != nil && ulen == 0 {
		return -1, "用户不能为空."
	}
	if elen == 0 {
		return -1, "邮箱不能为空."
	}

	if plen == 0 || rplen == 0 {
		return -1, "密码不能为空."
	}

	if !bytes.Equal(utils.ItoByte(password), utils.ItoByte(r_password)) {
		return -1, "密码与确认密码不一致."
	}
	delete(u.Data, "r_password")

	if username != nil {
		u.Data["username"] = utils.ItoString(u.Data["username"])
	}
	if password != nil {
		u.Data["hash"] = u.hash
		flag, password := u.Password(utils.ItoString(u.Data["password"]), u.hash)
		if !flag {
			return -1, "密码操作失败."
		}
		u.Data["password"] = password
	}
	if email != nil {
		u.Data["email"] = utils.ItoString(u.Data["email"])
	}
	if groupid != nil {
		u.Data["groupid"] = utils.ItoString(u.Data["groupid"])
	}

	return 0, ""
}

func (u *User) LoginValid() (int64, string) {
	if _, length := u.GetData("username"); length == 0 {
		return u.code + 1, ""
	}

	if _, length := u.GetData("password"); length == 0 {
		return u.code + 2, ""
	}
	return 0, ""
}
