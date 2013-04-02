package controllers

import (
	"io"
	"dogo"
    models "models"
)

type Admin struct {
	dogo.Controller
}

func (c *Admin) Init() {
	//c.Layout("header", "layout/header.html")
	//c.Layout("footer", "layout/footer.html")
}


func (c *Admin) Index() {
	/*user := models.NewUserModel()

	d := make(map[string]interface{})
	d["departname"] = "研发部门"
	d["username"] = "astaxie"*/
	

	//gets users	
	// affect,_ := user.Filters(d).Gets()
	// fmt.Println(affect)

	//delete users
	// user.Filters(d).Delete()


	// d["created"] = "2012-12-12"
	/*id,_ := user.Sets(d).Add();
	fmt.Println(id)*/
	//
	
}

func (c *Admin) Login() {
	 c.DisableView()
}

func (c *Admin) Login_Post() {
	params := []string{"username", "password"}

	values := c.GetPost(params)
	
	user := models.NewUserModel()
	user.Sets(values)
	
	if code, msg := user.Valid(); code != 0{
		io.WriteString(c.GetReponse(), msg)
	}
	// models.User.New().Sets(values)
    user.Sets(values)
}
