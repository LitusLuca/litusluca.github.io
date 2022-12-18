package controller

import (
	"math"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/litusluca/litusluca.github.io/src/events"
	"github.com/litusluca/litusluca.github.io/src/input"
	"github.com/litusluca/litusluca.github.io/src/scene/camera"
)

type FPSController struct {
	Camera *camera.PerspectiveCamera
	position mgl32.Vec3
	yaw, pitch float32
	front, right, worldUp, up mgl32.Vec3
	mouseX, mouseY float32
	Speed, Sensivity float32
}

func NewFPSController(cam *camera.PerspectiveCamera, position mgl32.Vec3, startYaw, startPitch, startMouseX, startMouseY float32) *FPSController{
	controller := new(FPSController)
	controller.Camera = cam
	controller.position = position
	controller.yaw = startYaw
	controller.pitch = startPitch
	controller.mouseX = startMouseX
	controller.mouseY = startMouseY
	controller.Speed = 1
	controller.Sensivity = 0.5

	controller.worldUp = mgl32.Vec3{0,1,0}
	controller.updateFront()
	return controller
}

func (controller *FPSController) updateFront()  {
	controller.front = mgl32.Vec3{
		float32(math.Cos(float64(controller.pitch)*math.Pi/180)*math.Sin(float64(controller.yaw)*math.Pi/180)),
		float32(math.Sin(float64(controller.pitch)*math.Pi/180)),
		float32(math.Cos(float64(controller.pitch)*math.Pi/180)*math.Cos(float64(controller.yaw)*math.Pi/180)),
	}
	controller.right = controller.front.Cross(controller.worldUp).Normalize()
	controller.up = controller.right.Cross(controller.front)
	controller.Camera.LookAt(controller.position, controller.position.Add(controller.front), controller.up)
}

func (controller *FPSController) OnUpdate(dt float32)  {
	controller.updateFront()
	move := mgl32.Vec3{0}
	if input.IsKeyPressed(input.KEY_W) {
		move = move.Add(controller.front)
	}
	if input.IsKeyPressed(input.KEY_S) {
		move = move.Add(controller.front.Mul(-1))
	}
	if input.IsKeyPressed(input.KEY_A) {
		move = move.Add(controller.right.Mul(-1))
	}
	if input.IsKeyPressed(input.KEY_D) {
		move = move.Add(controller.right)
	}
	l := move.Len()
	if  l != 0{
		move = move.Mul(1/l)
	}
	controller.position = controller.position.Add(move.Mul(controller.Speed*dt))
	controller.Camera.LookAt(controller.position, controller.position.Add(controller.front), controller.up)
}

func (controller *FPSController) OnMouseMove(ev events.Event) bool {
	moveEv := ev.(*events.MouseMoveEvent)
	dx := moveEv.DX
	dy := -moveEv.DY
	controller.mouseX = moveEv.X
	controller.mouseY = moveEv.Y
	controller.yaw -= dx * controller.Sensivity
	controller.pitch += dy * controller.Sensivity

	if controller.pitch > 89 {
		controller.pitch = 89
	}else if controller.pitch < -89{
		controller.pitch = -89
	}
	return false
}