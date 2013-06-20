package admin

import (
	"dogo"
	"strconv"
)

type Errors struct {
	dogo.Controller
	errors map[int]string
}

func (c *Errors) Init() {
	c.errors = map[int]string{
		1000: "处理异常",
		1001: "用户名不能为空.",
		1002: "密码不能为空.",
		1003: "用户不存在.",
		1004: "用户名或者密码不正确.",
	}
}

func (c *Errors) Index() {
	code, _ := strconv.Atoi(c.GetInput("code"))
	c.Assign("msg", c.errors[code])
}
