package main

import (
	"fmt"
	"time"

	"github.com/litusluca/litusluca.github.io/src/app"
)

type MyApp struct {
	parentApp *app.Application
	hi string
}

func (a *MyApp) OnCreate(application *app.Application) {
	a.hi = "Helloworld"
	a.parentApp = application
}

func (a *MyApp) OnUpdate(dt time.Duration) {
	fmt.Printf("%v  fps: %v\n", a.hi, 1. / dt.Seconds())
}

func (a *MyApp) OnDestroy() {
	fmt.Printf("GoodBye")
}