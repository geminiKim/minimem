package store

import "net/http"

type Handle struct {
	Method string
	Path string
	Function func(response http.ResponseWriter, request *http.Request)
}

