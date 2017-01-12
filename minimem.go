package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store"
)

func main() {
	keyValueStore := store.NewKeyValueStore()
	storeHandler := store.NewStoreHttpHandler(keyValueStore)
	httpServerStart("8011", storeHandler)
}

func httpServerStart(port string, handler *store.StoreHttpHandler) {
	server := mux.NewRouter() // HTTP 요청 멀티플렉서 생성
	server.HandleFunc("/store/map/set", handler.Set)
	server.HandleFunc("/store/map/get/{key}", handler.Get)
	log.Fatal(http.ListenAndServe(":" + port, server))
}