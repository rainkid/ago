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

	//bootstart and run
	app.Bootstrap(router).SetDefaultModule("admin").Run()
}

func getRouter() *dogo.Router {
	// new dogo router
	var router = dogo.NewRouter()
	basepath, _ := os.Getwd()

	//add static router
	router.AddStaticRoute("/statics", path.Join(basepath, "src/statics/"))

	//add sample route
	router.AddSampleRoute("admin", &admin.Errors{})
	router.AddSampleRoute("admin", &admin.Login{})
	router.AddSampleRoute("admin", &admin.Index{})

	//add regex router and default is sample route
	router.AddRegexRoute("/login", "/admin/login/index")

	return router
}
