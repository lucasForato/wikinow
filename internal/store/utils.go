package store

import (
	"context"
	"fmt"
)

func Print(store *context.Context) {
	fmt.Println("Store {")
	for k, v := range (*store).Value(mapKey).(map[string]string) {
		fmt.Println("  ", k, ":", v)
	}
	fmt.Println("}")
}
