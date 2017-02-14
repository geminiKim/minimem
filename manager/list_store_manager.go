package manager

import (
	"github.com/geminikim/minimem/util"
	"github.com/geminikim/minimem/constant"
	"github.com/geminikim/minimem/store"
)

type ListStoreManager struct {
	store *store.ListStore
}

func NewListStoreManager() Manager {
	manager := new(ListStoreManager)
	manager.store = store.NewListStore()
	return manager
}

func (manager ListStoreManager) Process(message Message) string {
	switch message.Command {
	case constant.LEFT_PUSH: return manager.store.LeftPush(message.Value[constant.KEY], message.Value[constant.VALUE])
	case constant.LEFT_PEEK: return manager.store.LeftPeek(message.Value[constant.KEY])
	case constant.LEFT_POP: return manager.store.LeftPop(message.Value[constant.KEY])
	case constant.RIGHT_PUSH: return manager.store.RightPush(message.Value[constant.KEY], message.Value[constant.VALUE])
	case constant.RIGHT_PEEK: return manager.store.RightPeek(message.Value[constant.KEY])
	case constant.RIGHT_POP: return manager.store.RightPop(message.Value[constant.KEY])
	case constant.BY_RANGE: return util.ToJsonString(manager.store.ByRange(message.Value[constant.KEY], util.GetInt(message.Value[constant.INDEX]), util.GetInt(message.Value[constant.COUNT])))
	default: return constant.NOT_SUPPORTED_COMMAND
	}
}

func (manager ListStoreManager) GetType() string {
	return constant.LIST
}