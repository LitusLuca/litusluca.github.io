package events

import "github.com/litusluca/litusluca.github.io/src/input"

type MousePressEvent struct {
	H       bool
	KeyCode input.MouseButton
}

func (ev *MousePressEvent) Handled() bool {
	return ev.H
}

func (ev *MousePressEvent) SetHandled(handled bool) {
	if !ev.H {
		ev.H = handled
	}
}


type MouseReleaseEvent struct {
	H       bool
	KeyCode input.MouseButton
}

func (ev *MouseReleaseEvent) Handled() bool {
	return ev.H
}

func (ev *MouseReleaseEvent) SetHandled(handled bool) {
	if !ev.H {
		ev.H = handled
	}
}

type MouseScrollEvent struct {
	H       bool
	DX, DY float32
}

func (ev *MouseScrollEvent) Handled() bool {
	return ev.H
}

func (ev *MouseScrollEvent) SetHandled(handled bool) {
	if !ev.H {
		ev.H = handled
	}
}

func (ev *MouseScrollEvent) Type() EventType {
	return MouseScroll
}

type MouseMoveEvent struct {
	H       bool
	X, Y float32
	DX, DY float32
}

func (ev *MouseMoveEvent) Handled() bool {
	return ev.H
}

func (ev *MouseMoveEvent) SetHandled(handled bool) {
	if !ev.H {
		ev.H = handled
	}
}