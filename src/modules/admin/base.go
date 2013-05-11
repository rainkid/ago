package admin

import (
	"dogo"
	libs "libs"
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

func (c *AdminBase) Assigns() {
	adminroot := libs.GetConfig("app", "adminroot")
	c.Set("adminroot", adminroot)
	c.Set("token", c.GetToken())
}
