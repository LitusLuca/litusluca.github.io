package events

type WindowResizeEvent struct {
	H             bool
	Width, Height int32
}

func (ev *WindowResizeEvent) Handled() bool {
	return ev.H
}

func (ev *WindowResizeEvent) SetHandled(handled bool) {
	if !ev.H {
		ev.H = handled
	}
}