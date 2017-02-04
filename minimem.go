package main

import (
	"log"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store/list"
	"github.com/geminikim/minimem/store/hash"
	"net/http"
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/store/factory"
)

func main() {
	managers := factory.GetStoreManagers()
	handlers := factory.GetHttpHandler(managers)

	listStore := list.NewListStore()
	listHandler := list.NewListHttpHandler(listStore)

	hashStore := hash.NewHashStore()
	hashHandler := hash.NewHashHttpHandler(hashStore)

	httpServerStart("8011", []store.HttpHandler{handlers[0], listHandler, hashHandler})
}
func httpServerStart(port string, handlers []store.HttpHandler) {
	server := mux.NewRouter()
	for _, handler := range handlers {
		for _, handle := range handler.GetHandles() {
			server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
		}
	}
	log.Fatal(http.ListenAndServe(":" + port, server))
}