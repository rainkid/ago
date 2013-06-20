package admin

import (
	"dogo"
	utils "libs/utils"
	models "models"
)

type AdminBase struct {
	dogo.Controller
}

func (c *AdminBase) CheckLogin() {
	user := models.NewUserModel()
	cookieStr := c.GetCookie("Admin_User")

	if user.IsLogin(cookieStr) == false {
		c.Redirect("/admin/login/index", nil)
	}
}

func (c *AdminBase) InitParams() {
	adminroot := utils.GetConfig("app", "adminroot")
	c.Assign("adminroot", adminroot)
	c.Assign("token", c.GetToken())
}
