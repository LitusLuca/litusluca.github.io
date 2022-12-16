package renderer

import "github.com/litusluca/litusluca.github.io/src/glapi"

var sRenderer Renderer

type Renderer struct {
	gapi *glapi.GLapi
}

func Init(gl *glapi.GLapi) {
	sRenderer.gapi = gl
	SetClearColor(0.1,0.1,0.1,1.)
	gl.Enable(glapi.DEPTH_TEST)
}

func DrawIndexed(vao *VertexArray)  {
	vao.Bind()
	count := vao.GetIndexBuffer().GetCount()
	vao.GetIndexBuffer().Bind()
	sRenderer.gapi.DrawElements(glapi.TRIANGLES, count, glapi.UNSIGNED_INT, 0)
}

func Clear()  {
	sRenderer.gapi.Clear(glapi.COLOR_BUFFER_BIT)
}

func SetClearColor(r,g,b,a float32)  {
	sRenderer.gapi.SetClearColor(r,g,b,a)
}