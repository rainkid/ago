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

func (g *Group) Valid() (int64, string) {
	name, nlen := g.GetData("name")
	descrip, dlen := g.GetData("descrip")
	rvalue, rlen := g.GetData("rvalue")

	if nlen == 0 {
		return g.code + 1, "名称不能为空."
	}

	if dlen == 0 {
		return g.code + 2, "描述不能为空."
	}

	if rlen == 0 {
		return g.code + 3, "请选择至少一个权限."
	}

	g.Data["name"] = name
	g.Data["descrip"] = descrip
	g.Data["rvalue"] = rvalue

	return 0, ""
}
