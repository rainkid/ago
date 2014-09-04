package models

import (
	"bytes"
	"fmt"
	"github.com/rainkid/dogo"
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
		code:  1000,
	}
}

func (u *User) Login(mData map[string]string) (int64, string) {
	username, ulen := mData["username"], len(mData["username"])
	password, plen := mData["password"], len(mData["password"])

	if ulen == 0 || plen == 0 {
		return u.code + 2, ""
	}

	result, err := u.Where("username =? ", username).Get()
	if err != nil {
		return u.code + 3, ""
	}

	flag, s := u.Password(password, fmt.Sprintf("%s", result["hash"]))
	if !flag {
		return u.code + 3, ""
	}

	if !bytes.Equal([]byte(s), utils.ItoByte(result["password"])) {
		return u.code + 4, ""
	}

	str := fmt.Sprintf("%d|%s|%s", result["uid"], result["username"], result["hash"])
	cstr, err := utils.Encrypt(str, "12345678")

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
	cookieStr := fmt.Sprintf("%s", dogo.Register.Get("Admin_User_Cookie"))
	if cookieStr == "" {
		return false, nil
	}
	destr, err := utils.Decrypt(cookieStr, "12345678")
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

func (u *User) Valid(mData *map[string]string) (int, string) {
	d := *mData
	username, ulen := d["username"], len(d["username"])
	password, plen := d["password"], len(d["password"])
	_, elen := d["email"], len(d["email"])
	//groupid, _ := d["groupid"], len(d["groupid"])
	r_password, rplen := d["r_password"], len(d["r_password"])
	hash := utils.RandString(8)

	if username != "" && ulen == 0 {
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
	delete(d, "r_password")

	if password != "" {
		d["hash"] = hash
		flag, password := u.Password(d["password"], hash)
		if !flag {
			return -1, "密码操作失败."
		}
		d["password"] = password
	}
	
	return 0, ""
}

func (u *User) LoginValid(mData map[string]string) (int64, string) {
	if _, length := mData["username"], len(mData["username"]); length == 0 {
		return u.code + 1, ""
	}

	if _, length := mData["password"], len(mData["password"]); length == 0 {
		return u.code + 2, ""
	}
	return 0, ""
}
