package store

import "context"

func Read(store *context.Context, key string) (string, bool) {
	stringMap, ok := (*store).Value(mapKey).(map[string]string)
	if !ok {
		return "", false
	}

	value, found := stringMap[key]
	return value, found
}
