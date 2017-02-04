package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/store/factory"
)

func main() {
	managers := factory.GetManagers()
	handlers := factory.GetHttpHandlers(managers)

	httpServerStart("8011", handlers)
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