package list

import (
	"net/http/httptest"
	"testing"
	"net/http"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
)

func Test_LeftPushByListHttpHandler(t *testing.T) {
	store := NewListStore()
	storeHandler := NewListHttpHandler(store)

	request, _ := http.NewRequest("POST", "/list/hello/leftPush", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/list/{key}/leftPush", storeHandler.LeftPush).Methods("POST")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", store.leftPop("hello"))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_LeftPeekByListHttpHandler(t *testing.T) {
	store := NewListStore()
	storeHandler := NewListHttpHandler(store)

	store.leftPush("hello", "HelloWorld")

	request, _ := http.NewRequest("GET", "/list/hello/leftPeek", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/list/{key}/leftPeek", storeHandler.LeftPeek).Methods("GET")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_LeftPopByListHttpHandler(t *testing.T) {
	store := NewListStore()
	storeHandler := NewListHttpHandler(store)

	store.leftPush("hello", "HelloWorld")

	request, _ := http.NewRequest("GET", "/list/hello/leftPop", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/list/{key}/leftPop", storeHandler.LeftPop).Methods("GET")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPushByListHttpHandler(t *testing.T) {
	store := NewListStore()
	storeHandler := NewListHttpHandler(store)

	request, _ := http.NewRequest("POST", "/list/hello/rightPush", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/list/{key}/rightPush", storeHandler.RightPush).Methods("POST")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", store.rightPop("hello"))
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPeekByListHttpHandler(t *testing.T) {
	store := NewListStore()
	storeHandler := NewListHttpHandler(store)

	store.rightPush("hello", "HelloWorld")

	request, _ := http.NewRequest("GET", "/list/hello/rightPeek", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/list/{key}/rightPeek", storeHandler.RightPeek).Methods("GET")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RightPopByListHttpHandler(t *testing.T) {
	store := NewListStore()
	storeHandler := NewListHttpHandler(store)

	store.rightPush("hello", "HelloWorld")

	request, _ := http.NewRequest("GET", "/list/hello/rightPop", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/list/{key}/rightPop", storeHandler.RightPop).Methods("GET")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "HelloWorld", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

func Test_RangeByListHttpHandler(t *testing.T) {
	store := NewListStore()
	storeHandler := NewListHttpHandler(store)

	store.rightPush("hello", "HelloWorld_0")
	store.rightPush("hello", "HelloWorld_1")
	store.rightPush("hello", "HelloWorld_2")
	store.rightPush("hello", "HelloWorld_3")
	store.rightPush("hello", "HelloWorld_4")

	request, _ := http.NewRequest("GET", "/list/hello/1/3", bytes.NewBufferString("HelloWorld"))
	recorder := httptest.NewRecorder()

	server := mux.NewRouter()
	server.HandleFunc("/list/{key}/{index}/{count}", storeHandler.RangeGet).Methods("GET")
	server.ServeHTTP(recorder, request)

	assert.Equal(t, "[\"HelloWorld_1\",\"HelloWorld_2\",\"HelloWorld_3\"]", recorder.Body.String())
	assert.Equal(t, http.StatusOK, recorder.Code)
}

