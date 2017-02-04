package store

func getDataStructure() []DataStructure {
	return []DataStructure {
		{"STRING", newTestStore},
	}
}

type DataStructure struct {
	Structure string
	Store func() DataStore
}

func newTestStore() DataStore {
	store := new(testStore)
	store.Initialize()
	return store
}

type testStore struct {
}

func (store testStore) Initialize() {
}

func (store testStore) GetStructureType() string {
	return "TEST"
}

type DataStore interface {
	Initialize()
	GetStructureType() string
}