package list

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/util"
)

type ListHttpHandler struct {
	manager store.Manager
}

func NewListHttpHandler(manager store.Manager) store.HttpHandler {
	handler := new(ListHttpHandler)
	handler.manager = manager
	return handler
}

func (handler ListHttpHandler) LeftPush(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]
	value["value"] = util.ReadAll(request)

	result := handler.manager.Process(store.Message{"LEFT_PUSH", value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) LeftPop(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]

	result := handler.manager.Process(store.Message{"LEFT_POP", value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) LeftPeek(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]

	result := handler.manager.Process(store.Message{"LEFT_PEEK", value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPush(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]
	value["value"] = util.ReadAll(request)

	result := handler.manager.Process(store.Message{"RIGHT_PUSH", value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPop(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]

	result := handler.manager.Process(store.Message{"RIGHT_POP", value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPeek(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]

	result := handler.manager.Process(store.Message{"RIGHT_PEEK", value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RangeGet(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value["key"] = vars["key"]
	value["index"] = vars["index"]
	value["count"] = vars["count"]

	result := handler.manager.Process(store.Message{"BY_RANGE", value})
	response.Write([]byte(result))
}

func (handler ListHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{"POST", "/list/{key}/leftPush", handler.LeftPush},
		{"GET", "/list/{key}/leftPeek", handler.LeftPeek},
		{"GET", "/list/{key}/leftPop", handler.LeftPop},
		{"POST", "/list/{key}/rightPush", handler.RightPush},
		{"GET", "/list/{key}/rightPeek", handler.RightPeek},
		{"GET", "/list/{key}/rightPop", handler.RightPop},
		{"GET", "/list/{key}/{index}/{count}", handler.RangeGet},
	}
}