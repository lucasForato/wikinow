package utils

import (
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ReadMarkdown(path string) []string {
	mainInput, err := os.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{
			"source": path,
		}).Fatal("Error reading main documentation file.")
	}

	file := string(mainInput)
	return strings.Split(file, "\n")
}

func IndexOfSubstring(str, subStr string) int {
	for i := 0; i < len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			return i
		}
	}
	return -1
}
