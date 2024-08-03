package store

import "context"

func Create() *Store {
	store := context.Background()
	stringMap := map[string]string{}
	store = context.WithValue(store, mapKey, stringMap)
	return &store
}
