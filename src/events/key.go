package events

import "github.com/litusluca/litusluca.github.io/src/input"

type KeyPressEvent struct {
	H bool
	KeyCode input.KeyCode
}

func (ev *KeyPressEvent) Handled() bool {
	return ev.H
}

func (ev *KeyPressEvent) SetHandled(handled bool) {
	if !ev.H {
		ev.H = handled
	}
}

func (ev *KeyPressEvent) Type() EventType {
	return KeyPress
}

type KeyReleaseEvent struct {
	H bool
	KeyCode input.KeyCode
}

func (ev *KeyReleaseEvent) Handled() bool {
	return ev.H
}

func (ev *KeyReleaseEvent) SetHandled(handled bool) {
	if !ev.H {
		ev.H = handled
	}
}

func (ev *KeyReleaseEvent) Type() EventType {
	return KeyRelease
}