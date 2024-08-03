package store

import (
	"regexp"
	"strings"
)

func Load(store *Store, lines *[]string) {
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
				Update(store, strings.Trim(split[0], " "), strings.Trim(split[1], " "))
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
		Update(store, strings.Trim(match[1], " "), strings.Trim(match[2], " "))
	}
}
