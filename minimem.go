package main

import (
	"log"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store/list"
	"github.com/geminikim/minimem/store/hash"
	"github.com/geminikim/minimem/store/string"
	"net/http"
	"github.com/geminikim/minimem/handler/http"
)

func main() {
	stringStore := strings.NewStringStore()
	listStore := list.NewListStore()
	hashStore := hash.NewHashStore()

	stringHandler := strings.NewStringHttpHandler(stringStore)
	listHandler := list.NewListHttpHandler(listStore)
	hashHandler := hash.NewHashHttpHandler(hashStore)

	httpServerStart("8011", []handler.HttpHandler{stringHandler, listHandler, hashHandler})
}
func httpServerStart(port string, handlers []handler.HttpHandler) {
	server := mux.NewRouter()
	for _, handler := range handlers {
		for _, handle := range handler.GetHandles() {
			server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
		}
	}
	log.Fatal(http.ListenAndServe(":" + port, server))
}