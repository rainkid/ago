package admin

import (
	"dogo"
	"fmt"
	models "models"
)

type Login struct {
	dogo.Controller
}

func (c *Login) Index() {
	c.Set("token", c.GetToken())
}

func (c *Login) Login() {
	params := []string{"username", "password"}

	values := c.GetPost(params)

	user := models.NewUserModel()
	user.Sets(values)

	if code, _ := user.Valid(); code != 0 {
		c.Redirect(fmt.Sprintf("/admin/errors/index?code=%d", code), nil)
	}

	// models.User.Login()
	code, msg := user.Login()

	if code != 0 {
		c.Redirect(fmt.Sprintf("/admin/errors/index?code=%d", code), nil)
	}
	c.SetCookie("Admin_User", msg, 60*60*24)
	c.Redirect("/admin/index/index", nil)
}
