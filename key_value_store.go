package minimem

type keyValueStore struct {
	keyValueMap map[string]string
}

func (store keyValueStore) put(key string, value string) {
	store.keyValueMap[key] = value
}

func (store keyValueStore) get(key string) string {
	return store.keyValueMap[key]
}

func NewKeyValueStore() *keyValueStore {
	store := new(keyValueStore)
	store.keyValueMap = make(map[string]string)
	return store
}