package events

import "github.com/litusluca/litusluca.github.io/src/input"

type KeyPressEvent struct {
	H uint8
	KeyCode input.KeyCode
}

func (ev *KeyPressEvent) Handled() uint8 {
	return ev.H
}

func (ev *KeyPressEvent) SetHandled(handled uint8) {
	ev.H |= handled
}

type KeyReleaseEvent struct {
	H uint8
	KeyCode input.KeyCode
}

func (ev *KeyReleaseEvent) Handled() uint8 {
	return ev.H
}

func (ev *KeyReleaseEvent) SetHandled(handled uint8) {
	ev.H |= handled
}
