//go:build wasm
// +build wasm

package glapi

import (
	"syscall/js"
)

type GLapi struct {
	glctx js.Value
}


func New(webglctx js.Value) (*GLapi, error) {
	gapi := new(GLapi)
	gapi.reset()
	gapi.glctx = webglctx

	gapi.setDefaults()
	return gapi, nil
}

func (gapi *GLapi) reset()  {
}

func (gapi *GLapi) setDefaults()  {
	gapi.ClearColor(0, 0, 0, 1)
}

func (gapi *GLapi) ClearColor(r, g, b, a float32)  {
	gapi.glctx.Call("clearColor", r, g, b, a)
}

func (gapi *GLapi) Clear(bits uint)  {
	gapi.glctx.Call("clear", int(bits))
}