package list

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/util"
	"github.com/geminikim/minimem/constant"
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
	value[constant.KEY] = vars[constant.KEY]
	value[constant.VALUE] = util.ReadAll(request)

	result := handler.manager.Process(store.Message{constant.LEFT_PUSH, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) LeftPop(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]

	result := handler.manager.Process(store.Message{constant.LEFT_POP, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) LeftPeek(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]

	result := handler.manager.Process(store.Message{constant.LEFT_PEEK, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPush(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]
	value[constant.VALUE] = util.ReadAll(request)

	result := handler.manager.Process(store.Message{constant.RIGHT_PUSH, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPop(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]

	result := handler.manager.Process(store.Message{constant.RIGHT_POP, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPeek(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]

	result := handler.manager.Process(store.Message{constant.RIGHT_PEEK, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RangeGet(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]
	value[constant.INDEX] = vars[constant.INDEX]
	value[constant.COUNT] = vars[constant.COUNT]

	result := handler.manager.Process(store.Message{constant.BY_RANGE, value})
	response.Write([]byte(result))
}

func (handler ListHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{http.MethodPost, constant.URL_LIST_LEFT_PUSH, handler.LeftPush},
		{http.MethodGet, constant.URL_LIST_LEFT_PEEK, handler.LeftPeek},
		{http.MethodGet, constant.URL_LIST_LEFT_POP, handler.LeftPop},
		{http.MethodPost, constant.URL_LIST_RIGHT_PUSH, handler.RightPush},
		{http.MethodGet, constant.URL_LIST_RIGHT_PEEK, handler.RightPeek},
		{http.MethodGet, constant.URL_LIST_RIGHT_POP, handler.RightPop},
		{http.MethodGet, constant.URL_LIST_BY_RANGE, handler.RangeGet},
	}
}