package main

import (
	"github.com/rainkid/dogo"
	admin "modules/admin"
	debug "modules/debug"
	front "modules/front"
)

func AddSampleRoute(router *dogo.Router) {
	router.AddSampleRoute("admin", &admin.Errors{})
	router.AddSampleRoute("admin", &admin.Group{})
	router.AddSampleRoute("admin", &admin.Index{})
	router.AddSampleRoute("admin", &admin.Login{})
	router.AddSampleRoute("admin", &admin.User{})
	router.AddSampleRoute("debug", &debug.Pprof{})
	router.AddSampleRoute("front", &front.Index{})
}