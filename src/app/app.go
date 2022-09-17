package app

import "time"

type IApp interface {
	OnCreate(application *Application)
	OnUpdate(dt time.Duration)
	OnDestroy()
}