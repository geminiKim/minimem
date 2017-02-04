package list

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store"
)

func Test_LeftPushByListHttpHandler(t *testing.T) {
	manager := NewListStoreManager()
	handler := NewListHttpHandler(manager)

	request, _ := http.NewRequest("POST", "/list/hello/leftPush", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)


	value := make(map[string]string)
	value["key"] = "hello"

	assert.Equal(t, "HelloWorld", manager.Process(store.Message{"LEFT_POP", value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_LeftPeekByListHttpHandler(t *testing.T) {
	manager := NewListStoreManager()
	handler := NewListHttpHandler(manager)

	value := make(map[string]string)
	value["key"] = "hello"
	value["value"] = "HelloWorld"
	manager.Process(store.Message{"LEFT_PUSH", value})

	request, _ := http.NewRequest("GET", "/list/hello/leftPeek", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_LeftPopByListHttpHandler(t *testing.T) {
	manager := NewListStoreManager()
	handler := NewListHttpHandler(manager)

	value := make(map[string]string)
	value["key"] = "hello"
	value["value"] = "HelloWorld"
	manager.Process(store.Message{"LEFT_PUSH", value})

	request, _ := http.NewRequest("GET", "/list/hello/leftPop", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPushByListHttpHandler(t *testing.T) {
	manager := NewListStoreManager()
	handler := NewListHttpHandler(manager)

	request, _ := http.NewRequest("POST", "/list/hello/rightPush", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	value := make(map[string]string)
	value["key"] = "hello"

	assert.Equal(t, "HelloWorld", manager.Process(store.Message{"RIGHT_POP", value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPeekByListHttpHandler(t *testing.T) {
	manager := NewListStoreManager()
	handler := NewListHttpHandler(manager)

	value := make(map[string]string)
	value["key"] = "hello"
	value["value"] = "HelloWorld"
	manager.Process(store.Message{"RIGHT_PUSH", value})

	request, _ := http.NewRequest("GET", "/list/hello/rightPeek", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPopByListHttpHandler(t *testing.T) {
	manager := NewListStoreManager()
	handler := NewListHttpHandler(manager)

	value := make(map[string]string)
	value["key"] = "hello"
	value["value"] = "HelloWorld"
	manager.Process(store.Message{"RIGHT_PUSH", value})

	request, _ := http.NewRequest("GET", "/list/hello/rightPop", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RangeByListHttpHandler(t *testing.T) {
	manager := NewListStoreManager()
	handler := NewListHttpHandler(manager)


	value := make(map[string]string)
	value["key"] = "hello"
	value["value"] = "HelloWorld_0"
	manager.Process(store.Message{"RIGHT_PUSH", value})
	value["value"] = "HelloWorld_1"
	manager.Process(store.Message{"RIGHT_PUSH", value})
	value["value"] = "HelloWorld_2"
	manager.Process(store.Message{"RIGHT_PUSH", value})
	value["value"] = "HelloWorld_3"
	manager.Process(store.Message{"RIGHT_PUSH", value})
	value["value"] = "HelloWorld_4"
	manager.Process(store.Message{"RIGHT_PUSH", value})

	request, _ := http.NewRequest("GET", "/list/hello/1/3", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "[\"HelloWorld_1\",\"HelloWorld_2\",\"HelloWorld_3\"]", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

