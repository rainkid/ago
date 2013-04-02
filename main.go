package main

import (
	// "fmt"
	"os"
	"path"
	"dogo"
	controllers "controllers"
	// models "models"
)

func main() {
	router := getRouter()
	//get config
	basePath, _ := os.Getwd()
	configPath := path.Join(basePath, "src/configs", "app.yaml")

	// bootstrap and return a app
	app := dogo.Bootstrap(configPath, "develop")
	app.AddRouter(router)

	//app run
	app.Run()
}

func getRouter() *dogo.Router{
	// new dogo router
	var router = dogo.NewRouter()
	
	//AddRegexRoute
	regexRoute := dogo.NewRegexRoute("/get/:uid", controllers.Home)
	router.AddRegexRoute(regexRoute)

	//add map route 
	mapRoute := dogo.NewMapRoute(&controllers.Admin{})
	router.AddMapRoute(mapRoute)
	return router
}
