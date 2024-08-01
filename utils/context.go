package utils

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

type contextKey string

const mapKey contextKey = "storageContext"


func CreateContext() *context.Context {
	ctx := context.Background()
	stringMap := map[string]string{}
	ctx = context.WithValue(ctx, mapKey, stringMap)
	return &ctx
}

func LoadContext(ctx *context.Context, lines *[]string) {
	var metadataStart int
	var metadataEnd int

	if len(*lines) == 0 {
		return
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
				AddToContext(ctx, strings.Trim(split[0], " "), strings.Trim(split[1], " "))
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
		AddToContext(ctx, strings.Trim(match[1], " "), strings.Trim(match[2], " "))
	}
}

func AddToContext(ctx *context.Context, key, value string) {
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

func GetFromContext(ctx *context.Context, key string) (string, bool) {
	stringMap, ok := (*ctx).Value(mapKey).(map[string]string)
	if !ok {
		return "", false
	}

	value, found := stringMap[key]
	return value, found
}

func PrintContext(ctx *context.Context) {
	fmt.Println("Store {")
	for k, v := range (*ctx).Value(mapKey).(map[string]string) {
		fmt.Println("  ", k, ":", v)
	}
	fmt.Println("}")
}
