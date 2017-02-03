package store


type HttpHandler interface {
	GetHandles() []HttpHandle
}