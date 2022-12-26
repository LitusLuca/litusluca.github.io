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
}

func NewDispatcher(ev Event) *Dispatcher {
	d := new(Dispatcher)
	d.EV = ev
	return d
}

type Dispatcher struct {
	EV Event
}

type Callback[T Event] func(T) bool

//func (dp *Dispatcher) Dispatch(evType EventType, cb Callback) {
//	if dp.EV.Type() == evType {
//		dp.EV.SetHandled(cb(dp.EV))
//	}
//}

func Dispatch[T Event](dispatcher *Dispatcher, cb Callback[T]) {
	if ev, ok := dispatcher.EV.(T); ok {
		ev.SetHandled(cb(ev))
	}
}