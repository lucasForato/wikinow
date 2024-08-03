package store

import "context"

func Update(store *context.Context, key, value string) {
	stringMap, ok := (*store).Value(mapKey).(map[string]string)
	if !ok {
		stringMap = make(map[string]string)
	}

	newMap := make(map[string]string)
	for k, v := range stringMap {
		newMap[k] = v
	}
	newMap[key] = value

	*store = context.WithValue(*store, mapKey, newMap)
}
