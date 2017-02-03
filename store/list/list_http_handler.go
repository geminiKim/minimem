package list

import (
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"strconv"
	"io/ioutil"
	"github.com/geminikim/minimem/store"
)

type ListHttpHandler struct {
	store *listStore
}

func NewListHttpHandler(store *listStore) *ListHttpHandler {
	handler := new(ListHttpHandler)
	handler.store = store
	return handler
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

func (handler ListHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{"POST", "/list/{key}/leftPush", handler.LeftPush},
		{"GET", "/list/{key}/leftPeek", handler.LeftPeek},
		{"GET", "/list/{key}/leftPop", handler.LeftPop},
		{"POST", "/list/{key}/rightPush", handler.RightPush},
		{"GET", "/list/{key}/rightPeek", handler.RightPeek},
		{"GET", "/list/{key}/rightPop", handler.RightPop},
		{"POST", "/list/{key}/{index}/{count}", handler.RangeGet},
	}
}