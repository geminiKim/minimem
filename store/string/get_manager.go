package strings

import "github.com/geminikim/minimem/store"

func GetStoreManagers() []store.Manager {
	manager := NewStringStoreManager()
	return []store.Manager{
		manager,
	}
}

func GetHttpHandler(managers []store.Manager) []store.HttpHandler {
	handlers := make([]store.HttpHandler, len(managers))
	for index, manager := range managers {
		switch manager.GetType() {
		case "STRING": handlers[index] = NewStringHttpHandler(manager)
		}
	}
	return handlers
}