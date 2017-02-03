package main

import (
	"log"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store/list"
	"github.com/geminikim/minimem/store/hash"
	"github.com/geminikim/minimem/store/string"
	"net/http"
	"github.com/geminikim/minimem/store"
)

func main() {
	stringStore := strings.NewStringStore()
	stringHandler := strings.NewStringHttpHandler(stringStore)

	listStore := list.NewListStore()
	listHandler := list.NewListHttpHandler(listStore)

	hashStore := hash.NewHashStore()
	hashHandler := hash.NewHashHttpHandler(hashStore)

	httpServerStart("8011", stringHandler, listHandler, hashHandler)
}
func httpServerStart(port string, stringHandler store.HttpHandler, listHandler *list.ListHttpHandler, hashHandler *hash.HashHttpHandler) {
	server := mux.NewRouter()

	for _, handle := range stringHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}

	server.HandleFunc("/list/{key}/leftPush", listHandler.LeftPush).Methods("POST")
	server.HandleFunc("/list/{key}/leftPeek", listHandler.LeftPeek).Methods("GET")
	server.HandleFunc("/list/{key}/leftPop", listHandler.LeftPop).Methods("GET")
	server.HandleFunc("/list/{key}/rightPush", listHandler.RightPush).Methods("POST")
	server.HandleFunc("/list/{key}/rightPeek", listHandler.RightPeek).Methods("GET")
	server.HandleFunc("/list/{key}/rightPop", listHandler.RightPop).Methods("GET")
	server.HandleFunc("/list/{key}/{index}/{count}", listHandler.RangeGet).Methods("GET")

	server.HandleFunc("/hash/{key}/{field}", hashHandler.Set).Methods("POST")
	server.HandleFunc("/hash/{key}/{field}", hashHandler.Get).Methods("GET")

	for _, handle := range stringHandler.GetHandles() {
		server.HandleFunc(handle.Path, handle.Function).Methods(handle.Method)
	}

	log.Fatal(http.ListenAndServe(":" + port, server))
}