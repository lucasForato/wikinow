package utils

import (
	"errors"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func ReadMarkdown(path string) ([]string, error) {
	mainInput, err := os.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{
			"source": path,
		}).Error(err)
    return nil, err
	}

	file := string(mainInput)
	return strings.Split(file, "\n"), nil
}

func GetFileTitleAndOrder(path string) (string, int64, error) {
	mainInput, err := os.ReadFile(path)
	if err != nil {
		log.WithFields(log.Fields{
			"source": path,
		}).Fatal("Error reading main documentation file.")
	}

	file := string(mainInput)
	split := strings.Split(file, "---")
	if len(split) < 2 {
		return "", 0, errors.New("File format is incorrect")
	}

	info, err := os.Stat(path)
	if err != nil {
		return "", 0, errors.New("Could not retrieve last update date from file")
	}
	order := info.ModTime().Unix()
	title := ""
	metadata := strings.Split(split[1], "\n")
	for _, line := range metadata {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "title: ") {
			title = strings.TrimPrefix(line, "title: ")
		}
	}

	if title == "" {
		return title, 0, errors.New("Every file must contain a 'title'")
	}

	return title, order, nil
}

func IndexOfSubstring(str, subStr string) int {
	for i := 0; i < len(str); i++ {
		if str[i:i+len(subStr)] == subStr {
			return i
		}
	}
	return -1
}
