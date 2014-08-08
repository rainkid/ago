package models

type Group struct {
	Model
	code int64
}

func NewGroupModel() *Group {
	return &Group{
		Model: Model{TableName: "admin_group", PrimaryKey: "groupid"},
		code:  1200,
	}
}

func (g *Group) Valid(mData *map[string]string) (int64, string) {
	d := *mData
	if _, ok := d["name"]; !ok {
		return g.code + 1, "名称不能为空."
	}
	if _, ok := d["descrip"]; !ok {
		return g.code + 2, "描述不能为空."
	}
	if _, ok := d["rvalue"]; !ok {
		return g.code + 3, "请选择至少一个权限."
	}
	return 0, ""
}
