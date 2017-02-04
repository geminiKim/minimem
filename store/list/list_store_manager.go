package list

import (
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/util"
)

type ListStoreManager struct {
	store *listStore
}

func NewListStoreManager() store.Manager {
	manager := new(ListStoreManager)
	manager.store = NewListStore()
	return manager
}

func (manager ListStoreManager) Process(message store.Message) string {
	switch message.Command {
	case "LEFT_PUSH": return manager.store.leftPush(message.Value["key"], message.Value["value"])
	case "LEFT_PEEK": return manager.store.leftPeek(message.Value["key"])
	case "LEFT_POP": return manager.store.leftPop(message.Value["key"])
	case "RIGHT_PUSH": return manager.store.rightPush(message.Value["key"], message.Value["value"])
	case "RIGHT_PEEK": return manager.store.rightPeek(message.Value["key"])
	case "RIGHT_POP": return manager.store.rightPop(message.Value["key"])
	case "BY_RANGE": return manager.store.byRange(message.Value["key"], util.GetInt(message.Value["index"]), util.GetInt(message.Value["count"]))
	default: return "Not Supported Command"
	}
}

func (manager ListStoreManager) GetType() string {
	return "LIST"
}