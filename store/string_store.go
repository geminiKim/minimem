package store

import "github.com/geminikim/minimem/constant"

type StringStore struct {
	stringMap map[string]string
}

func NewStringStore() *StringStore {
	store := new(StringStore)
	store.stringMap = make(map[string]string)
	return store
}

func (store StringStore) Set(key string, value string) string {
	store.stringMap[key] = value
	return constant.OK
}

func (store StringStore) Get(key string) string {
	return store.stringMap[key]
}