package handler


type HttpHandler interface {
	GetHandles() []HttpHandle
}