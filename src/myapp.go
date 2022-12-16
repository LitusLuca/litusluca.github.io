package main

import (
	"fmt"

	"github.com/litusluca/litusluca.github.io/src/app"
)

type MyApp struct {
	parentApp *app.Application
	hi string
}

func (a *MyApp) OnCreate(application *app.Application) {
	a.hi = "Helloworld"
	a.parentApp = application
	application.GetLayerStack().PushLayer(CreateBaseLayer())
}

func (a *MyApp) OnDestroy() {
	fmt.Printf("GoodBye")
}