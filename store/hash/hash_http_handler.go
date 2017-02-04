package hash

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"github.com/geminikim/minimem/store"
)

type HashHttpHandler struct {
	manager store.Manager
}

func NewHashHttpHandler(manager store.Manager) *HashHttpHandler {
	handler := new(HashHttpHandler)
	handler.manager = manager
	return handler
}

func (handler HashHttpHandler) Set(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	body, _ := ioutil.ReadAll(request.Body)

	value := make(map[string]string)
	value["key"] = vars["key"]
	value["field"] = vars["field"]
	value["value"] = string(body)

	result := handler.manager.Process(store.Message{"SET", value})
	response.Write([]byte(result))
}

func (handler HashHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]
	value["field"] = vars["field"]

	result := handler.manager.Process(store.Message{"GET", value})
	response.Write([]byte(result))
}

func (handler HashHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{"POST", "/hash/{key}/{field}", handler.Set},
		{"GET", "/hash/{key}/{field}", handler.Get},
	}
}