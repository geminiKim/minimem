package store

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
)

func Test_SetByHashHttpHandler(t *testing.T) {
	store := NewHashStore()
	storeHandler := NewHashHttpHandler(store)

	request, _ := http.NewRequest("POST", "/hash/hello/world", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/hash/{key}/{field}", storeHandler.Set).Methods("POST")
	server.ServeHTTP(recorder, request)

	hashValue := store.get("hello", "world")
	assert.Equal(t, "HelloWorld", hashValue)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_GetByHashHttpHandler(t *testing.T) {
	store := NewHashStore()
	storeHandler := NewHashHttpHandler(store)

	store.set("hello", "world", "HelloWorld")

	request, _ := http.NewRequest("GET", "/hash/hello/world", nil)

	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/hash/{key}/{field}", storeHandler.Get).Methods("GET")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}