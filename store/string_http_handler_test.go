package store

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
)

func Test_SetByStringHttpHandler(t *testing.T) {
	store := NewStringStore()
	storeHandler := NewStringHttpHandler(store)

	request, _ := http.NewRequest("POST", "/string/hello", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/string/{key}", storeHandler.Set).Methods("POST")
	server.ServeHTTP(recorder, request)

	hashValue := store.get("hello")
	assert.Equal(t, "HelloWorld", hashValue)
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_GetByStringHttpHandler(t *testing.T) {
	store := NewStringStore()
	storeHandler := NewStringHttpHandler(store)

	store.set("hello", "HelloWorld")

	request, _ := http.NewRequest("GET", "/string/hello", nil)

	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/string/{key}", storeHandler.Get).Methods("GET")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}