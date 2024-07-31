package utils

import (
	"context"
	"fmt"
	"regexp"
	"strings"
)

type contextKey string

const mapKey contextKey = "storageContext"

type Context struct {
	ctx context.Context
}

func CreateContext() Context {
	ctx := context.Background()
	stringMap := map[string]string{}
	ctx = context.WithValue(ctx, mapKey, stringMap)
	return Context{ctx}
}

func (c *Context) Load(lines *[]string) {
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
				c.Add(strings.Trim(split[0], " "), strings.Trim(split[1], " "))
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
		c.Add(strings.Trim(match[1], " "), strings.Trim(match[2], " "))
	}
}

func (c *Context) Add(key, value string) {
	stringMap, ok := c.ctx.Value(mapKey).(map[string]string)
	if !ok {
		stringMap = make(map[string]string)
	}

	newMap := make(map[string]string)
	for k, v := range stringMap {
		newMap[k] = v
	}
	newMap[key] = value

	c.ctx = context.WithValue(c.ctx, mapKey, newMap)
}

func (c *Context) Remove(key string) {
	stringMap, ok := c.ctx.Value(mapKey).(map[string]string)
	if !ok {
		stringMap = make(map[string]string)
	}

	newMap := make(map[string]string)
	for k, v := range stringMap {
		if k != key {
			newMap[k] = v
		}
	}

	c.ctx = context.WithValue(c.ctx, mapKey, newMap)
}

func (c *Context) Get(key string) (string, bool) {
	stringMap, ok := c.ctx.Value(mapKey).(map[string]string)
	if !ok {
		return "", false
	}

	value, found := stringMap[key]
	return value, found
}

func (c *Context) Print() {
	fmt.Println("Store {")
	for k, v := range c.ctx.Value(mapKey).(map[string]string) {
		fmt.Println("  ", k, ":", v)
	}
	fmt.Println("}")
}
