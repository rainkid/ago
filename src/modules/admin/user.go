package admin

import (
	"bytes"
	"fmt"
	utils "libs/utils"
	models "models"
	"time"
)

type User struct {
	AdminBase
}

func (c *User) Index() {
	user := models.NewUserModel()
	users, err := user.Gets()
	if err == nil {
		for _, user := range users {
			user["registertime"] = time.Unix(utils.ItoInt(user["registertime"]), 0).String()
			if utils.ItoInt(user["groupid"]) == 0 {
				user["groupname"] = "超级管理员"
			} else {
				user["groupname"] = c.getGroupName(utils.ItoByte(user["groupid"]))
			}
		}
		c.Assign("users", users)
	}

}

func (c *User) Add() {
	//group list
	mgroup := models.NewGroupModel()
	groups, err := mgroup.Gets()
	if err == nil {
		c.Assign("groups", groups)
	}
}

func (c *User) Add_post() {
	values := c.GetPosts([]string{"username", "email", "password", "r_password", "groupid"})
	user := models.NewUserModel()
	user.WithHash(utils.RandString(8)).SetData(values)
	if code, msg := user.Valid(); code != 0 {
		c.Json(code, msg, nil)
		return
	}
	user.Insert()
	c.Json(0, "操作成功", nil)
}

func (c *User) Edit() {
	id := c.GetInput("id")
	user := models.NewUserModel()
	auser, err := user.Wherep(id).Get()
	if err == nil {
		c.Assign("auser", auser)
	}
	//group list
	mgroup := models.NewGroupModel()
	groups, err := mgroup.Gets()
	if err == nil {
		c.Assign("groups", groups)
	}
}

func (c *User) Edit_post() {
	id := c.GetInput("id")
	values := c.GetPosts([]string{"email", "password", "r_password", "groupid"})
	user := models.NewUserModel()

	user.SetData(values)
	if code, msg := user.Valid(); code != 0 {
		c.Json(code, msg, nil)
		return
	}
	user.Where("uid = ?", id).Update()
	c.Json(0, "操作成功", nil)
}

func (c *User) Delete() {
	id := c.GetInput("id")
	muser := models.NewUserModel()
	affect, err := muser.Wherep(id).Delete()
	if err != nil {
		c.Json(-1, "操作失败", nil)
		return
	}
	c.Json(0, "操作成功", affect)
}

func (c *User) Passwd_post() {
	values := c.GetInputs([]string{"current_password", "password", "r_password"})
	cookieStr := c.GetCookie("Admin_User")

	user := models.NewUserModel().WithCookie(cookieStr)
	flag, msg := user.CheckPasswd(utils.ItoString(values["current_password"]))
	if !flag {
		c.Json(-1, msg, nil)
		return
	}

	_, info := user.GetLoginUser()
	fmt.Println(info)
	data := map[string]interface{}{"password": values["password"]}
	user.SetData(data).Wherep(info[0]).Update()
	c.Json(0, "修改成功.", nil)
}

func (c *User) getGroupName(groupid []byte) string {
	//group list
	mgroup := models.NewGroupModel()
	groups, err := mgroup.Gets()
	if err != nil {
		return ""
	}
	for _, group := range groups {
		if bytes.Equal(groupid, utils.ItoByte(group["groupid"])) {
			return utils.ItoString(group["name"])
		}
	}
	return ""
}
