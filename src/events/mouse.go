package events

import "github.com/litusluca/litusluca.github.io/src/input"

type MousePressEvent struct {
	H       uint8
	KeyCode input.MouseButton
}

func (ev *MousePressEvent) Handled() uint8 {
	return ev.H
}

func (ev *MousePressEvent) SetHandled(handled uint8) {
	ev.H |= handled
}


type MouseReleaseEvent struct {
	H       uint8
	KeyCode input.MouseButton
}

func (ev *MouseReleaseEvent) Handled() uint8 {
	return ev.H
}

func (ev *MouseReleaseEvent) SetHandled(handled uint8) {
	ev.H |= handled
}

type MouseScrollEvent struct {
	H       uint8
	DX, DY float32
}

func (ev *MouseScrollEvent) Handled() uint8 {
	return ev.H
}

func (ev *MouseScrollEvent) SetHandled(handled uint8) {
	ev.H |= handled
}

func (ev *MouseScrollEvent) Type() EventType {
	return MouseScroll
}

type MouseMoveEvent struct {
	H       uint8
	X, Y float32
	DX, DY float32
}

func (ev *MouseMoveEvent) Handled() uint8 {
	return ev.H
}

func (ev *MouseMoveEvent) SetHandled(handled uint8) {
	ev.H |= handled
}