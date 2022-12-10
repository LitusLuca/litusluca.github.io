package main

import (
	"fmt"
	"time"

	"github.com/litusluca/litusluca.github.io/src/renderer"
)

type BaseLayer struct {
	runtime float64
	triangle *renderer.VertexArray
	shader *renderer.ShaderProgram
}

func CreateBaseLayer() *BaseLayer {
	layer := new(BaseLayer)
	layer.runtime = 0

	layer.triangle = renderer.NewVAO()

	positions := []float32{
		0., 0.5, 0.,
		-0.5, -0.5, 0.,
		0.5, -0.5, 0.,
	}
	vbo := renderer.NewVBO(positions)
	
	layout := renderer.NewBufferLayout([]renderer.BufferElement{renderer.MakeBufferElement(renderer.TypeFloat3, "Pos")})

	vbo.SetLayout(layout)

	layer.triangle.AddVertexBuffer(vbo)

	indices := []uint32{0,1,2}
	ibo := renderer.NewIBO(indices)
	layer.triangle.AddIndexBuffer(ibo)


	vertexShader := `#version 300 es
	
	layout(location=0) in vec3 a_Pos;

	void main() {
		gl_Position = vec4(a_Pos, 1.0);
	}
	`
	fragmentShader := `#version 300 es

	precision mediump float;
	
	out vec4 FragmentColor;

	void main() {
		FragmentColor = vec4(1.0);
	}
	`

	layer.shader = renderer.NewShader("basic", vertexShader, fragmentShader)

	return layer
}

func (layer *BaseLayer) OnAttach() {
	fmt.Println("Attached!")
}

func (layer *BaseLayer) OnUpdate(dt time.Duration) {
	layer.runtime += dt.Seconds()
	fmt.Println(layer.runtime)
	renderer.SetClearColor(0.3,0.1,0.3,1.)
	renderer.Clear()
	layer.shader.Bind()
	renderer.DrawIndexed(layer.triangle)
}

func (layer *BaseLayer) OnDetach()  {
	fmt.Println("Detached!")
}