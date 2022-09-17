//go:build wasm
// +build wasm

package window

import (
	"fmt"
	"syscall/js"

	"github.com/litusluca/litusluca.github.io/src/glapi"
)

type Window struct {
	canvas js.Value
	handle string
	width, height int
	gl *glapi.GLapi
}

func InitWindow(handle string) (*Window, error){
	if sWindow != nil {
		return nil, fmt.Errorf("!!! window: window already exists")
	}
	win := new(Window)
	sWindow = win

	doc := js.Global().Get("document")
	win.canvas = doc.Call("getElementById", handle)
	win.handle = handle
	if win.canvas.Equal(js.Null()){
		return nil, fmt.Errorf("!!! window: cannot find canvas with handle: %v", handle)
	}

	webglctx := win.canvas.Call("getContext", "webgl2")
	if webglctx.Equal(js.Null()) {
		return nil, fmt.Errorf("!!! window: browser does not support webgl2")
	}

	win.width = win.canvas.Get("width").Int()
	win.height = win.canvas.Get("height").Int()
	//create graphics api
	win.gl, _ = glapi.New(webglctx)

	win.canvas.Set("oncontextmenu", js.FuncOf(func(this js.Value, args []js.Value) interface{} {return false}))
	//TODO ev callbacks
	return win, nil
}

func (win *Window)OnUpdate()  {
	
}

func (win *Window)GetSize() (int,int) {
	return win.width, win.height
}
func (win *Window) SetSize(width, height int)  {
	win.canvas.Set("width", width)
	win.canvas.Set("height", height)
	win.width = width
	win.height = height
}

func (win *Window) GLapi() *glapi.GLapi {
	return win.gl
}