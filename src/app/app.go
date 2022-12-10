package app

type IApp interface {
	OnCreate(application *Application)
	OnDestroy()
}