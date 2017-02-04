package list

import (
	"github.com/geminikim/minimem/store"
	"github.com/geminikim/minimem/util"
	"github.com/geminikim/minimem/constant"
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
	case constant.LEFT_PUSH: return manager.store.leftPush(message.Value[constant.KEY], message.Value[constant.VALUE])
	case constant.LEFT_PEEK: return manager.store.leftPeek(message.Value[constant.KEY])
	case constant.LEFT_POP: return manager.store.leftPop(message.Value[constant.KEY])
	case constant.RIGHT_PUSH: return manager.store.rightPush(message.Value[constant.KEY], message.Value[constant.VALUE])
	case constant.RIGHT_PEEK: return manager.store.rightPeek(message.Value[constant.KEY])
	case constant.RIGHT_POP: return manager.store.rightPop(message.Value[constant.KEY])
	case constant.BY_RANGE: return manager.store.byRange(message.Value[constant.KEY], util.GetInt(message.Value[constant.INDEX]), util.GetInt(message.Value[constant.COUNT]))
	default: return constant.NOT_SUPPORTED_COMMAND
	}
}

func (manager ListStoreManager) GetType() string {
	return constant.LIST
}