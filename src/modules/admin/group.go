package admin

import (
	admin "libs/admin"
	models "models"
	"strings"
)

type Group struct {
	AdminBase
}

func (c *Group) Index() {
	mgroup := models.NewGroupModel()
	groups, err := mgroup.Gets()
	if err == nil {
		c.Assign("groups", groups)
	}
}

func (c *Group) Add() {
	menu := admin.NewMenu()
	c.Assign("menus", menu.Menus)
}

func (c *Group) Add_post() {
	values := c.GetInputs([]string{"name", "descrip"})
	values["rvalue"] = strings.Join(c.GetInputList("rvalue"), ",")
	mgroup := models.NewGroupModel()

	if code, msg := mgroup.Valid(&values); code != 0 {
		c.Json(code, msg, nil)
		return
	}
	mgroup.Insert(values)
	c.Json(0, "操作成功", nil)
}

func (c *Group) Edit() {
	id := c.GetInput("id")
	mgroup := models.NewGroupModel()
	agroup, err := mgroup.Where("groupid = ?", id).Get()

	if err == nil {
		c.Assign("agroup", agroup)
	}

	menu := admin.NewMenu()
	c.Assign("menus", menu.Menus)
}

func (c *Group) Edit_post() {
	id := c.GetInput("id")
	values := c.GetInputs([]string{"name", "descrip"})
	values["rvalue"] = strings.Join(c.GetInputList("rvalue"), ",")
	mgroup := models.NewGroupModel()

	if code, msg := mgroup.Valid(&values); code != 0 {
		c.Json(code, msg, nil)
		return
	}
	_, err := mgroup.Where("groupid=?", id).Update(values)
	if err != nil {
		c.Json(-1, "操作失败", nil)
		return
	}
	c.Json(0, "操作成功", nil)
}

func (c *Group) Delete() {
	id := c.GetInput("id")
	mgroup := models.NewGroupModel()
	affect, err := mgroup.Where("groupid = ?", id).Delete()
	if err != nil {
		c.Json(-1, "操作失败", nil)
		return
	}
	c.Json(0, "操作成功", affect)
}
