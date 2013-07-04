package main

import (
	"github.com/rainkid/dogo"
	"os"
	"path"
)

func main() {
	router := getRouter()

	// bootstrap and return a app
	basepath, _ := os.Getwd()
	file := path.Join(basepath, "src/configs", "app.ini")
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
	AddSampleRoute(router)

	//add regex router and default is sample route
	router.AddRegexRoute("/", "/admin/login/index")

	return router
}
