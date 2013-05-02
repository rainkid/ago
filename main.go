package main

import (
	"dogo"
	"fmt"
	admin "modules/admin"
	"os"
	"path"
)

func main() {
	router := getRouter()
	//get config
	basePath, _ := os.Getwd()
	configPath := path.Join(basePath, "src/configs", "app.yaml")

	// bootstrap and return a app
	app, err := dogo.Bootstrap(configPath)
	if err != nil {
		fmt.Println(err)
	}
	app.AddRouter(router)

	//app run
	app.Run()
}

func getRouter() *dogo.Router {
	// new dogo router
	var router = dogo.NewRouter()

	//AddRegexRoute
	//router.AddRegexRoute("/get/:uid", controllers.Home)

	//add map route
	router.AddMapRoute("admin", &admin.Errors{})
	router.AddMapRoute("admin", &admin.Login{})
	router.AddMapRoute("admin", &admin.Index{})
	return router
}
