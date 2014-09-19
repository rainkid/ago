package main

import (
	"flag"
	"fmt"
	"github.com/rainkid/dogo"
	"os"
	"path"
)

func main() {
	cfgdir := flag.String("c", "", "please input build dir with")
	flag.Parse()
	l := len(*cfgdir)
	if l == 0 {
		fmt.Println("please input build dir with -c")
		os.Exit(0)
	}
	defer func() {
		fmt.Println("...")
		if err := recover(); err != nil {
			dogo.Loger.Println("run time panic: ", err)
		}
	}()

	router := getRouter()

	// bootstrap and return a app
	app := dogo.NewApp(*cfgdir + "/app.ini")

	//Bootstrap and run
	app.Bootstrap(router).SetDefaultModule("api").Run()
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
