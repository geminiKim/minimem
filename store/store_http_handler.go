package store

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

type StoreHttpHandler struct {
	store *keyValueStore
}

func (handler StoreHttpHandler) Set(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var paramMap map[string]string
	err := decoder.Decode(&paramMap)

	if err != nil {
	panic(err)
	}
	handler.store.set(paramMap["key"], paramMap["value"])
}

func (handler StoreHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Write([]byte(handler.store.get(vars["key"])))
}

func NewStoreHttpHandler(store *keyValueStore) *StoreHttpHandler {
	handler := new(StoreHttpHandler)
	handler.store = store
	return handler
}