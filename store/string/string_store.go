package strings

type stringStore struct {
	stringMap map[string]string
}

func (store stringStore) set(key string, value string) string {
	store.stringMap[key] = value
	return "OK"
}

func (store stringStore) get(key string) string {
	return store.stringMap[key]
}

func NewStringStore() *stringStore {
	store := new(stringStore)
	store.stringMap = make(map[string]string)
	return store
}