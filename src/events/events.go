package events

type EventType int

const (
	WindowResize EventType = iota
	WindowFocus
	WindowClose

	KeyPress
	KeyRelease

	MouseMove
	MousePress
	MouseRelease
	MouseScroll
)

type Event interface {
	Handled() bool
	SetHandled(Handled bool)
	Type() EventType
}

type Dispatcher struct {
	EV Event
}

type Callback func(Event) bool

func (dp *Dispatcher) Dispatch(evType EventType, cb Callback) {
	if dp.EV.Type() == evType {
		dp.EV.SetHandled(cb(dp.EV))
	}
}