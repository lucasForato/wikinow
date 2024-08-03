package parser 

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type Ctx = context.Context

type contextKey string
const mapKey contextKey = "storageContext"

func CreateCtx() *Ctx {
	store := context.Background()
	stringMap := map[string]string{}
	store = context.WithValue(store, mapKey, stringMap)
	return &store
}

func ReadCtx(ctx *Ctx, key string) (string, bool) {
	stringMap, ok := (*ctx).Value(mapKey).(map[string]string)
	if !ok {
		return "", false
	}

	value, found := stringMap[key]
	return value, found
}

func LoadCtx(ctx *Ctx, lines *[]string) error {
	var metadataStart int
	var metadataEnd int

	if len(*lines) == 0 {
		return errors.New("Empty file")
	}

	if (*lines)[0] == "---" {
		metadataStart = 1
		for i := metadataStart; i < len(*lines); i++ {
			if (*lines)[i] == "---" {
				metadataEnd = i
				break
			}
		}
	}

	if metadataEnd != 0 {
		metadataLines := (*lines)[metadataStart:metadataEnd]
		for _, line := range metadataLines {
			split := strings.SplitN(line, ":", 2)
			if len(split) == 2 {
				key := strings.Trim(split[0], " ")
				value := strings.Trim(split[1], " ")
				if _, ok := ReadCtx(ctx, key); ok {
					return errors.New(fmt.Sprintf("Duplicate key: %s", key))
				}

				UpdateCtx(ctx, key, value)
			}
		}
	}

	// regex to find link definitions
	re := regexp.MustCompile(`\[(.*?)\]:\s*(.*)`)
	for _, key := range *lines {
		match := re.FindStringSubmatch(key)
		if match == nil {
			continue
		}
		key = strings.Trim(match[1], " ")
		value := strings.Trim(match[2], " ")
		if _, ok := ReadCtx(ctx, key); ok {
			return errors.New(fmt.Sprintf("Duplicate key: %s", key))
		}

		UpdateCtx(ctx, key, value)
	}
	return nil
}

func UpdateCtx(ctx *Ctx, key, value string) {
	stringMap, ok := (*ctx).Value(mapKey).(map[string]string)
	if !ok {
		stringMap = make(map[string]string)
	}

	newMap := make(map[string]string)
	for k, v := range stringMap {
		newMap[k] = v
	}
	newMap[key] = value

	*ctx = context.WithValue(*ctx, mapKey, newMap)
}

func PrintCtx(ctx *Ctx) {
	fmt.Println("Store {")
	for k, v := range (*ctx).Value(mapKey).(map[string]string) {
		fmt.Println("  ", k, ":", v)
	}
	fmt.Println("}")
}
