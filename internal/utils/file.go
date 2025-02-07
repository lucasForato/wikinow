package utils

import (
	"errors"
	"os"
	"strings"
	"wikinow/infra/logger"
)

func GetFileTitleAndOrder(path string) (string, int64, error) {
	mainInput, err := os.ReadFile(path)
	if err != nil {
    logger.Error("Error reading main documentation file.", err)
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

func GetFileTitleAndContent(path string) (string, string, error) {
	lines, err := ReadMarkdown(path)
	if err != nil {
		return "", "", err
	}

	title, err := GetFileTitle(path)
	if err != nil {
		return "", "", err
	}

	return title, strings.Join(lines, "\n"), nil
}

func GetFileTitle(path string) (string, error) {
	lines, err := ReadMarkdown(path)
	if err != nil {
		return "", err
	}

	iStart, iEnd, err := GetMetadataStartAndEnd(lines)
	if err != nil {
		return "", err
	}

	for _, line := range lines[iStart:iEnd] {
		if strings.HasPrefix(line, "title: ") {
			return strings.TrimPrefix(line, "title: "), nil
		}
	}
	return "", errors.New("No title found in metadata")
}

func GetMetadataStartAndEnd(lines []string) (int, int, error) {
	if len(lines) == 0 {
		return 0, 0, errors.New("File is empty")
	}

	iStart := 0
	iEnd := 0

	if !strings.Contains(lines[iStart], "---") {
		return 0, 0, errors.New("No metadata found in file")
	}

	for i, line := range lines[1:] {
		if strings.Contains(line, "---") {
			iEnd = i + 1
		}
	}

	return iStart, iEnd, nil
}

func FindInFile(path string, query string) (bool, error) {
	if strings.Contains(strings.ToLower(path), strings.ToLower(query)) {
		return true, nil
	}

	lines, err := ReadMarkdown(path)
	if err != nil {
		return false, err
	}
	for _, line := range lines {
		if strings.Contains(Normalize(line), query) {
			return true, nil
		}
	}
	return false, nil
}
