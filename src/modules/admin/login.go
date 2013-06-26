package admin

import (
	"fmt"
	models "models"
)

type Login struct {
	AdminBase
}

func (c *Login) Index() {

}

func (c *Login) Login() {
	values := c.GetInputs([]string{"username", "password"})

	user := models.NewUserModel()
	user.SetData(values)

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
	return
}

func (c *Login) Logout() {
	c.DelCookie("Admin_User")
	c.Redirect("/admin/index/index", nil)
	return
}
