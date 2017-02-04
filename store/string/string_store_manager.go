package strings

import "github.com/geminikim/minimem/store"

type StringStoreManager struct {
	store *stringStore
}

func (manager StringStoreManager) Process(message store.Message) string {
	switch message.Command {
	case "GET": return manager.store.get(message.Value["key"])
	case "SET": return manager.store.set(message.Value["key"], message.Value["value"])
	default: return "Not Supported Command"
	}
}

func NewStringStoreManager(store *stringStore) *StringStoreManager {
	manager := new(StringStoreManager)
	manager.store = store
	return manager
}