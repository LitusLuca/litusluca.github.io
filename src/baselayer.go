package main

import (
	"fmt"
	"time"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/litusluca/litusluca.github.io/src/app"
	"github.com/litusluca/litusluca.github.io/src/events"
	"github.com/litusluca/litusluca.github.io/src/input"
	"github.com/litusluca/litusluca.github.io/src/renderer"
	"github.com/litusluca/litusluca.github.io/src/scene/camera"
	"github.com/litusluca/litusluca.github.io/src/scene/controller"
)

type BaseLayer struct {
	runtime float64
	triangle *renderer.VertexArray
	plane *renderer.VertexArray
	shader *renderer.ShaderProgram
	aspect float32
	camController *controller.FPSController
	tex *renderer.Texture2D
	ground *renderer.Texture2D
}

func CreateBaseLayer() *BaseLayer {
	layer := new(BaseLayer)
	layer.runtime = 0

	layer.triangle = renderer.NewVAO()

	positions := []float32{
	-0.5, -0.5, -0.5, 1.,1.,
     0.5, -0.5, -0.5, 0., 1.,
     0.5,  0.5, -0.5, 0., 0.,
    -0.5,  0.5, -0.5, 1., 0.,

    -0.5, -0.5,  0.5, 1.,1.,
     0.5, -0.5,  0.5, 0., 1.,
     0.5,  0.5,  0.5, 0., 0.,
    -0.5,  0.5,  0.5, 1., 0.,

    -0.5,  0.5,  0.5, 1.,1.,
    -0.5,  0.5, -0.5, 0., 1.,
    -0.5, -0.5, -0.5, 0., 0.,
    -0.5, -0.5,  0.5, 1., 0.,

     0.5,  0.5,  0.5, 1.,1.,
     0.5,  0.5, -0.5, 0., 1.,
     0.5, -0.5, -0.5, 0., 0.,
     0.5, -0.5,  0.5, 1., 0.,

    -0.5, -0.5, -0.5, 1.,1.,
     0.5, -0.5, -0.5, 0., 1.,
     0.5, -0.5,  0.5, 0., 0.,
    -0.5, -0.5,  0.5, 1., 0.,

    -0.5,  0.5, -0.5, 1.,1.,
     0.5,  0.5, -0.5, 0., 1.,
     0.5,  0.5,  0.5, 0., 0.,
    -0.5,  0.5,  0.5, 1., 0.,
}
	vbo := renderer.NewVBO(positions)
	
	layout := renderer.NewBufferLayout([]renderer.BufferElement{
		renderer.MakeBufferElement(renderer.TypeFloat3, "Pos"), 
		renderer.MakeBufferElement(renderer.TypeFloat2, "TexCoord")})

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

	planeVecs := []float32{
		-1.,  0., -1.,  10.,10.,
		-1.,  0.,  1.,  0., 10.,
		 1.,  0.,  1.,  0., 0.,
		 1.,  0., -1.,  10., 0.,
	}
	planeIndices := []uint32{
		0,1,2,
		2,3,0,
	}
	layer.plane = renderer.NewVAO()
	planeVbo := renderer.NewVBO(planeVecs)
	planeIbo := renderer.NewIBO(planeIndices)
	planeVbo.SetLayout(layout)
	layer.plane.AddVertexBuffer(planeVbo)
	layer.plane.AddIndexBuffer(planeIbo)


	layer.shader = renderer.NewShaderFromFile("basicColor.glsl")
	layer.aspect = 640./480.


	cam := camera.NewPerspectiveCamera(90, 16/9, 0.1, 100)
	layer.camController = controller.NewFPSController(cam, mgl32.Vec3{0,0,0}, 0, 0, 0, 0)

	layer.tex = renderer.NewTexture2D("container.jpeg", renderer.TrueColor)
	layer.ground = renderer.NewTexture2D("rock.tif", renderer.TrueColor)

	return layer
}

func (layer *BaseLayer) OnAttach() {
	fmt.Println("Attached!")
}

func (layer *BaseLayer) OnUpdate(dt time.Duration) {
	layer.runtime += dt.Seconds()
	renderer.SetClearColor(0.3,0.1,0.3,1.)
	renderer.Clear()
	layer.shader.Bind()
	layer.camController.OnUpdate(float32(dt.Seconds()))
	layer.shader.SetMat4("uViewProjection", layer.camController.Camera.ViewProjection())
	model := mgl32.HomogRotate3D(float32(layer.runtime), mgl32.Vec3{0.6,0.4,0.2}.Normalize())
	model = mgl32.Scale3D(1,1,1).Mul4(model)
	model = mgl32.Translate3D(0,0,2).Mul4(model)
	layer.shader.SetMat4("uModel", model)
	layer.tex.Bind(0)
	renderer.DrawIndexed(layer.triangle)
	layer.ground.Bind(0)
	layer.shader.SetMat4("uModel",mgl32.Translate3D(0,-1,0).Mul4(mgl32.Scale3D(10,10,10)))
	renderer.DrawIndexed(layer.plane)
}

func (layer *BaseLayer) OnDetach()  {
	fmt.Println("Detached!")
}

func (layer *BaseLayer) OnEvent(ev events.Event) {
	dispatcher := events.Dispatcher{EV:ev}
	dispatcher.Dispatch(events.KeyPress, layer.OnKeyPress)
	dispatcher.Dispatch(events.WindowResize, layer.camController.Camera.OnResize)
	dispatcher.Dispatch(events.MouseMove, layer.camController.OnMouseMove)
}

func (layer *BaseLayer) OnKeyPress(ev events.Event) bool {
	KpEV := ev.(*events.KeyPressEvent)
	switch KpEV.KeyCode {
	case input.KEY_F:
		app.GetApp().GetWindow().ToggleFullscreen()
	}
	return false
}
