package admin

import (
	"encoding/json"
	admin "libs/admin"
)

type Index struct {
	AdminBase
}

func (c *Index) Index() {
	menu := admin.NewMenu()
	views, err := json.Marshal(menu.Views)
	if err == nil {
		c.Assign("JsonViews", string(views))
	}
	menus, err := json.Marshal(menu.Menus)
	if err == nil {
		c.Assign("JsonMenus", string(menus))
		c.Assign("Menus", menu.Menus)
	}
}

func (c *Index) Default() {

}
