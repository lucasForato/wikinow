// package filetree provides utility function to create and manipulate the filetree
package filetree

import (
	"errors"
	"os"
	"path/filepath"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

type TreeNodeType string

const (
	Root TreeNodeType = "Root"
	Dir  TreeNodeType = "Dir"
	File TreeNodeType = "File"
)

type TreeNode struct {
	Title        string       `json:"title"`
	IsActive     bool         `json:"isActive"`
	Order        int64        `json:"order"`
	Type         TreeNodeType `json:"type"`
	Path         string       `json:"path"`
	RelativePath string       `json:"relativePath"`
	Children     []TreeNode   `json:"children,omitempty"`
}

func GetFileTree(rootPath string, currentPath string) *TreeNode {
  currentPath = strings.Replace(currentPath, "/wiki", "", 1)
	rootNode := &TreeNode{
		Title:        "root",
		Type:         Root,
		Path:         rootPath,
		RelativePath: "./",
		Children:     []TreeNode{},
	}

	if currentPath == "/" {
		rootNode.IsActive = true
	}

	dir := filepath.Dir(rootPath + currentPath)
	err := buildTree(rootNode, rootPath, dir, rootPath+currentPath)
	if err != nil {
		log.Fatal("Error while building the file tree:", err)
	}

	return rootNode
}

func buildTree(parentNode *TreeNode, currentPath string, directory string, accessedPath string) error {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return err
	}

	if parentNode.Type == Root {
		relativePath, err := filepath.Rel(directory, currentPath)
		if err != nil {
			return errors.New("Error while getting relative path")
		}
		if strings.Contains(relativePath, ".md") {
			parentNode.RelativePath = strings.TrimRight(relativePath, ".md")
		} else {
			parentNode.RelativePath = strings.Join([]string{relativePath, "/"}, "")
		}
	}

	for _, entry := range entries {
		path := filepath.Join(currentPath, entry.Name())

		relativePath, err := filepath.Rel(directory, path)
		if err != nil {
			return errors.New("Error while getting relative path")
		}

		node := TreeNode{
			Path:     path,
			Children: []TreeNode{},
		}

		if accessedPath[len(accessedPath)-1] == '/' {
			alteredPath := accessedPath + "main.md"
			if alteredPath == path {
				parentNode.IsActive = true
			}
		} else {
			alteredPath := accessedPath + ".md"
			if alteredPath == path {
				node.IsActive = true
			}
		}

		if strings.Contains(relativePath, ".md") {
			node.RelativePath = strings.TrimRight(relativePath, ".md")
		} else {
			node.RelativePath = strings.Join([]string{relativePath, "/"}, "")
		}

		if entry.Name() == "main.md" {
			title, order, err := GetFileTitleAndOrder(path)
			if err != nil {
				return err
			}
			(*parentNode).Title = title
			(*parentNode).Order = order

		} else if entry.IsDir() {
			node.Type = Dir
			err = buildTree(&node, path, directory, accessedPath)
			if err != nil {
				return err
			}
		} else {
			title, order, err := GetFileTitleAndOrder(node.Path)
			if err != nil {
				return err
			}
			node.Type = File
			node.Title = title
			node.Order = order
		}

		// Append the node to the parent
		sort.Slice(parentNode.Children, func(i, j int) bool {
			return parentNode.Children[i].Order < parentNode.Children[j].Order
		})

		parentNode.Children = append(parentNode.Children, node)
	}

	return nil
}

func GetFileTitleAndOrder(path string) (string, int64, error) {
	mainInput, err := os.ReadFile(path)
	if err != nil {
		log.WithError(err).Fatal("Error reading main documentation file.")
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

func GetRelativePath(url string, path string) string {
	relativePath, err := filepath.Rel(url, path)
	if err != nil {
		log.Fatal("Error while getting relative path")
	}

	return relativePath
}
