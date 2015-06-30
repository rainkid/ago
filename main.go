package main

import (
	"flag"
	"github.com/rainkid/dogo"
	pserver "libs/pserver"
	websock "libs/websock"
	"os"
	"path"
)

var (
	host = flag.String("h", "127.0.0.1", "hostname for app runtime")
	port = flag.String("p", "8090", "port for app runtime")
)

func main() {
	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			dogo.Loger.E("run time panic: ", err)
		}
	}()

	pserver.Start()
	websock.Start()

	router := getRouter()

	// bootstrap and return a app
	app := dogo.NewApp(*host, *port)
	//Bootstrap and run

	app.Bootstrap(router).SetDefaultModule("admin").Run()
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
