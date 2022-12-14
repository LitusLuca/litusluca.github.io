package main

import (
	"fmt"
	"time"

	"github.com/go-gl/mathgl/mgl32"
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
	-0.5, -0.5, -0.5, 0.5,0.,0.,
     0.5, -0.5, -0.5, 0.5,0.,0.,
     0.5,  0.5, -0.5, 0.5,0.,0.,
    -0.5,  0.5, -0.5, 0.5,0.,0.,

    -0.5, -0.5,  0.5, 0.5,0.,0.,
     0.5, -0.5,  0.5, 0.5,0.,0.,
     0.5,  0.5,  0.5, 0.5,0.,0.,
    -0.5,  0.5,  0.5, 0.5,0.,0.,

    -0.5,  0.5,  0.5, 0.,0.5,0.,
    -0.5,  0.5, -0.5, 0.,0.5,0.,
    -0.5, -0.5, -0.5, 0.,0.5,0.,
    -0.5, -0.5,  0.5, 0.,0.5,0.,

     0.5,  0.5,  0.5, 0.,0.5,0.,
     0.5,  0.5, -0.5, 0.,0.5,0.,
     0.5, -0.5, -0.5, 0.,0.5,0.,
     0.5, -0.5,  0.5, 0.,0.5,0.,

    -0.5, -0.5, -0.5, 0.,0.,0.5,
     0.5, -0.5, -0.5, 0.,0.,0.5,
     0.5, -0.5,  0.5, 0.,0.,0.5,
    -0.5, -0.5,  0.5, 0.,0.,0.5,

    -0.5,  0.5, -0.5, 0.,0.,0.5,
     0.5,  0.5, -0.5, 0.,0.,0.5,
     0.5,  0.5,  0.5, 0.,0.,0.5,
    -0.5,  0.5,  0.5, 0.,0.,0.5,
}
	vbo := renderer.NewVBO(positions)
	
	layout := renderer.NewBufferLayout([]renderer.BufferElement{renderer.MakeBufferElement(renderer.TypeFloat3, "Pos"),
	 renderer.MakeBufferElement(renderer.TypeFloat3, "Color")})

	vbo.SetLayout(layout)

	layer.triangle.AddVertexBuffer(vbo)

	indices := []uint32{
		0,1,2,
		2,3,0,

		4,5,6,
		6,7,4,

		8,9,10,
		10,11,8,

		12,13,14,
		14,15,12,

		16,17,18,
		18,19,16,

		20,21,22,
		22,23,20,
	}
	ibo := renderer.NewIBO(indices)
	layer.triangle.AddIndexBuffer(ibo)

	layer.shader = renderer.NewShaderFromFile("test.glsl")

	return layer
}

func (layer *BaseLayer) OnAttach() {
	fmt.Println("Attached!")
}

func (layer *BaseLayer) OnUpdate(dt time.Duration) {
	layer.runtime += dt.Seconds()
	//fmt.Println(layer.runtime)
	renderer.SetClearColor(0.3,0.1,0.3,1.)
	renderer.Clear()
	layer.shader.Bind()
	layer.shader.SetMat4("uViewProjection", mgl32.Perspective(90, 640./480., 0.1, 100))
	model := mgl32.HomogRotate3D(float32(layer.runtime), mgl32.Vec3{0.6,0.4,0.2}.Normalize())
	model = mgl32.Scale3D(1,1,1).Mul4(model)
	model = mgl32.Translate3D(0,0,-2).Mul4(model)
	layer.shader.SetMat4("uModel", model)
	renderer.DrawIndexed(layer.triangle)
}

func (layer *BaseLayer) OnDetach()  {
	fmt.Println("Detached!")
}