package parser

import (
	"context"
	"errors"
	"fmt"
	"html/template"
	"regexp"
	"slices"
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

func ReadCtx(c *Ctx, key string) (string, bool) {
	stringMap, ok := (*c).Value(mapKey).(map[string]string)
	if !ok {
		return "", false
	}

	value, found := stringMap[key]
	return value, found
}

func ReadCtxSkipError(c *Ctx, key string) string {
  value, _ := ReadCtx(c, key)
  return value
}

func LoadCtx(c *Ctx, lines *[]string) error {
	var metadataStart int
	var metadataEnd int
  var keys []string

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
				if _, ok := ReadCtx(c, key); ok {
					return errors.New(fmt.Sprintf("Duplicate key: %s", key))
				}
				parsedValue := ParseInline(value, c)
				if parsedValue != template.HTML(value) {
					return errors.New(fmt.Sprintf("value can only contain text: %s", value))
				}
        keys = append(keys, key)
				UpdateCtx(c, key, value)
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
		if _, ok := ReadCtx(c, key); ok {
			return errors.New(fmt.Sprintf("Duplicate key: %s", key))
		}
		parsedValue := ParseInline(value, c)
		if parsedValue != template.HTML(value) {
			return errors.New(fmt.Sprintf("value can only contain text: %s", value))
		}

    keys = append(keys, key)
		UpdateCtx(c, key, value)
	}

  if slices.Contains(keys, "title") == false || len(keys) == 0 {
    return errors.New("title must be set")
  }

	return nil
}

func UpdateCtx(c *Ctx, key, value string) {
	stringMap, ok := (*c).Value(mapKey).(map[string]string)
	if !ok {
		stringMap = make(map[string]string)
	}

	newMap := make(map[string]string)
	for k, v := range stringMap {
		newMap[k] = v
	}
	newMap[key] = value

	*c = context.WithValue(*c, mapKey, newMap)
}

func PrintCtx(c *Ctx) {
	fmt.Println("Store {")
	for k, v := range (*c).Value(mapKey).(map[string]string) {
		fmt.Println("  ", k, ":", v)
	}
	fmt.Println("}")
}
