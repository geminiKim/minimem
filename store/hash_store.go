package store

import "github.com/geminikim/minimem/constant"

type HashStore struct {
	hashMap map[string]map[string]string
}

func NewHashStore() *HashStore {
	store := new(HashStore)
	store.hashMap = make(map[string]map[string]string)
	return store
}

func (store HashStore) Set(key string, field string, value string) string {
	if store.hashMap[key] == nil {
		store.hashMap[key] = make(map[string]string);
	}
	store.hashMap[key][field] = value;
	return constant.OK
}

func (store HashStore) Get(key string, field string) string {
	return store.hashMap[key][field]
}
