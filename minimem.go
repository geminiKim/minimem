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

	hashStore := store.NewHashStore()
	hashHandler := store.NewHashHttpHandler(hashStore)

	httpServerStart("8011", stringHandler, listHandler, hashHandler)
}

func httpServerStart(port string, stringHandler *store.StringHttpHandler, listHandler *store.ListHttpHandler, hashHandler *store.HashHttpHandler) {
	server := mux.NewRouter()
	server.HandleFunc("/string/{key}", stringHandler.Set).Methods("POST")
	server.HandleFunc("/string/{key}", stringHandler.Get).Methods("GET")

	server.HandleFunc("/list/{key}/leftPush", listHandler.LeftPush).Methods("POST")
	server.HandleFunc("/list/{key}/leftPeek", listHandler.LeftPeek).Methods("GET")
	server.HandleFunc("/list/{key}/leftPop", listHandler.LeftPop).Methods("GET")
	server.HandleFunc("/list/{key}/rightPush", listHandler.RightPush).Methods("POST")
	server.HandleFunc("/list/{key}/rightPeek", listHandler.RightPeek).Methods("GET")
	server.HandleFunc("/list/{key}/rightPop", listHandler.RightPop).Methods("GET")
	server.HandleFunc("/list/{key}/rangeGet/{index}/{count}", listHandler.RangeGet).Methods("GET")

	server.HandleFunc("/hash/{key}/{field}", hashHandler.Set).Methods("POST")
	server.HandleFunc("/hash/{key}/{field}", hashHandler.Get).Methods("GET")

	log.Fatal(http.ListenAndServe(":" + port, server))
}