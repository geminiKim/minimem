package factory

import (
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/store/string"
	"github.com/geminikim/minimem/store/list"
	"github.com/geminikim/minimem/store/hash"
)

func GetManagers() []store.Manager {
	return []store.Manager{
		strings.NewStringStoreManager(),
		list.NewListStoreManager(),
		hash.NewHashStoreManager(),
	}
}

func GetHttpHandlers(managers []store.Manager) []store.HttpHandler {
	handlers := make([]store.HttpHandler, len(managers))
	for index, manager := range managers {
		switch manager.GetType() {
		case "STRING": handlers[index] = strings.NewStringHttpHandler(manager)
		}
	}
	return handlers
}