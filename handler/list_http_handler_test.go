package handler

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/constant"
	"github.com/geminikim/minimem/manager"
)

func Test_LeftPushByListHttpHandler(t *testing.T) {
	storeManager := manager.NewListStoreManager()
	storeHandler := NewListHttpHandler(storeManager)

	request, _ := http.NewRequest(http.MethodPost, "/list/hello/leftPush", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)


	value := make(map[string]string)
	value[constant.KEY] = "hello"

	assert.Equal(t, "HelloWorld", storeManager.Process(manager.Message{constant.LEFT_POP, value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_LeftPeekByListHttpHandler(t *testing.T) {
	storeManager := manager.NewListStoreManager()
	storeHandler := NewListHttpHandler(storeManager)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	storeManager.Process(manager.Message{constant.LEFT_PUSH, value})

	request, _ := http.NewRequest(http.MethodGet, "/list/hello/leftPeek", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_LeftPopByListHttpHandler(t *testing.T) {
	storeManager := manager.NewListStoreManager()
	storeHandler := NewListHttpHandler(storeManager)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	storeManager.Process(manager.Message{constant.LEFT_PUSH, value})

	request, _ := http.NewRequest(http.MethodGet, "/list/hello/leftPop", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPushByListHttpHandler(t *testing.T) {
	storeManager := manager.NewListStoreManager()
	storeHandler := NewListHttpHandler(storeManager)

	request, _ := http.NewRequest(http.MethodPost, "/list/hello/rightPush", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	value := make(map[string]string)
	value[constant.KEY] = "hello"

	assert.Equal(t, "HelloWorld", storeManager.Process(manager.Message{constant.RIGHT_POP, value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPeekByListHttpHandler(t *testing.T) {
	storeManager := manager.NewListStoreManager()
	storeHandler := NewListHttpHandler(storeManager)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	storeManager.Process(manager.Message{constant.RIGHT_PUSH, value})

	request, _ := http.NewRequest(http.MethodGet, "/list/hello/rightPeek", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPopByListHttpHandler(t *testing.T) {
	storeManager := manager.NewListStoreManager()
	storeHandler := NewListHttpHandler(storeManager)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	storeManager.Process(manager.Message{constant.RIGHT_PUSH, value})

	request, _ := http.NewRequest(http.MethodGet, "/list/hello/rightPop", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RangeByListHttpHandler(t *testing.T) {
	storeManager := manager.NewListStoreManager()
	storeHandler := NewListHttpHandler(storeManager)


	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld_0"
	storeManager.Process(manager.Message{constant.RIGHT_PUSH, value})
	value[constant.VALUE] = "HelloWorld_1"
	storeManager.Process(manager.Message{constant.RIGHT_PUSH, value})
	value[constant.VALUE] = "HelloWorld_2"
	storeManager.Process(manager.Message{constant.RIGHT_PUSH, value})
	value[constant.VALUE] = "HelloWorld_3"
	storeManager.Process(manager.Message{constant.RIGHT_PUSH, value})
	value[constant.VALUE] = "HelloWorld_4"
	storeManager.Process(manager.Message{constant.RIGHT_PUSH, value})

	request, _ := http.NewRequest(http.MethodGet, "/list/hello/1/3", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "[\"HelloWorld_1\",\"HelloWorld_2\"]", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

