package list

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
	"io/ioutil"
)

type ListHttpHandler struct {
	store *listStore
}

func (handler ListHttpHandler) LeftPush(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	body, _ := ioutil.ReadAll(request.Body)
	handler.store.leftPush(vars["key"], string(body))
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
	vars := mux.Vars(request)
	body, _ := ioutil.ReadAll(request.Body)
	handler.store.rightPush(vars["key"], string(body))
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
	index, _ := strconv.Atoi(vars["index"])
	count, _ := strconv.Atoi(vars["count"])

	list := handler.store.rangeGet(vars["key"], index, count)
	result, _ := json.Marshal(list)
	response.Write(result)
}

func NewListHttpHandler(store *listStore) *ListHttpHandler {
	handler := new(ListHttpHandler)
	handler.store = store
	return handler
}