package front

type Index struct {
	FrontBase
}

func (c *Index) Index() {
	c.Layout("layout/front/header.html")
	c.Layout("layout/front/footer.html")
}
