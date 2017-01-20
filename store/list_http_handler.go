package store

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
)

type ListHttpHandler struct {
	store *listStore
}

func (handler ListHttpHandler) LeftPush(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var paramMap map[string]string
	err := decoder.Decode(&paramMap)

	if err != nil {
		panic(err)
	}
	handler.store.leftPush(paramMap["key"], paramMap["value"])
}
func (handler ListHttpHandler) LeftPop(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Write([]byte(handler.store.leftPop(vars["key"])))
}
func (handler ListHttpHandler) LeftPeek(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Write([]byte(handler.store.leftPeek(vars["key"])))
}
func (handler ListHttpHandler) RightPush(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var paramMap map[string]string
	err := decoder.Decode(&paramMap)

	if err != nil {
		panic(err)
	}
	handler.store.rightPush(paramMap["key"], paramMap["value"])
}
func (handler ListHttpHandler) RightPop(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Write([]byte(handler.store.rightPop(vars["key"])))
}
func (handler ListHttpHandler) RightPeek(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	response.Write([]byte(handler.store.rightPeek(vars["key"])))
}
func (handler ListHttpHandler) RangeGet(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	startIndex, _ := strconv.Atoi(vars["startIndex"])
	endIndex, _ := strconv.Atoi(vars["endIndex"])
	response.Write([]byte(handler.store.rangeGet(vars["key"], startIndex, endIndex)))
}

func NewListHttpHandler(store *listStore) *ListHttpHandler {
	handler := new(ListHttpHandler)
	handler.store = store
	return handler
}