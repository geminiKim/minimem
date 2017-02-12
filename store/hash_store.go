package hash

import "github.com/geminikim/minimem/constant"

type hashStore struct {
	hashMap map[string]map[string]string
}

func NewHashStore() *hashStore {
	store := new(hashStore)
	store.hashMap = make(map[string]map[string]string)
	return store
}

func (store hashStore) set(key string, field string, value string) string {
	if store.hashMap[key] == nil {
		store.hashMap[key] = make(map[string]string);
	}
	store.hashMap[key][field] = value;
	return constant.OK
}

func (store hashStore) get(key string, field string) string {
	return store.hashMap[key][field]
}
