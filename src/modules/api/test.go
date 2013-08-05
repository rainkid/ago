package api

type Test struct {
	ApiBase
}

func (c *Test) Index() {
	c.Json(0, "aaa", "Content")
}
