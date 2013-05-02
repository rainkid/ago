package admin

import (
	"dogo"
	"strconv"
)

type Errors struct {
	dogo.Controller
	errors map[int]string
}

func (c *Errors) Init(){
	c.errors = map[int]string{
		1001:"用户名不能为空.",
		1002:"密码不能为空.",
		1003:"用户名或者密码不正确.",
	}
}

func (c *Errors) Index() {
	code,_ := strconv.Atoi(c.GetInput("code"))
	c.Set("msg", c.errors[code])
}
