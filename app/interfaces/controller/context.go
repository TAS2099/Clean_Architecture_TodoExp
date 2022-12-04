package controller

type Context interface {
	Bind(interface{}) error
}
