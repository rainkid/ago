package admin

import (
	admin "libs/admin"
)

type Index struct {
	AdminBase
}

func (c *Index) Init() {
	c.CheckLogin()
	c.InitParams()
}

func (c *Index) Index() {
	menu := admin.NewMenu()
	c.Assign("JsonViews", menu.ToJson(menu.Views))
	c.Assign("JsonMenus", menu.ToJson(menu.Menus))
	c.Assign("Menus", menu.Menus)
}

func (c *Index) Default() {
	c.Layout("header", "layout/header.html")
	c.Layout("footer", "layout/footer.html")
}
