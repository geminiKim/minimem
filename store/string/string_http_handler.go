package strings

import (
	"net/http"
	"github.com/gorilla/mux"
	"io/ioutil"
	"github.com/geminikim/minimem/store"
)

type StringHttpHandler struct {
	store *stringStore
}

func NewStringHttpHandler(store *stringStore) *StringHttpHandler {
	handler := new(StringHttpHandler)
	handler.store = store
	return handler
}

func (handler StringHttpHandler) Set(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	body, _ := ioutil.ReadAll(request.Body)
	handler.store.set(vars["key"], string(body))
}

func (handler StringHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Write([]byte(handler.store.get(vars["key"])))
}

func (handler StringHttpHandler) GetHandles() []store.Handle {
	return []store.Handle {
		{"POST", "/string/{key}", handler.Set},
		{"GET", "/string/{key}", handler.Get},
	}
}