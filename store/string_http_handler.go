package store

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type StringHttpHandler struct {
	store *stringStore
}

func (handler StringHttpHandler) Set(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var paramMap map[string]string
	err := decoder.Decode(&paramMap)

	if err != nil {
	panic(err)
	}
	handler.store.set(paramMap["key"], paramMap["value"])
}

func (handler StringHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Write([]byte(handler.store.get(vars["key"])))
}

func NewStringHttpHandler(store *stringStore) *StringHttpHandler {
	handler := new(StringHttpHandler)
	handler.store = store
	return handler
}