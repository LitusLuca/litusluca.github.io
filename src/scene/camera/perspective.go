package camera

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/litusluca/litusluca.github.io/src/events"
)

type PerspectiveCamera struct {
	projection mgl32.Mat4
	view mgl32.Mat4
	viewProjection mgl32.Mat4
	fov, aspect, near, far float32
}

func NewPerspectiveCamera(fov, aspect, near, far float32) *PerspectiveCamera {
	cam := new(PerspectiveCamera)
	fov = fov * math.Pi / 180
	cam.fov, cam.aspect, cam.near, cam.far = fov, aspect, near, far
	cam.projection = mgl32.Perspective(fov, aspect, near, far)
	cam.view = mgl32.Ident4()
	cam.viewProjection = cam.projection.Mul4(cam.view)
	return cam
}

func (cam *PerspectiveCamera) LookAt(eye, center, up mgl32.Vec3)  {
	cam.view = mgl32.LookAtV(eye,center,up)
	cam.updateMatrices()
}

func (cam *PerspectiveCamera) updateMatrices()  {
	cam.viewProjection = cam.projection.Mul4(cam.view)
}

func (cam *PerspectiveCamera) ViewProjection() mgl32.Mat4 {
	return cam.viewProjection
}

func (cam *PerspectiveCamera) OnResize(ev *events.WindowResizeEvent) uint8 {
	aspect := float32(ev.Width)/float32(ev.Height)
	cam.projection = mgl32.Perspective(cam.fov, aspect, cam.near, cam.far)
	return events.UnHandled
}