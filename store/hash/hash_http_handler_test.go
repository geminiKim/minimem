package hash

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store"
)

func Test_SetByHashHttpHandler(t *testing.T) {
	manager := NewHashStoreManager()
	storeHandler := NewHashHttpHandler(manager)

	request, _ := http.NewRequest("POST", "/hash/hello/world", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/hash/{key}/{field}", storeHandler.Set).Methods("POST")
	server.ServeHTTP(recorder, request)

	value := make(map[string]string)
	value["key"] = "hello"
	value["field"] = "world"

	assert.Equal(t, "HelloWorld", manager.Process(store.Message{"GET", value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_GetByHashHttpHandler(t *testing.T) {
	manager := NewHashStoreManager()
	storeHandler := NewHashHttpHandler(manager)

	value := make(map[string]string)
	value["key"] = "hello"
	value["field"] = "world"
	value["value"] = "HelloWorld"
	manager.Process(store.Message{"SET", value})

	request, _ := http.NewRequest("GET", "/hash/hello/world", nil)
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/hash/{key}/{field}", storeHandler.Get).Methods("GET")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}