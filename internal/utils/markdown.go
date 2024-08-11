package utils

import (
	"errors"
	"os"
	"strconv"
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

func GetFileTitleAndOrder(path string) (string, int, error) {
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

	metadata := strings.Split(split[1], "\n")
	title := ""
	order := -1
	for _, line := range metadata {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "title: ") {
			title = strings.TrimPrefix(line, "title: ")
		}

    if strings.HasPrefix(line, "order: ") {
      i, err := strconv.Atoi(strings.TrimPrefix(line, "order: "))
      if err != nil {
        return title, order, errors.New("Order must be an integer")
      }
      order = i
    }
	}

  if title == "" {
	  return title, 0, errors.New("Every file must contain a 'title'")
  }

  if order == -1 {
    return title, 0, errors.New("Every file must contain an 'order'")
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
