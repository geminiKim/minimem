package strings

//import (
//	"net/http/httptest"
//	"testing"
//	"net/http"
//	"bytes"
//	"github.com/stretchr/testify/assert"
//	"github.com/gorilla/mux"
//	"github.com/geminikim/minimem/store"
//)
//
//func Test_SetByStringHttpHandler(t *testing.T) {
//	stringManager := NewStringStoreManager()
//	storeHandler := NewStringHttpHandler(stringManager)
//
//	request, _ := http.NewRequest("POST", "/string/hello", bytes.NewBufferString("HelloWorld"))
//	recorder := httptest.NewRecorder()
//
//	server := mux.NewRouter()
//	server.HandleFunc("/string/{key}", storeHandler.Set).Methods("POST")
//	server.ServeHTTP(recorder, request)
//
//	value := make(map[string]string)
//	value["key"] = "hello"
//
//	assert.Equal(t, "HelloWorld", stringManager.Process(store.Message{"GET", value}))
//	assert.Equal(t, http.StatusOK, recorder.Code)
//}
//
//func Test_GetByStringHttpHandler(t *testing.T) {
//	stringManager := NewStringStoreManager()
//	storeHandler := NewStringHttpHandler(stringManager)
//
//	value := make(map[string]string)
//	value["key"] = "hello"
//	value["value"] = "HelloWorld"
//	stringManager.Process(store.Message{"SET", value})
//
//	request, _ := http.NewRequest("GET", "/string/hello", nil)
//
//	recorder := httptest.NewRecorder()
//
//	server := mux.NewRouter()
//	server.HandleFunc("/string/{key}", storeHandler.Get).Methods("GET")
//	server.ServeHTTP(recorder, request)
//
//	assert.Equal(t, "HelloWorld", recorder.Body.String())
//	assert.Equal(t, http.StatusOK, recorder.Code)
//}