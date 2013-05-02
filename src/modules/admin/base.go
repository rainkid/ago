package admin

import (
	"dogo"
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
