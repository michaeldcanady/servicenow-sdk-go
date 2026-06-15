package internal

import "github.com/microsoft/kiota-abstractions-go/store"

type mockNilBSModel struct{ BaseModel }

func (m *mockNilBSModel) GetBackingStore() store.BackingStore { return nil }
func (m *mockNilBSModel) GetBackingStoreFactory() store.BackingStoreFactory {
	return store.NewInMemoryBackingStore
}
func (m *mockNilBSModel) SetBackingStoreFactory(f store.BackingStoreFactory) error { return nil }
