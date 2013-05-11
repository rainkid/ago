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
	c.Assigns()
}

func (c *Index) Index() {
	menu := configs.NewMenu()
	c.Set("JsonViews", menu.ToJson(menu.Views))
	c.Set("JsonMenus", menu.ToJson(menu.Menus))
	c.Set("Menus", menu.Menus)
}

func (c *Index) Default() {
	c.Layout("header", "layout/header.html")
	c.Layout("footer", "layout/footer.html")
}
