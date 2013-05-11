package admin

import (
	// "fmt"
	// models "models"
	// am_lib "lib/admin"
	configs "configs"
)

type Index struct {
	AdminBase
}

func (c *Index) Init() {
	c.CheckLogin()
	c.InitParams()
}

func (c *Index) Index() {
	menu := configs.NewMenu()
	c.Assign("JsonViews", menu.ToJson(menu.Views))
	c.Assign("JsonMenus", menu.ToJson(menu.Menus))
	c.Assign("Menus", menu.Menus)
}

func (c *Index) Default() {
	c.Layout("header", "layout/header.html")
	c.Layout("footer", "layout/footer.html")
}
