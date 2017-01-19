package store

type ListStoreHttpHandler struct {
	store *listStore
}

func NewListStoreHttpHandler(store *listStore) *ListStoreHttpHandler {
	handler := new(ListStoreHttpHandler)
	handler.store = store
	return handler
}