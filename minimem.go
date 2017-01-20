package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/geminikim/minimem/store"
)

func main() {
	stringStore := store.NewStringStore()
	stringHandler := store.NewStringHttpHandler(stringStore)

	listStore := store.NewListStore()
	listHandler := store.NewListHttpHandler(listStore)

	httpServerStart("8011", stringHandler, listHandler)
}

func httpServerStart(port string, stringHandler *store.StringHttpHandler, listHandler *store.ListHttpHandler) {
	server := mux.NewRouter()
	server.HandleFunc("/string/set", stringHandler.Set)
	server.HandleFunc("/string/get/{key}", stringHandler.Get)

	server.HandleFunc("/list/leftPush", listHandler.LeftPush)
	server.HandleFunc("/list/leftPeek/{key}", listHandler.LeftPeek)
	server.HandleFunc("/list/leftPop/{key}", listHandler.LeftPop)
	server.HandleFunc("/list/rightPush", listHandler.RightPush)
	server.HandleFunc("/list/rightPeek/{key}", listHandler.RightPeek)
	server.HandleFunc("/list/rightPop/{key}", listHandler.RightPop)
	server.HandleFunc("/list/rightPop/{key}/{startIndex}/{endIndex}", listHandler.RangeGet)

	log.Fatal(http.ListenAndServe(":" + port, server))
}