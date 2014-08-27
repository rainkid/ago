package admin

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"github.com/rainkid/dogo"
	"html/template"
	"io"
	utils "libs/utils"
	models "models"
	"strconv"
	"time"
)

type AdminBase struct {
	dogo.Controller
	UserInfo map[string]interface{}
}

func (c *AdminBase) Init() {
	c.InitParams()
	//is not on login page
	if c.ControllerName != "login" {
		c.CheckLogin()

		c.Layout("layout/header.html")
		c.Layout("layout/footer.html")
	}

	c.TplFuncs = template.FuncMap{
		"buffer": bytes.NewBuffer,
		"string": (*bytes.Buffer).String,
	}

	//is ajax disable view
	if c.IsAjax() {
		c.DisableView = true
	}
}

func (c *AdminBase) CheckLogin() {
	user := models.NewUserModel()
	cookieStr := c.GetCookie("Admin_User")
	dogo.Register.Set("Admin_User_Cookie", cookieStr)

	flag, info := user.IsLogin()
	if !flag {
		c.Redirect("/admin/login/index", nil)
	}
	c.UserInfo = info
	c.Assign("UserInfo", info)
}

//get a token
func (c *AdminBase) GetToken() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (c *AdminBase) InitParams() {
	adminroot := utils.GetConfig("app", "adminroot")
	c.Assign("adminroot", adminroot)
	c.Assign("token", c.GetToken())
}
