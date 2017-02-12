package hash

import (
	"net/http"
	"github.com/geminikim/minimem/constant"
	"github.com/geminikim/minimem/util"
)

type HashHttpHandler struct {
	manager store.Manager
}

func NewHashHttpHandler(manager store.Manager) *HashHttpHandler {
	handler := new(HashHttpHandler)
	handler.manager = manager
	return handler
}

func (handler HashHttpHandler) Set(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMapWithBody(request, []string{constant.KEY,constant.FIELD}, constant.VALUE)
	result := handler.manager.Process(store.Message{constant.SET, value})
	response.Write([]byte(result))
}

func (handler HashHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	value := util.GetMessageMap(request, []string{constant.KEY,constant.FIELD})
	result := handler.manager.Process(store.Message{constant.GET, value})
	response.Write([]byte(result))
}

func (handler HashHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{http.MethodPost, constant.URL_HASH_SET, handler.Set},
		{http.MethodGet, constant.URL_HASH_GET, handler.Get},
	}
}