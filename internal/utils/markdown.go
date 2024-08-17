package utils

import (
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
