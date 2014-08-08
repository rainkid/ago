package admin

import (
	"fmt"
	// s "libs/service"
	models "models"
	// "reflect"
)

type Login struct {
	AdminBase
}

func (c *Login) Index() {
	/*userService := s.NewUserService()
	mWhere := make(map[string]interface{})
	mWhere["uid"] = []string{"=", "5"}

	mOrderBy := make(map[string]interface{})

	users := userService.GetList(mParams, 0, 20, mOrderby)*/

	// mOrderBy := make(map[string]interface{})
	// mOrderBy["uid"] = "DESC"
	/*mData := make(map[string]interface{})
	mData["username"] = "test"
	mData["password"] = "test)()()()"
	mData["email"] = "test@gmail.com"
	mData["hash"] = "hash000"
	mData["registertime"] = "1388059054"
	mData["registerip"] = "127.0.0.1"
	mData["groupid"] = "2"
	//
	auser, err := bd.Insert(mData)
	if err != nil {
		fmt.Println(err)
	}*/

	/*_, err = bd.DeleteBy(mWhere)
	if err != nil {
		fmt.Println(err)
	}*/
	/*rt := reflect.TypeOf(mWhere["uid"])
	fmt.Println(rt.Kind() == reflect.Slice)
	fmt.Println(reflect.ValueOf(mWhere["uid"]))*/
	// fmt.Println(auser)
}

func (c *Login) Login() {
	values := c.GetInputs([]string{"username", "password"})

	user := models.NewUserModel()

	if code, _ := user.LoginValid(values); code != 0 {
		c.Redirect(fmt.Sprintf("/admin/errors/index?code=%d", code), nil)
	}

	// models.User.Login()
	code, msg := user.Login(values)

	if code != 0 {
		c.Redirect(fmt.Sprintf("/admin/errors/index?code=%d", code), nil)
	}
	c.SetCookie("Admin_User", msg, 60*60*24)
	c.Redirect("/admin/index/index", nil)
}

func (c *Login) Logout() {
	c.DelCookie("Admin_User")
	c.Redirect("/admin/index/index", nil)
}
