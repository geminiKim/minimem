package strings

import (
	"net/http"
	"github.com/geminikim/minimem/constant"
	"github.com/geminikim/minimem/util"
)

type StringHttpHandler struct {
	manager store.Manager
}

func NewStringHttpHandler(manager store.Manager) store.HttpHandler {
	handler := new(StringHttpHandler)
	handler.manager = manager
	return handler
}

func (handler StringHttpHandler) Set(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMapWithBody(request, []string{constant.KEY}, constant.VALUE)
	result := handler.manager.Process(store.Message{constant.SET, value})
	response.Write([]byte(result))
}

func (handler StringHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMap(request, []string{constant.KEY})
	result := handler.manager.Process(store.Message{constant.GET, value})
	response.Write([]byte(result))
}

func (handler StringHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{http.MethodPost, "/string/{key}", handler.Set},
		{http.MethodGet, "/string/{key}", handler.Get},
	}
}