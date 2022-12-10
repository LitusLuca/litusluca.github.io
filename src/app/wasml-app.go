//go:build wasm
// +build wasm

package app

import (
	"syscall/js"
	"time"

	"github.com/litusluca/litusluca.github.io/src/layers"
	"github.com/litusluca/litusluca.github.io/src/renderer"
	"github.com/litusluca/litusluca.github.io/src/window"
)

var a *Application

type Application struct {
	userApp IApp
	window *window.Window
	running bool
	runtime time.Time
	lastFrameTime time.Time
	layerStack layers.LayerStack
}

func App(handle string, userApp IApp) *Application {
	if a != nil {
		return a
	}

	a = new(Application)
	var err error
	a.window, err = window.InitWindow(handle)
	if err != nil {
		panic(err)
	}
	renderer.Init(a.window.GLapi())
	a.layerStack = *layers.NewLayerStack()

	a.userApp = userApp
	userApp.OnCreate(a)
	return a
}
func (a *Application) Run(){

	done := make(chan bool)
	a.running = true
	a.lastFrameTime = time.Now()

	var tick js.Func
	tick = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		currentTime := time.Now()
		dt := currentTime.Sub(a.lastFrameTime)
		a.runtime.Add(dt)
		a.lastFrameTime = currentTime
		
		for i := 0; i < a.layerStack.GetLayerCount(); i++ {
			a.layerStack.GetLayerByIndex(i).OnUpdate(dt)
		}

		if a.running {
			js.Global().Call("requestAnimationFrame", tick)
		} else {
			done <- true
		}

		return nil
	})
	defer tick.Release()

	js.Global().Call("requestAnimationFrame", tick)

	<-done
	//destroywindow

	a.userApp.OnDestroy()
}

func (a *Application) Exit(){
	a.running = false
}

func (a *Application) GetWindow() *window.Window {
	return a.window
}

func (a *Application) GetLayerStack() *layers.LayerStack {
	return &a.layerStack
}