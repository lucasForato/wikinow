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
