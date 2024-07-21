package utils

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func OpenWriteFile(path string) *os.File {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func AppendToFile(path string, entry string) {
	file := OpenWriteFile(path)

	_, err := file.WriteString(entry)
	if err != nil {
		log.Fatal(err)
	}

	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func RewriteFile(path string, content []string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for _, line := range content {
		line = line + "\n"
		_, err := file.WriteString(line)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BuildEntry(projectName string, projectPath string) string {
	return fmt.Sprintf("%v: \"%v\"\n", projectName, projectPath)
}
