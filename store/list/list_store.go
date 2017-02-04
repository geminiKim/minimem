package list

import (
	"github.com/geminikim/minimem/constant"
)

type listStore struct {
	listMap map[string][]string
}

func NewListStore() *listStore {
	store := new(listStore)
	store.listMap = make(map[string][]string)
	return store
}

func (store *listStore) leftPush(key string, value string) string {
	list := store.listMap[key]
	list = append(list, value)
	copy(list[1:], list[0:])
	list[0] = value
	store.listMap[key] = list
	return constant.OK
}

func (store *listStore) leftPop(key string) string {
	if len(store.listMap[key]) == 0 {
		return constant.EMPTY
	}
	list := store.listMap[key]
	value := list[0]
	store.listMap[key] = list[1:]
	return value
}

func (store *listStore) leftPeek(key string) string {
	if len(store.listMap) == 0 {
		return constant.EMPTY
	}
	return store.listMap[key][0]
}

func (store *listStore) rightPush(key string, value string) string {
	store.listMap[key] = append(store.listMap[key], value)
	return constant.OK
}

func (store *listStore) rightPop(key string) string {
	if len(store.listMap[key]) == 0 {
		return constant.EMPTY
	}
	list := store.listMap[key]
	value := list[len(list) - 1]
	store.listMap[key] = list[:len(list) - 1]
	return value
}

func (store *listStore) rightPeek(key string) string {
	if len(store.listMap) == 0 {
		return constant.EMPTY
	}
	return store.listMap[key][len(store.listMap[key])-1]
}

func (store *listStore) byRange(key string, index int, count int) []string {
	return store.listMap[key][index:count + 1]
}