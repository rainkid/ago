package admin

import (
	configs "configs"
	"encoding/json"
	"fmt"
)

type MenuItem map[string]string
type MenuMap map[string]interface{}

type Menu struct {
	menus []MenuItem
}

func NewMenu() *Menu {
	return &Menu{}
}

func (m *Menu) GetJson() {
	var vo interface{}
	views := configs.GetMenu()
	json.Unmarshal(views, &vo)
	for _, v := range vo.(map[string]interface{}) {
		m.Match(v)
	}

	b, err := json.Marshal(m.menus)
	if err != nil {
		fmt.Println(err, b)
	}
}

func (m *Menu) Match(pitem interface{}) {
	switch el := pitem.(type) {
	case map[string]interface{}:

		m.AddItem(el)

		if el["items"] != nil {
			m.Match(el["items"].([]interface{}))
		}
	case []interface{}:
		for _, v := range el {

			child := v.(map[string]interface{})
			m.AddItem(child)

			if child["items"] != nil {
				m.Match(child["items"].([]interface{}))
			}
		}
	default:
	}
}

func (m *Menu) AddItem(child MenuMap) {
	item := make(MenuItem)
	if child["url"] != nil {
		item["url"] = child["url"].(string)
	}
	item["name"] = child["name"].(string)
	item["id"] = child["id"].(string)
	m.menus = append(m.menus, item)
}
