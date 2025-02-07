package utils

import (
	"os"
	"strings"
	"wikinow/infra/logger"

)

func ReadMarkdown(path string) ([]string, error) {
	mainInput, err := os.ReadFile(path)
	if err != nil {
		logger.Error(err, map[string]string{"source": path})
		return nil, err
	}

	file := string(mainInput)
	return strings.Split(file, "\n"), nil
}
