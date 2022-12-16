//go:build wasm
// +build wasm

package window

import (
	"fmt"
	"syscall/js"

	"github.com/litusluca/litusluca.github.io/src/events"
	"github.com/litusluca/litusluca.github.io/src/glapi"
	"github.com/litusluca/litusluca.github.io/src/input"
)

type Window struct {
	canvas js.Value
	handle string
	width, height int32
	gl *glapi.GLapi
	eventCallback func(ev events.Event)
	fullScreen bool
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

	win.width = int32(win.canvas.Get("width").Int())
	win.height = int32(win.canvas.Get("height").Int())
	//create graphics api
	win.gl, _ = glapi.New(webglctx)

	win.canvas.Set("oncontextmenu", js.FuncOf(func(this js.Value, args []js.Value) interface{} {return false}))
	//TODO ev callbacks
	js.Global().Call("addEventListener", "keydown", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ev := args[0]

		evCode := ev.Get("code").String()
		KeyPress := events.KeyPressEvent{H: false, KeyCode: keyMap[evCode]}
		input.OnKeyEvent(keyMap[evCode],true)
		win.eventCallback(&KeyPress)
		return nil
	}))

	js.Global().Call("addEventListener", "keyup", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		ev := args[0]

		evCode := ev.Get("code").String()
		keyrelease := events.KeyReleaseEvent{H: false, KeyCode: keyMap[evCode]}
		input.OnKeyEvent(keyMap[evCode],false)
		win.eventCallback(&keyrelease)
		return nil
	}))

	
	win.canvas.Call("addEventListener", "mousemove", js.FuncOf(func(this js.Value, args []js.Value) any {
		ev := args[0]
		mousemove := events.MouseMoveEvent{H: false, X: float32(ev.Get("screenX").Float()), Y: float32(ev.Get("screenY").Float()),
											DX: float32(ev.Get("movementX").Float()), DY: float32(ev.Get("movementY").Float())}
		win.eventCallback(&mousemove)
		return nil
	}))

	win.canvas.Call("addEventListener", "fullscreenchange", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		if !win.fullScreen {
			win.SetSize(1920,1080)
			win.fullScreen = true
			win.gl.Viewport(0,0, win.width, win.height)
			//win.canvas.Call("requestPointerLock")
		}else{
			win.fullScreen = false
			win.SetSize(640,480)
			win.gl.Viewport(0,0, win.width, win.height)
		}
		fmt.Println(win.width, win.height)
		Resize := events.WindowResizeEvent{H: false, Width: win.width, Height: win.height}
		win.eventCallback(&Resize)
		
		return nil
	}))
	js.Global().Call("addEventListener", "pointerlockchange", js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("!!!")
		return nil
	}))

	return win, nil
}
func (win *Window) SetEventCallback(cb func(ev events.Event))  {
	win.eventCallback = cb
}

func (win *Window)OnUpdate()  {
	
}

func (win *Window)GetSize() (int32,int32) {
	return win.width, win.height
}
func (win *Window) SetSize(width, height int32)  {
	win.canvas.Set("width", width)
	win.canvas.Set("height", height)
	win.width = width
	win.height = height
}

func (win *Window) ToggleFullscreen()  {
	
	win.canvas.Call("requestPointerLock")
	win.canvas.Call("requestFullscreen")
	
}

func (win *Window) GLapi() *glapi.GLapi {
	return win.gl
}