package events

type EventType int

const (
	Handled   = 1
	UnHandled = 0
)

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

type IEvent interface {
	Handled() uint8
	SetHandled(Handled uint8)
}

func NewDispatcher(ev IEvent) *Dispatcher {
	d := new(Dispatcher)
	d.EV = ev
	return d
}

type Dispatcher struct {
	EV IEvent
}

type Callback[T IEvent] func(T) uint8

//func (dp *Dispatcher) Dispatch(evType IEventType, cb Callback) {
//	if dp.EV.Type() == evType {
//		dp.EV.SetHandled(cb(dp.EV))
//	}
//}

func Dispatch[T IEvent](dispatcher *Dispatcher, cb Callback[T]) {
	if ev, ok := dispatcher.EV.(T); ok {
		ev.SetHandled(cb(ev))
	}
}