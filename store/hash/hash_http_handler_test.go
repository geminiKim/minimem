package hash

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/constant"
)

func Test_SetByHashHttpHandler(t *testing.T) {
	manager := NewHashStoreManager()
	storeHandler := NewHashHttpHandler(manager)

	request, _ := http.NewRequest(http.MethodPost, "/hash/hello/world", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc(constant.URL_HASH_SET, storeHandler.Set).Methods(http.MethodPost)
	server.ServeHTTP(recorder, request)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.FIELD] = "world"

	assert.Equal(t, "HelloWorld", manager.Process(store.Message{constant.GET, value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_GetByHashHttpHandler(t *testing.T) {
	manager := NewHashStoreManager()
	storeHandler := NewHashHttpHandler(manager)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.FIELD] = "world"
	value[constant.VALUE] = "HelloWorld"
	manager.Process(store.Message{"SET", value})

	request, _ := http.NewRequest(http.MethodGet, "/hash/hello/world", nil)
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc(constant.URL_HASH_GET, storeHandler.Get).Methods(http.MethodGet)
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}