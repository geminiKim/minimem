package hash

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"github.com/geminikim/minimem/handler/http"
)

type HashHttpHandler struct {
	store *hashStore
}

func NewHashHttpHandler(store *hashStore) *HashHttpHandler {
	handler := new(HashHttpHandler)
	handler.store = store
	return handler
}

func (handler HashHttpHandler) Set(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	body, _ := ioutil.ReadAll(request.Body)
	handler.store.set(vars["key"], vars["field"], string(body))
}

func (handler HashHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Write([]byte(handler.store.get(vars["key"], vars["field"])))
}

func (handler HashHttpHandler) GetHandles() []handler.HttpHandle {
	return []handler.HttpHandle {
		{"POST", "/hash/{key}/{field}", handler.Set},
		{"GET", "/hash/{key}/{field}", handler.Get},
	}
}