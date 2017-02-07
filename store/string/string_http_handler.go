package strings

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store"
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
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]
	value[constant.VALUE] = util.ReadAll(request)

	result := handler.manager.Process(store.Message{constant.SET, value})
	response.Write([]byte(result))
}

func (handler StringHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]

	result := handler.manager.Process(store.Message{constant.GET, value})
	response.Write([]byte(result))
}

func (handler StringHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{http.MethodPost, "/string/{key}", handler.Set},
		{http.MethodGet, "/string/{key}", handler.Get},
	}
}