package admin

import (
	"fmt"
	// utils "libs/utils"
	models "models"
)

type User struct {
	AdminBase
}

func (c *User) Index() {
	user := models.NewUserModel()
	users, err := user.Gets()
	if err == nil {
		c.Assign("users", users)
	}
}

func (c *User) Edit() {
	id := c.GetInput("id")
	user := models.NewUserModel()
	auser, err := user.Where("uid = ? ", id).Get()
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
	values := c.GetPosts([]string{"email", "password", "r_password", "groupid"})
	user := models.NewUserModel()
	user.SetData(values)

	if code, msg := user.Valid(); code != 0 {
		c.Json(code, msg, nil)
		return
	}
	user.Insert()
	c.Json(0, "操作成功", nil)
}
