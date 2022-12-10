//go:build wasm
// +build wasm

package glapi

import (
	"syscall/js"

	"github.com/litusluca/litusluca.github.io/src/utils/wasm"
)



type GLapi struct {
	glctx js.Value
	glObjectMap map[uint32]js.Value
	objectMapIndex uint32
}


func New(webglctx js.Value) (*GLapi, error) {
	gapi := new(GLapi)
	gapi.reset()
	gapi.glctx = webglctx
	gapi.glObjectMap = make(map[uint32]js.Value)
	gapi.objectMapIndex = 0

	return gapi, nil
}

func (gapi *GLapi) reset()  {
}

func (gapi *GLapi) SetClearColor(r, g, b, a float32)  {
	gapi.glctx.Call("clearColor", r, g, b, a)
}

func (gapi *GLapi) Clear(bits uint)  {
	gapi.glctx.Call("clear", int(bits))
}

func (gapi *GLapi) CreateBuffer() uint32 {
	gapi.glObjectMap[gapi.objectMapIndex] = gapi.glctx.Call("createBuffer")
	index := gapi.objectMapIndex
	gapi.objectMapIndex++
	return index
}

func (gapi *GLapi) BindBuffer(target uint32, id uint32, )  {
	gapi.glctx.Call("bindBuffer", target, gapi.glObjectMap[id])
}

func (gapi *GLapi) DeleteBuffer(id uint32)  {
	gapi.glctx.Call("deleteBuffer", gapi.glObjectMap[id])
	delete(gapi.glObjectMap, id)
}

func (gapi GLapi) BufferData(target uint32, data interface{}, usage uint32)  {
	gapi.glctx.Call("bufferData", target, wasm.SliceToTypedArray(data), usage)
}

func (gapi *GLapi) CreateVertexArray() uint32 {
	gapi.glObjectMap[gapi.objectMapIndex] = gapi.glctx.Call("createVertexArray")
	index := gapi.objectMapIndex
	gapi.objectMapIndex++
	return index
}

func (gapi 	*GLapi) BindVertexArray(id uint32) {
	gapi.glctx.Call("bindVertexArray", gapi.glObjectMap[id])
}

func (gapi *GLapi) DeleteVertexArray(id uint32)  {
	gapi.glctx.Call("deleteVertexArray", gapi.glObjectMap[id])
	delete(gapi.glObjectMap, id)
}

func (gapi *GLapi) EnableVertexAttribArray(index uint32)  {
	gapi.glctx.Call("enableVertexAttribArray", index)
}

func (gapi *GLapi) VertexAttribPointer(index, size, datatype, normalized, stride, offset uint32)  {
	gapi.glctx.Call("vertexAttribPointer", index, size, datatype, normalized, stride, offset)
}

func (gapi *GLapi) VertexAttribIPointer(index, size, datatype, stride, offset uint32)  {
	gapi.glctx.Call("vertexAttribIPointer", index, size, datatype, stride, offset)
}

func (gapi *GLapi) CreateProgram() uint32 {
	gapi.glObjectMap[gapi.objectMapIndex] = gapi.glctx.Call("createProgram")
	index := gapi.objectMapIndex
	gapi.objectMapIndex++
	return index
}

func (gapi *GLapi) DeleteProgram(id uint32)  {
	gapi.glctx.Call("deleteProgram", gapi.glObjectMap[id])
	delete(gapi.glObjectMap, id)
}

func (gapi *GLapi) UseProgram(id uint32)  {
	gapi.glctx.Call("useProgram", gapi.glObjectMap[id])
}

func (gapi *GLapi) AttachShader(program, shader uint32)  {
	gapi.glctx.Call("attachShader", gapi.glObjectMap[program], gapi.glObjectMap[shader])
}
func (gapi *GLapi) DetachShader(program, shader uint32)  {
	gapi.glctx.Call("detachShader", gapi.glObjectMap[program], gapi.glObjectMap[shader])
}

func (gapi *GLapi) LinkProgram(id uint32)  {
	gapi.glctx.Call("linkProgram", gapi.glObjectMap[id])
}

func (gapi *GLapi) GetProgramiv(id, pname uint32) int32 {
	res := gapi.glctx.Call("getProgramParameter", gapi.glObjectMap[id], pname)
	if pname == LINK_STATUS {
		switch res.Bool() {
		case true: return TRUE
		case false: return FALSE	
		}
	}
	return int32(res.Int())
}

func (gapi *GLapi) GetProgramInfoLog(id uint32) string {
	return gapi.glctx.Call("getProgramInfoLog", gapi.glObjectMap[id]).String()
}

func (gapi *GLapi) CreateShader(shaderType uint32) uint32 {
	gapi.glObjectMap[gapi.objectMapIndex] = gapi.glctx.Call("createShader", shaderType)
	index := gapi.objectMapIndex
	gapi.objectMapIndex++
	return index
}

func (gapi *GLapi) DeleteShader(id uint32)  {
	gapi.glctx.Call("deleteShader", gapi.glObjectMap[id])
	delete(gapi.glObjectMap, id)
}

func (gapi *GLapi) ShaderSource(id uint32, source string)  {
	gapi.glctx.Call("shaderSource", gapi.glObjectMap[id], source)
}

func (gapi *GLapi) CompileShader(id uint32)  {
	gapi.glctx.Call("compileShader", gapi.glObjectMap[id])
}

func (gapi *GLapi) GetShaderiv(id, pname uint32) int32 {
	res := gapi.glctx.Call("getShaderParameter", gapi.glObjectMap[id], pname)
	if pname == COMPILE_STATUS {
		switch res.Bool() {
		case true: return TRUE
		case false: return FALSE	
		}
	}
	return int32(res.Int())
}

func (gapi *GLapi) GetShaderInfoLog(id uint32) string {
	return gapi.glctx.Call("getShaderInfoLog", gapi.glObjectMap[id]).String()
}

func (gapi *GLapi) DrawElements(mode, count, indexType, offset uint32)  {
	gapi.glctx.Call("drawElements",mode, count, indexType, offset)
}