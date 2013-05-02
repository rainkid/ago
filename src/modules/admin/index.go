package admin

import (
// "fmt"
// models "models"
)

type Index struct {
	AdminBase
}

func (c *Index) Init() {
	c.CheckLogin()
}

func (c *Index) Index() {

	// user := models.NewUserModel()

	//gets users	

	//delete users
	// user.Wheres(d).Delete()

	// d["created"] = "2012-12-12"
	/*id,_ := user.Sets(d).Add();
	fmt.Println(id)*/
	//
}
