package events

type WindowResizeEvent struct {
	H             uint8
	Width, Height int32
}

func (ev *WindowResizeEvent) Handled() uint8 {
	return ev.H
}

func (ev *WindowResizeEvent) SetHandled(handled uint8) {
	ev.H |= handled
}