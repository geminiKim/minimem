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

func Test_SetByStringHttpHandler(t *testing.T) {
	storeManager := manager.NewStringStoreManager()
	storeHandler := NewStringHttpHandler(storeManager)

	request, _ := http.NewRequest(http.MethodPost, "/string/hello", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	value := make(map[string]string)
	value[constant.KEY] = "hello"

	assert.Equal(t, "HelloWorld", storeManager.Process(manager.Message{constant.GET, value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_GetByStringHttpHandler(t *testing.T) {
	storeManager := manager.NewStringStoreManager()
	storeHandler := NewStringHttpHandler(storeManager)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	storeManager.Process(manager.Message{constant.SET, value})

	request, _ := http.NewRequest(http.MethodGet, "/string/hello", nil)
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range storeHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}