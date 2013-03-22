package main

import (
	// "fmt"
	"dogo"
	controllers "controllers"
	// models "models"
)

func main() {
	// new dogo router
	var router = dogo.NewRouter()
	
	//AddRegexRoute
	regexRoute := dogo.NewRegexRoute("/get/:uid", controllers.Home)
	router.AddRegexRoute(regexRoute)

	//add map route 
	mapRoute := dogo.NewMapRoute(&controllers.Admin{})
	router.AddMapRoute(mapRoute)

	// bootstrap and return a app
	app := dogo.Bootstrap()
	app.AddRouter(router)

	//

	app.Run()
}
