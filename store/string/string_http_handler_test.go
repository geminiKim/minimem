package strings

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

func Test_SetByStringHttpHandler(t *testing.T) {
	manager := NewStringStoreManager()
	handler := NewStringHttpHandler(manager)

	request, _ := http.NewRequest(http.MethodPost, "/string/hello", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	value := make(map[string]string)
	value[constant.KEY] = "hello"

	assert.Equal(t, "HelloWorld", manager.Process(store.Message{constant.GET, value}))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_GetByStringHttpHandler(t *testing.T) {
	manager := NewStringStoreManager()
	handler := NewStringHttpHandler(manager)

	value := make(map[string]string)
	value[constant.KEY] = "hello"
	value[constant.VALUE] = "HelloWorld"
	manager.Process(store.Message{constant.SET, value})

	request, _ := http.NewRequest(http.MethodGet, "/string/hello", nil)
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	for _, handle := range handler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}