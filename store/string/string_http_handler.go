package strings

import (
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"github.com/geminikim/minimem/store"
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
	body, _ := ioutil.ReadAll(request.Body)

	value := make(map[string]string)
	value["key"] = vars["key"]
	value["value"] = string(body)

	result := handler.manager.Process(store.Message{"SET", value})
	response.Write([]byte(result))
}

func (handler StringHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]

	result := handler.manager.Process(store.Message{"GET", value})
	response.Write([]byte(result))
}

func (handler StringHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{"POST", "/string/{key}", handler.Set},
		{"GET", "/string/{key}", handler.Get},
	}
}