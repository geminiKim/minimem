package hash

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/constant"
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
	vars := mux.Vars(request)
	body, _ := ioutil.ReadAll(request.Body)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]
	value[constant.FIELD] = vars[constant.FIELD]
	value[constant.VALUE] = string(body)

	result := handler.manager.Process(store.Message{constant.SET, value})
	response.Write([]byte(result))
}

func (handler HashHttpHandler) Get(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)

	value := make(map[string]string)
	value[constant.KEY] = vars[constant.KEY]
	value[constant.FIELD] = vars[constant.FIELD]

	result := handler.manager.Process(store.Message{constant.GET, value})
	response.Write([]byte(result))
}

func (handler HashHttpHandler) GetHandles() []store.HttpHandle {
	return []store.HttpHandle {
		{http.MethodPost, constant.URL_HASH_SET, handler.Set},
		{http.MethodGet, constant.URL_HASH_GET, handler.Get},
	}
}