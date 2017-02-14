package handler

import "net/http"

type HttpHandle struct {
	Method string
	Path string
	Function func(response http.ResponseWriter, request *http.Request)
}

