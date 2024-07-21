package utils

import (
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)


func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func GetDirFromCurr(path ...string) string {
	dir := GetCurrentDir()

	fullPath := append([]string{dir}, path...)

	newDir := filepath.Join(fullPath...)
	return newDir
}

func DirExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}
