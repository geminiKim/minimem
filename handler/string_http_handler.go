package handler

import (
	"net/http"
	"github.com/geminikim/minimem/constant"
	"github.com/geminikim/minimem/util"
	"github.com/geminikim/minimem/manager"
)

type StringHttpHandler struct {
	manager manager.Manager
}

func NewStringHttpHandler(manager manager.Manager) HttpHandler {
	handler := new(StringHttpHandler)
	handler.manager = manager
	return handler
}

func (handler StringHttpHandler) Set(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMapWithBody(request, []string{constant.KEY}, constant.VALUE)
	result := handler.manager.Process(manager.Message{constant.SET, value})
	response.Write([]byte(result))
}

func (handler StringHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMap(request, []string{constant.KEY})
	result := handler.manager.Process(manager.Message{constant.GET, value})
	response.Write([]byte(result))
}

func (handler StringHttpHandler) GetHandles() []HttpHandle {
	return []HttpHandle {
		{http.MethodPost, "/string/{key}", handler.Set},
		{http.MethodGet, "/string/{key}", handler.Get},
	}
}