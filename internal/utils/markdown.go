package utils

import (
	"errors"
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

func GetFileTitle(path string) (string, error) {
	mainInput, err := os.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{
			"source": path,
		}).Fatal("Error reading main documentation file.")
	}

	file := string(mainInput)
	split := strings.Split(file, "---")
	if len(split) < 2 {
		return "", errors.New("File format is incorrect")
	}

	metadata := strings.Split(split[1], "\n")
	for _, line := range metadata {
		line = strings.TrimSpace(line) // Remove any leading/trailing whitespace
		if strings.HasPrefix(line, "title: ") {
			title := strings.TrimPrefix(line, "title: ")
			return strings.TrimSpace(title), nil
		}
	}
	return "", errors.New("Every file must contain a 'title'")
}


func IndexOfSubstring(str, subStr string) int {
	for i := 0; i < len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			return i
		}
	}
	return -1
}
