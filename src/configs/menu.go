package configs

import (
	"encoding/json"
	"fmt"
)

type Menu struct {
	Menus []Item
	Views []Item
}

type Item struct {
	ID    string
	Name  string
	Url   string
	Items []Item
}

func NewMenu() *Menu {
	m := &Menu{}
	m.InitMenu()
	return m
}

func (m *Menu) Init() {
	m.Menus = []Item{
		Item{
			ID:   "admin_system",
			Name: "系统",
			Items: []Item{
				Item{
					ID:   "0",
					Name: "用户",
					Items: []Item{
						Item{
							ID:   "admin_user",
							Name: "用户管理",
							Url:  "/admin/user/index",
						},
						Item{
							ID:   "admin_group",
							Name: "用户组管理",
							Url:  "/admin/group/index",
						},
						Item{
							ID:   "admin_user_passwd",
							Name: "修改密码",
							Url:  "/admin/user/password",
						},
					},
				},
			},
		},
		Item{
			ID:   "admin_yuning",
			Name: "内容库",
			Items: []Item{
				Item{
					ID:   "1",
					Name: "用户",
					Items: []Item{
						Item{
							ID:   "admin_user1",
							Name: "用户管理1",
							Url:  "/admin/user/index",
						},
						Item{
							ID:   "admin_group2",
							Name: "用户组2",
							Url:  "/admin/group/index",
						},
						Item{
							ID:   "admin_user_passwd3",
							Name: "修改密码3",
							Url:  "/admin/user/password",
						},
					},
				},
			},
		},
	}
}

func (m *Menu) InitMenu() {
	m.Init()

	m.InitViews(m.Menus)
	// fmt.Println(m.ToJson(m.Views))
}

func (m *Menu) ToJson(Items []Item) string {
	b, err := json.Marshal(Items)
	if err != nil {
		fmt.Println("error:", err)
	}
	return string(b)
}

func (m *Menu) InitViews(Items []Item) {
	for _, v := range Items {
		if v.Items != nil {
			m.InitViews(v.Items)
		}
		m.Views = append(m.Views, Item{
			ID:   v.ID,
			Name: v.Name,
			Url:  v.Url,
		})
	}
}
