package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
)

type TestPair struct {
	Input    string
	Expected []string
}

func GetTestData(dir ...string) []TestPair {
	path := GetDirFromCurr(dir...)
	lines := ReadMarkdown(path)

	result := []TestPair{}

	for i := 0; i < len(lines)-1; i += 2 {
		testPair := new(TestPair)
		testPair.Input = string(lines[i])
		expected := strings.Split(lines[i+1], ", ")
		testPair.Expected = expected
		result = append(result, *testPair)
	}

	return result
}

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

func JsonPrettyPrint(in string) *error {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "  ")
	if err != nil {
		return &err
	}
	fmt.Print(out.String())
	return nil
}
