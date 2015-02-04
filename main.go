package main

import (
	"flag"
	"fmt"
	"github.com/rainkid/dogo"
	pserver "libs/pserver"
	websock "libs/websock"
	"os"
	"path"
)

var cfgdir = flag.String("c", "", "please input build dir with")

func main() {
	flag.Parse()
	l := len(*cfgdir)
	if l == 0 {
		fmt.Println("please input build dir with -c")
		os.Exit(0)
	}

	defer func() {
		if err := recover(); err != nil {
			dogo.Loger.E("run time panic: ", err)
		}
	}()

	pserver.Start()
	websock.Start()

	router := getRouter()
	app_ini := fmt.Sprintf("%s/app.ini", *cfgdir)

	dogo.Register.Set("app_ini", app_ini)
	dogo.Register.Set("cfg_path", *cfgdir)

	// bootstrap and return a app
	app := dogo.NewApp(app_ini)
	//Bootstrap and run

	app.Bootstrap(router).SetDefaultModule("api").Run()
}

func getRouter() *dogo.Router {
	// new dogo router
	var router = dogo.NewRouter()
	basepath, _ := os.Getwd()

	//add static router
	router.AddStaticRoute("/statics", path.Join(basepath, "src/statics/"))
	//add regex router and default is sample route
	router.AddRegexRoute("/", "/admin/login/index")
	//add sample route
	AddSampleRoute(router)
	return router
}
