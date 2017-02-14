package handler

import (
	"net/http"
	"github.com/geminikim/minimem/util"
	"github.com/geminikim/minimem/constant"
	"github.com/geminikim/minimem/manager"
)

type ListHttpHandler struct {
	manager manager.Manager
}

func NewListHttpHandler(manager manager.Manager) HttpHandler {
	handler := new(ListHttpHandler)
	handler.manager = manager
	return handler
}

func (handler ListHttpHandler) LeftPush(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMapWithBody(request, []string{constant.KEY}, constant.VALUE)
	result := handler.manager.Process(manager.Message{constant.LEFT_PUSH, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) LeftPop(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMap(request, []string{constant.KEY})
	result := handler.manager.Process(manager.Message{constant.LEFT_POP, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) LeftPeek(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMap(request, []string{constant.KEY})
	result := handler.manager.Process(manager.Message{constant.LEFT_PEEK, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPush(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMapWithBody(request, []string{constant.KEY}, constant.VALUE)
	result := handler.manager.Process(manager.Message{constant.RIGHT_PUSH, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPop(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMap(request, []string{constant.KEY})
	result := handler.manager.Process(manager.Message{constant.RIGHT_POP, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RightPeek(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMap(request, []string{constant.KEY})
	result := handler.manager.Process(manager.Message{constant.RIGHT_PEEK, value})
	response.Write([]byte(result))
}
func (handler ListHttpHandler) RangeGet(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMap(request, []string{constant.KEY,constant.INDEX,constant.COUNT})
	result := handler.manager.Process(manager.Message{constant.BY_RANGE, value})
	response.Write([]byte(result))
}

func (handler ListHttpHandler) GetHandles() []HttpHandle {
	return []HttpHandle {
		{http.MethodPost, constant.URL_LIST_LEFT_PUSH, handler.LeftPush},
		{http.MethodGet, constant.URL_LIST_LEFT_PEEK, handler.LeftPeek},
		{http.MethodGet, constant.URL_LIST_LEFT_POP, handler.LeftPop},
		{http.MethodPost, constant.URL_LIST_RIGHT_PUSH, handler.RightPush},
		{http.MethodGet, constant.URL_LIST_RIGHT_PEEK, handler.RightPeek},
		{http.MethodGet, constant.URL_LIST_RIGHT_POP, handler.RightPop},
		{http.MethodGet, constant.URL_LIST_BY_RANGE, handler.RangeGet},
	}
}