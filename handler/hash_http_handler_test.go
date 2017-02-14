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

func Test_SetByHashHttpHandler(t *testing.T) {
	storeManager := manager.NewHashStoreManager()
	storeHandler := NewHashHttpHandler(storeManager)

	request, _ := http.NewRequest(http.MethodPost, "/hash/hello/world", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc(constant.URL_HASH_SET, storeHandler.Set).Methods(http.MethodPost)
	server.ServeHTTP(recorder, request)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.FIELD] = "world"

	assert.Equal(t, "HelloWorld", storeManager.Process(manager.Message{constant.GET, value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_GetByHashHttpHandler(t *testing.T) {
	storeManager := manager.NewHashStoreManager()
	storeHandler := NewHashHttpHandler(storeManager)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.FIELD] = "world"
	value[constant.VALUE] = "HelloWorld"
	storeManager.Process(manager.Message{"SET", value})

	request, _ := http.NewRequest(http.MethodGet, "/hash/hello/world", nil)
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc(constant.URL_HASH_GET, storeHandler.Get).Methods(http.MethodGet)
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}