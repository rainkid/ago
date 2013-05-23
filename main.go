package main

import (
	"dogo"
	// "fmt"
	admin "modules/admin"
	"os"
	"path"
)

func main() {
	router := getRouter()
	// bootstrap and return a app

	basepath, _ := os.Getwd()
	file := path.Join(basepath, "src/configs", "app.yaml")
	app := dogo.NewApp(file)

	//register controller
	app.RegisterCtrl("admin", &admin.Errors{})
	app.RegisterCtrl("admin", &admin.Login{})
	app.RegisterCtrl("admin", &admin.Index{})

	//bootstart and run
	app.Bootstrap(router).SetDefaultModule("admin").Run()
}

func getRouter() *dogo.Router {
	// new dogo router
	var router = dogo.NewRouter()
	basepath, _ := os.Getwd()

	//add static router
	router.AddStaticRoute("/statics", path.Join(basepath, "src/statics/"))

	//add regex router and default is sample reoute
	router.AddRegexRoute("/login", "/admin/login/index")

	return router
}
