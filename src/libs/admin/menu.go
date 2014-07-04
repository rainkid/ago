package admin

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
	m.Init()
	m.InitViews(m.Menus)
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
	}
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
